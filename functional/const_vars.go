package main

const (
	ActionGetKey  = "actions-get"
	ActionListKey = "actions-list"
	AuditGetKey   = "audits-get"

	ClientCreateKey = "clients-create"
	ClientDeleteKey = "clients-delete"
	ClientGetKey    = "clients-get"
	ClientListKey   = "clients-list"

	CompiledListKey = "compiled-list"

	PolicyCreateKey    = "policies-create"
	PolicyDeleteKey    = "policies-delete"
	PolicyGetKey       = "policies-get"
	PolicyListKey      = "policies-list"
	PolicyUpdateKey    = "policies-update"
	PrincipalCreateKey = "principals-create"
	PrincipalDeleteKey = "principals-delete"
	PrincipalGetKey    = "principals-get"
	PrincipalListKey   = "principals-list"
	PrincipalUpdateKey = "principals-update"
	ResourceCreateKey  = "resources-create"
	ResourceDeleteKey  = "resources-delete"
	ResourceGetKey     = "resources-get"
	ResourceListKey    = "resources-list"
	ResourceUpdateKey  = "resources-update"
	RoleCreateKey      = "roles-create"
	RoleDeleteKey      = "roles-delete"
	RoleGetKey         = "roles-get"
	RoleListKey        = "roles-list"
	RoleUpdateKey      = "roles-update"
	StatsGetKey        = "stats-get"
	UserCreateKey      = "users-create"
	UserDeleteKey      = "users-delete"
	UserGetKey         = "users-get"
	UserListKey        = "users-list"
)

var (
	// ResourcesAndActionsByMethod maps the resource kind and action for each
	// http handler method available in the rest API.
	ResourcesAndActionsByMethod = map[string][]string{
		ActionGetKey:  {"authz.actions", "get"},
		ActionListKey: {"authz.actions", "list"},

		AuditGetKey: {"authz.audits", "get"},

		ClientCreateKey: {"authz.clients", "create"},
		ClientDeleteKey: {"authz.clients", "delete"},
		ClientGetKey:    {"authz.clients", "get"},
		ClientListKey:   {"authz.clients", "list"},

		CompiledListKey: {"authz.compiled", "list"},

		PrincipalCreateKey: {"authz.policies", "create"},
		PrincipalDeleteKey: {"authz.policies", "delete"},
		PrincipalGetKey:    {"authz.policies", "get"},
		PrincipalListKey:   {"authz.policies", "list"},
		PrincipalUpdateKey: {"authz.policies", "update"},

		PolicyCreateKey: {"authz.principals", "create"},
		PolicyDeleteKey: {"authz.principals", "delete"},
		PolicyGetKey:    {"authz.principals", "get"},
		PolicyListKey:   {"authz.principals", "list"},
		PolicyUpdateKey: {"authz.principals", "update"},

		ResourceCreateKey: {"authz.resources", "create"},
		ResourceDeleteKey: {"authz.resources", "delete"},
		ResourceGetKey:    {"authz.resources", "get"},
		ResourceListKey:   {"authz.resources", "list"},
		ResourceUpdateKey: {"authz.resources", "update"},

		RoleCreateKey: {"authz.roles", "create"},
		RoleDeleteKey: {"authz.roles", "delete"},
		RoleGetKey:    {"authz.roles", "get"},
		RoleListKey:   {"authz.roles", "list"},
		RoleUpdateKey: {"authz.roles", "update"},

		StatsGetKey: {"authz.stats", "get"},

		UserCreateKey: {"authz.users", "create"},
		UserDeleteKey: {"authz.users", "delete"},
		UserGetKey:    {"authz.users", "get"},
		UserListKey:   {"authz.users", "list"},
	}
)
