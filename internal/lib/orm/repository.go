package orm

import "github.com/vulcangz/gf2-authz/internal/model/entity"

// Repositories is a constraint interface that allows only authz library repositories.
type Repositories interface {
	base[entity.Action] |
		base[entity.Client] |
		base[entity.CompiledPolicy] |
		base[entity.Policy] |
		base[entity.Principal] |
		base[entity.Resource] |
		base[entity.Role] |
		base[entity.User]
}

type (
	ActionRepository         Base[entity.Action]
	AttributeRepository      Base[entity.Attribute]
	AuditRepository          Base[entity.Audit]
	ClientRepository         Base[entity.Client]
	CompiledPolicyRepository Base[entity.CompiledPolicy]
	PolicyRepository         Base[entity.Policy]
	PrincipalRepository      Base[entity.Principal]
	ResourceRepository       Base[entity.Resource]
	RoleRepository           Base[entity.Role]
	StatsRepository          Base[entity.Stats]
	UserRepository           Base[entity.User]
)
