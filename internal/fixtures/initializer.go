package fixtures

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

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
	g.Log().Info(ctx, "initialize start...")

	appCfg, _ := service.SysConfig().GetApp(ctx)
	appName = appCfg.Name
	defaultAdminPassword = i.cfg.AdminDefaultPassword

	database.GetDatabase(ctx)

	if i.checkAlreadyInitialized() {
		return nil
	}

	if err := i.initializeResources(); err != nil {
		return err
	}

	if err := i.initializePolicies(); err != nil {
		return err
	}

	if err := i.initializeRoles(); err != nil {
		return err
	}

	if err := i.initializeUser(); err != nil {
		return err
	}

	g.Log().Info(ctx, "initialize end.")

	return nil
}

func (i *initializer) initializeResources() error {
	g.Log().Info(ctx, "initializeResources start……")
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

	g.Log().Info(ctx, "initializeResources ok.")
	return nil
}

func (i *initializer) checkAlreadyInitialized() bool {
	adminUser, err := service.UserManager().GetRepository().GetByFields(map[string]orm.FieldValue{
		"username": {Operator: "=", Value: defaultAdminUsername},
	})

	if err == nil && adminUser != nil {
		_ = service.UserManager().UpdatePassword(defaultAdminUsername, defaultAdminPassword)
		g.Log().Info(ctx, "checkAlreadyInitialized update password ok.")
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

	g.Log().Info(ctx, "initializePolicies ok.")
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

	g.Log().Info(ctx, "initializeRoles ok.")
	return nil
}

func (i *initializer) initializeUser() error {
	adminUser, err := service.UserManager().GetRepository().GetByFields(map[string]orm.FieldValue{
		"username": {Operator: "=", Value: defaultAdminUsername},
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		g.Log().Error(ctx, "Unable to get default admin user", slog.Any("err", err))
		return err
	}

	if adminUser != nil {
		// User already exists, nothing to do.
		return nil
	}

	// Create user "admin" and principal named "authz-user-admin"
	user, err := service.UserManager().Create(ctx, defaultAdminUsername, defaultAdminPassword)
	if err != nil {
		return fmt.Errorf("unable to create default admin user: %v", err)
	}

	// Retrieve principal created following the user creation
	principal, err := service.PrincipalManager().GetRepository().Get(
		entity.UserPrincipal(user.Username),
	)
	if err != nil {
		return fmt.Errorf("unable to retrieve admin principal: %v", err)
	}

	principal.IsLocked = true

	// Attach role "authz-admin" to user principal "authz-admin"
	role, err := service.RoleManager().GetRepository().Get(fmt.Sprintf("%s-admin", appName))
	if err != nil {
		return fmt.Errorf("unable to retrieve admin role: %v", err)
	}

	principal.Roles = []*entity.Role{role}

	g.Log().Info(ctx, "initializeUser start update……")
	return service.PrincipalManager().GetRepository().Update(principal)
}
