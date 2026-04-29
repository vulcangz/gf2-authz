package fixtures

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

const (
	defaultAdminUsername = "admin"
)

var (
	appName, defaultAdminPassword string

	resources = map[string][]string{
		"actions":    {"list", "get"},
		"audits":     {"get"},
		"clients":    {"list", "get", "create", "delete"},
		"compiled":   {"list"},
		"policies":   {"list", "get", "create", "update", "delete"},
		"principals": {"list", "get", "create", "update", "delete"},
		"resources":  {"list", "get", "create", "update", "delete"},
		"roles":      {"list", "get", "create", "update", "delete"},
		"stats":      {"get"},
		"users":      {"list", "get", "create", "delete"},
	}

	ctx = context.Background()
)

type Initializer interface {
	Initialize() error
}

type initializer struct {
	cfg      *entity.UserConfig
	policies []string
}

func NewInitializer(
	cfg *entity.UserConfig,
) Initializer {
	return &initializer{
		cfg:      cfg,
		policies: []string{},
	}
}

// Initialize initializes default application resources.
func (i *initializer) Initialize() error {
	appCfg, _ := service.SysConfig().GetApp(ctx)
	appName = appCfg.Name
	defaultAdminPassword = i.cfg.AdminDefaultPassword

	database.GetDatabase(ctx)

	if i.checkAlreadyInitialized() {
		return nil
	}

	if err := i.initializeResources(); err != nil {
		if !strings.Contains(err.Error(), "a resource already exists with id") {
			return err
		}

		g.Log().Debug(ctx, err.Error())
	}

	if err := i.initializePolicies(); err != nil {
		if !strings.Contains(err.Error(), "a policy already exists with identifier") {
			return err
		}

		g.Log().Debug(ctx, err.Error())
	}

	if err := i.initializeRoles(); err != nil {
		if !strings.Contains(err.Error(), "a role already exists with identifier") {
			return err
		}

		g.Log().Debug(ctx, err.Error())
	}

	if err := i.initializeUser(); err != nil {
		if !strings.Contains(err.Error(), "Duplicate entry 'authz-user-admin'") {
			return err
		}

		g.Log().Debug(ctx, err.Error())
	}

	g.Log().Info(ctx, "initialize end.")

	return nil
}

func (i *initializer) initializeResources() error {
	for resourceType := range resources {
		resource, err := service.ResourceManager().Create(
			ctx,
			fmt.Sprintf("%s.%s.%s", appName, resourceType, "*"),
			fmt.Sprintf("%s.%s", appName, resourceType),
			"*",
			nil,
		)
		if err != nil {
			return err
		}

		resource.IsLocked = true

		if err = service.ResourceManager().GetRepository().Update(resource); err != nil {
			return err
		}
	}

	return nil
}

func (i *initializer) checkAlreadyInitialized() bool {
	adminUser, err := service.UserManager().GetRepository().GetByFields(map[string]orm.FieldValue{
		"username": {Operator: "=", Value: defaultAdminUsername},
	})

	if err == nil && adminUser != nil {
		_ = service.UserManager().UpdatePassword(defaultAdminUsername, defaultAdminPassword)
		return true
	}

	return false
}

func (i *initializer) initializePolicies() error {
	for resourceType, actions := range resources {
		policy, err := service.PolicyManager().Create(
			ctx,
			fmt.Sprintf("%s-%s-admin", appName, resourceType),
			[]string{
				fmt.Sprintf("%s.%s.%s", appName, resourceType, "*"),
			},
			actions,
			nil,
		)
		if err != nil {
			return err
		}

		i.policies = append(i.policies, policy.ID)
	}

	return nil
}

func (i *initializer) initializeRoles() error {
	_, err := service.RoleManager().Create(
		ctx,
		fmt.Sprintf("%s-admin", appName),
		i.policies,
	)
	if err != nil {
		return err
	}

	return nil
}

func (i *initializer) initializeUser() error {
	adminUser, err := service.UserManager().GetRepository().GetByFields(map[string]orm.FieldValue{
		"username": {Operator: "=", Value: defaultAdminUsername},
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		g.Log().Debugf(ctx, "Unable to get default admin user: %v", err)
		return err
	}

	if adminUser != nil {
		// User already exists, nothing to do.
		return nil
	}

	// Create user "admin" and principal named "authz-user-admin"
	user, err := service.UserManager().Create(ctx, defaultAdminUsername, defaultAdminPassword)
	if err != nil {
		g.Log().Debugf(ctx, "unable to create default admin user: %v", err)
		return err
	}

	// Retrieve principal created following the user creation
	principal, err := service.PrincipalManager().GetRepository().Get(
		entity.UserPrincipal(user.Username),
	)
	if err != nil {
		g.Log().Debugf(ctx, "unable to retrieve admin principal: %v", err)
		return err
	}

	principal.IsLocked = true

	// Attach role "authz-admin" to user principal "authz-admin"
	role, err := service.RoleManager().GetRepository().Get(fmt.Sprintf("%s-admin", appName))
	if err != nil {
		g.Log().Debugf(ctx, "unable to retrieve admin role: %v", err)
		return err
	}

	principal.Roles = []*entity.Role{role}

	return service.PrincipalManager().GetRepository().Update(principal)
}
