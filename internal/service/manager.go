package service

import (
	"context"

	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

type (
	IAuditManager interface {
		GetRepository() orm.AuditRepository
		BatchAdd(ctx context.Context, audits []*entity.Audit) error
	}
	IStatsManager interface {
		GetRepository() orm.StatsRepository
		BatchAddCheck(ctx context.Context, timestamp int64, allowed int64, denied int64) error
	}
	IUserManager interface {
		GetRepository() orm.UserRepository
		Create(ctx context.Context, username string, password string) (*entity.User, error)
		Delete(ctx context.Context, username string) error
		UpdatePassword(username string, password string) error
	}
	IActionManager interface {
		GetRepository() orm.ActionRepository
		Create(ctx context.Context, identifier string) (*entity.Action, error)
	}
	IAttributeManager interface {
		GetRepository() orm.AttributeRepository
		MapToSlice(ctx context.Context, attributes map[string]any) ([]*entity.Attribute, error)
	}
	IClientManager interface {
		GetRepository() orm.ClientRepository
		Create(ctx context.Context, name string, domain string) (*entity.Client, error)
		Delete(ctx context.Context, identifier string) error
	}
	ICompiledPolicyManager interface {
		GetRepository() orm.CompiledPolicyRepository
		Create(ctx context.Context, compiledPolicy []*entity.CompiledPolicy) error
		IsAllowed(ctx context.Context, principalID string, resourceKind string, resourceValue string, actionID string) (bool, error)
	}
	IMockCompiledPolicy interface {
		// EXPECT() *MockCompiledPolicyMockRecorder
		Create(ctx context.Context, compiledPolicy []*entity.CompiledPolicy) error
		GetRepository() orm.CompiledPolicyRepository
		IsAllowed(ctx context.Context, principalID, resourceKind, resourceValue, actionID string) (bool, error)
	}
	IPolicyManager interface {
		GetRepository() orm.PolicyRepository
		Create(ctx context.Context, identifier string, resources []string, actions []string, attributeRules []string) (*entity.Policy, error)
		Delete(ctx context.Context, identifier string) error
		Update(ctx context.Context, identifier string, resources []string, actions []string, attributeRules []string) (*entity.Policy, error)
	}
	IPrincipalManager interface {
		GetRepository() orm.PrincipalRepository
		Create(ctx context.Context, identifier string, roles []string, attributes map[string]any) (*entity.Principal, error)
		Update(ctx context.Context, identifier string, roles []string, attributes map[string]any) (*entity.Principal, error)
		Delete(ctx context.Context, identifier string) error
		FindMatchingAttribute(principalAttribute string) ([]*entity.PrincipalMatchingAttribute, error)
	}
	IResourceManager interface {
		GetRepository() orm.ResourceRepository
		Create(ctx context.Context, identifier string, kind string, value string, attributes map[string]any) (*entity.Resource, error)
		Delete(ctx context.Context, identifier string) error
		Update(ctx context.Context, identifier string, kind string, value string, attributes map[string]any) (*entity.Resource, error)
		FindMatchingAttribute(resourceAttribute string, options ...entity.ResourceQueryOption) ([]*entity.ResourceMatchingAttribute, error)
	}
	IRoleManager interface {
		GetRepository() orm.RoleRepository
		Create(ctx context.Context, identifier string, policies []string) (*entity.Role, error)
		Delete(ctx context.Context, identifier string) error
		Update(ctx context.Context, identifier string, policies []string) (*entity.Role, error)
	}
)

var (
	localResourceManager       IResourceManager
	localActionManager         IActionManager
	localAttributeManager      IAttributeManager
	localClientManager         IClientManager
	localCompiledPolicyManager ICompiledPolicyManager
	localMockCompiledPolicy    IMockCompiledPolicy
	localPolicyManager         IPolicyManager
	localPrincipalManager      IPrincipalManager
	localRoleManager           IRoleManager
	localAuditManager          IAuditManager
	localStatsManager          IStatsManager
	localUserManager           IUserManager
)

func ResourceManager() IResourceManager {
	if localResourceManager == nil {
		panic("implement not found for interface IResourceManager, forgot register?")
	}
	return localResourceManager
}

func RegisterResourceManager(i IResourceManager) {
	localResourceManager = i
}

func ActionManager() IActionManager {
	if localActionManager == nil {
		panic("implement not found for interface IActionManager, forgot register?")
	}
	return localActionManager
}

func RegisterActionManager(i IActionManager) {
	localActionManager = i
}

func AttributeManager() IAttributeManager {
	if localAttributeManager == nil {
		panic("implement not found for interface IAttributeManager, forgot register?")
	}
	return localAttributeManager
}

func RegisterAttributeManager(i IAttributeManager) {
	localAttributeManager = i
}

func ClientManager() IClientManager {
	if localClientManager == nil {
		panic("implement not found for interface IClientManager, forgot register?")
	}
	return localClientManager
}

func RegisterClientManager(i IClientManager) {
	localClientManager = i
}

func CompiledPolicyManager() ICompiledPolicyManager {
	if localCompiledPolicyManager == nil {
		panic("implement not found for interface ICompiledPolicyManager, forgot register?")
	}
	return localCompiledPolicyManager
}

func RegisterCompiledPolicyManager(i ICompiledPolicyManager) {
	localCompiledPolicyManager = i
}

func MockCompiledPolicy() IMockCompiledPolicy {
	if localMockCompiledPolicy == nil {
		panic("implement not found for interface IMockCompiledPolicy, forgot register?")
	}
	return localMockCompiledPolicy
}

func RegisterMockCompiledPolicy(i IMockCompiledPolicy) {
	localMockCompiledPolicy = i
}

func PolicyManager() IPolicyManager {
	if localPolicyManager == nil {
		panic("implement not found for interface IPolicyManager, forgot register?")
	}
	return localPolicyManager
}

func RegisterPolicyManager(i IPolicyManager) {
	localPolicyManager = i
}

func PrincipalManager() IPrincipalManager {
	if localPrincipalManager == nil {
		panic("implement not found for interface IPrincipalManager, forgot register?")
	}
	return localPrincipalManager
}

func RegisterPrincipalManager(i IPrincipalManager) {
	localPrincipalManager = i
}

func RoleManager() IRoleManager {
	if localRoleManager == nil {
		panic("implement not found for interface IRoleManager, forgot register?")
	}
	return localRoleManager
}

func RegisterRoleManager(i IRoleManager) {
	localRoleManager = i
}

func AuditManager() IAuditManager {
	if localAuditManager == nil {
		panic("implement not found for interface IAuditManager, forgot register?")
	}
	return localAuditManager
}

func RegisterAuditManager(i IAuditManager) {
	localAuditManager = i
}

func StatsManager() IStatsManager {
	if localStatsManager == nil {
		panic("implement not found for interface IStatsManager, forgot register?")
	}
	return localStatsManager
}

func RegisterStatsManager(i IStatsManager) {
	localStatsManager = i
}

func UserManager() IUserManager {
	if localUserManager == nil {
		panic("implement not found for interface IUserManager, forgot register?")
	}
	return localUserManager
}

func RegisterUserManager(i IUserManager) {
	localUserManager = i
}
