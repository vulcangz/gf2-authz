package consts

type contextKey string

// ghttp.HookBeforeServe
var (
	ContextKey       contextKey = "authz_local_user_context"
	ResourceKindKey  contextKey = "authz_resource_kind"
	ResourceValueKey contextKey = "authz_resource_value"
	ActionKey        contextKey = "authz_action"
)

var (
	// gRPC
	// ClaimsKey is the context key used for storing claims.
	ClaimsKey         contextKey = "claims"
	UserIdentifierKey contextKey = "userIdentifier"
)
