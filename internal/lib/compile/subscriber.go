package compile

import (
	"context"
	"regexp"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/vulcangz/gf2-authz/internal/event"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/spooler"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/observability/metric"
	"github.com/vulcangz/gf2-authz/internal/service"
	"golang.org/x/exp/slog"
)

var (
	Subscriber *subscriber
	once       sync.Once
)

type subscriber struct {
	compiler          Compiler
	dispatcher        event.Dispatcher
	flushDelay        time.Duration
	statsFlushDelay   time.Duration
	resourceKindRegex *regexp.Regexp
	metricEnabled     bool
	observer          metric.Observer
}

func NewSubscriber(
	compiler Compiler,
	dispatcher event.Dispatcher,
) *subscriber {
	ctx := context.Background()
	appCfg, _ := service.SysConfig().GetApp(ctx)
	obs, _ := metric.NewObserver(appCfg)

	return &subscriber{
		compiler:          compiler,
		dispatcher:        dispatcher,
		flushDelay:        appCfg.Audit.FlushDelay,
		statsFlushDelay:   appCfg.Stats.FlushDelay,
		resourceKindRegex: regexp.MustCompile(appCfg.Stats.ResourceKindRegex),
		metricEnabled:     appCfg.Metrics.Enabled,
		observer:          obs,
	}
}

func SubscriberInit(ctx context.Context) *subscriber {
	once.Do(func() {
		if Subscriber == nil {
			clock := ctime.NewClock()
			c := NewCompiler(clock)
			d := event.NewDispatcher(0, clock)
			Subscriber = NewSubscriber(c, d)
		}
	})

	return Subscriber
}

func (s *subscriber) Start(ctx context.Context) {
	policyEventChan := s.dispatcher.Subscribe(event.EventTypePolicy)
	principalEventChan := s.dispatcher.Subscribe(event.EventTypePrincipal)
	resourceEventChan := s.dispatcher.Subscribe(event.EventTypeResource)
	checkEventChan := s.dispatcher.Subscribe(event.EventTypeCheck)
	auditChan := s.dispatcher.Subscribe(event.EventTypeCheck)
	go s.handlePolicyEvents(policyEventChan)
	go s.handlePrincipalEvents(principalEventChan)
	go s.handleResourceEvents(resourceEventChan)
	go s.handleCheckEvents(checkEventChan)
	go s.handleAudit(auditChan)

	if s.metricEnabled {
		roleEventChan := s.dispatcher.Subscribe(event.EventTypeRole)
		// go s.handleItemEvents(policyEventChan, "policy")
		// go s.handleItemEvents(principalEventChan, "principal")
		// go s.handleItemEvents(resourceEventChan, "resource")
		go s.handleItemEvents(roleEventChan, "role")
	}

	g.Log().Info(ctx, "Compiler: subscribed to event dispatchers")
}

func (s *subscriber) Instance(ctx context.Context) *subscriber {
	return Subscriber
}

func (s *subscriber) Stop(ctx context.Context) {
	policyEventChan := s.dispatcher.Subscribe(event.EventTypePolicy)
	principalEventChan := s.dispatcher.Subscribe(event.EventTypePrincipal)
	resourceEventChan := s.dispatcher.Subscribe(event.EventTypeResource)
	checkEventChan := s.dispatcher.Subscribe(event.EventTypeCheck)
	close(policyEventChan)
	close(principalEventChan)
	close(resourceEventChan)
	close(checkEventChan)

	g.Log().Info(ctx, "Compiler: subscription to event dispatcher stopped")
}

func (s *subscriber) handlePolicyEvents(eventChan chan *event.Event) {
	for eventItem := range eventChan {
		itemEvent, ok := eventItem.Data.(*event.ItemEvent)
		if !ok {
			continue
		}

		if s.metricEnabled {
			s.observer.ObserveItemCreatedCounter("policy", string(itemEvent.Action))
		}

		policy := itemEvent.Data.(*entity.Policy)
		if err := s.compiler.CompilePolicy(context.Background(), policy); err != nil {
			g.Log().Warning(context.Background(),
				"Compiler: unable to compile policy",
				err,
				"policy_id", policy.ID,
			)
		}
	}
}

