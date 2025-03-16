// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/vulcangz/gf2-authz/internal/dao/internal"
)

// internalAuthzRolesPoliciesDao is internal type for wrapping internal DAO implements.
type internalAuthzRolesPoliciesDao = *internal.AuthzRolesPoliciesDao

// authzRolesPoliciesDao is the data access object for table authz_roles_policies.
// You can define custom methods on it to extend its functionality as you wish.
type authzRolesPoliciesDao struct {
	internalAuthzRolesPoliciesDao
}

var (
	// AuthzRolesPolicies is globally public accessible object for table authz_roles_policies operations.
	AuthzRolesPolicies = authzRolesPoliciesDao{
		internal.NewAuthzRolesPoliciesDao(),
	}
)

// Fill with you ideas below.
