package event

import (
	"context"
	"errors"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
)

var (
	instance *dispatcher
	once     sync.Once

	ErrEventChanCast            = errors.New("unable to cast event channel to []chan *Event")
	ErrNoSubscriberForEventType = errors.New("no subscriber for this event type")
)

type Dispatcher interface {
	Dispatch(eventType EventType, data any) error
	Subscribe(eventType EventType) chan *Event
	Unsubscribe(eventType EventType, eventChanToClose chan *Event) error
}

type dispatcher struct {
	clock         ctime.Clock
	subscribers   *sync.Map
	eventChanSize int
}

func NewDispatcher(
	eventChanSize int,
	clock ctime.Clock,
) *dispatcher {
	once.Do(func() {
		if eventChanSize == 0 {
			eventChanSize = g.Cfg().MustGet(context.Background(), "event.dispatcherChannelSize").Int()
		}

		instance = &dispatcher{
			clock:         clock,
			subscribers:   &sync.Map{},
			eventChanSize: eventChanSize,
		}
	})
	return instance
}

func (n *dispatcher) Dispatch(eventType EventType, data any) error {
	eventChanSlice, ok := n.subscribers.Load(eventType)
	if !ok {
		return nil
	}

	eventChans, ok := eventChanSlice.([]chan *Event)
	if !ok {
		return ErrEventChanCast
	}

	for _, eventChan := range eventChans {
		eventChan <- &Event{
			Data:      data,
			Timestamp: n.clock.Now().Unix(),
		}
	}

	return nil
}

func (n *dispatcher) Subscribe(eventType EventType) chan *Event {
	eventChan := make(chan *Event, n.eventChanSize)

	eventChanSlice, ok := n.subscribers.Load(eventType)
	if ok {
		eventChanSlice = append(eventChanSlice.([]chan *Event), eventChan)
	} else {
		eventChanSlice = []chan *Event{eventChan}
	}

	n.subscribers.Store(eventType, eventChanSlice)

	return eventChan
}

func (n *dispatcher) Unsubscribe(eventType EventType, eventChanToClose chan *Event) error {
	eventChanSlice, ok := n.subscribers.Load(eventType)
	if !ok {
		return ErrNoSubscriberForEventType
	}

	eventChans, ok := eventChanSlice.([]chan *Event)
	if !ok {
		return ErrEventChanCast
	}

	for index, eventChan := range eventChans {
		if eventChan != eventChanToClose {
			continue
		}

		eventChans = append(eventChans[:index], eventChans[index+1:]...)
		n.subscribers.LoadOrStore(eventType, eventChans)
	}

	return nil
}