func (s *subscriber) handleResourceEvents(eventChan chan *event.Event) {
	for eventItem := range eventChan {
		itemEvent, ok := eventItem.Data.(*event.ItemEvent)
		if !ok {
			continue
		}

		if s.metricEnabled {
			s.observer.ObserveItemCreatedCounter("resource", string(itemEvent.Action))
		}

		resource := itemEvent.Data.(*entity.Resource)
		if err := s.compiler.CompileResource(resource); err != nil {
			g.Log().Warning(context.Background(),
				"Compiler: unable to compile resource",
				err,
				slog.Any("policy_id", resource.ID),
			)
		}
	}
}

func (s *subscriber) handlePrincipalEvents(eventChan chan *event.Event) {
	for eventItem := range eventChan {
		itemEvent, ok := eventItem.Data.(*event.ItemEvent)
		if !ok {
			continue
		}

		if s.metricEnabled {
			s.observer.ObserveItemCreatedCounter("principal", string(itemEvent.Action))
		}

		principal := itemEvent.Data.(*entity.Principal)
		if err := s.compiler.CompilePrincipal(principal); err != nil {
			g.Log().Warning(context.Background(),
				"Compiler: unable to compile principal",
				err,
				slog.Any("policy_id", principal.ID),
			)
		}
	}
}

func (s *subscriber) handleCheckEvents(eventChan chan *event.Event) {
	ctx := context.Background()
	var spooler = spooler.New(func(values []*event.Event) {
		if len(values) == 0 {
			return
		}

		var allowed, denied int64
		var timestamp int64

		for _, value := range values {
			timestamp = value.Timestamp

			checkEvent, ok := value.Data.(*event.CheckEvent)
			if !ok {
				continue
			}

			if s.metricEnabled {
				s.observer.ObserveCheckCounter(checkEvent.ResourceKind, checkEvent.IsAllowed)
			}

			if checkEvent.IsAllowed {
				allowed++
			} else {
				denied++
			}
		}

		if err := service.StatsManager().BatchAddCheck(ctx, timestamp, allowed, denied); err != nil {
			g.Log().Error(ctx, "Stats: unable to add check event", err)
		}
	}, spooler.WithFlushInterval(s.statsFlushDelay))

	for eventItem := range eventChan {
		checkEvent, ok := eventItem.Data.(*event.CheckEvent)
		if !ok {
			continue
		}

		if s.resourceKindRegex.MatchString(checkEvent.ResourceKind) {
			spooler.Add(eventItem)
		}
	}
}

func (s *subscriber) handleAudit(eventChan chan *event.Event) {
	ctx := context.Background()
	var spooler = spooler.New(func(values []*event.Event) {
		if len(values) == 0 {
			return
		}

		var audits = []*entity.Audit{}
		var timestamp int64

		for _, value := range values {
			timestamp = value.Timestamp

			checkEvent, ok := value.Data.(*event.CheckEvent)
			if !ok {
				continue
			}

			audit := &entity.Audit{
				Date:          time.Unix(timestamp, 0),
				Principal:     checkEvent.Principal,
				ResourceKind:  checkEvent.ResourceKind,
				ResourceValue: checkEvent.ResourceValue,
				Action:        checkEvent.Action,
				IsAllowed:     gconv.Int(checkEvent.IsAllowed),
			}

			if checkEvent.CompiledPolicy != nil {
				audit.PolicyId = checkEvent.CompiledPolicy.PolicyID
			}

			audits = append(audits, audit)
		}

		if err := service.AuditManager().BatchAdd(ctx, audits); err != nil {
			g.Log().Error(ctx, "Audit: unable to batch add audit events", err)
		}
	}, spooler.WithFlushInterval(s.flushDelay))

	for eventItem := range eventChan {
		checkEvent, ok := eventItem.Data.(*event.CheckEvent)
		if !ok {
			continue
		}

		if s.resourceKindRegex.MatchString(checkEvent.ResourceKind) {
			spooler.Add(eventItem)
		}
	}
}

func (s *subscriber) handleItemEvents(eventChan chan *event.Event, itemType string) {
	for eventItem := range eventChan {
		if !s.metricEnabled {
			continue
		}

		itemEvent, ok := eventItem.Data.(*event.ItemEvent)
		if !ok {
			continue
		}

		s.observer.ObserveItemCreatedCounter(itemType, string(itemEvent.Action))
	}
}
