// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/vulcangz/gf2-authz/internal/dao/internal"
)

// internalAuthzPoliciesDao is internal type for wrapping internal DAO implements.
type internalAuthzPoliciesDao = *internal.AuthzPoliciesDao

// authzPoliciesDao is the data access object for table authz_policies.
// You can define custom methods on it to extend its functionality as you wish.
type authzPoliciesDao struct {
	internalAuthzPoliciesDao
}

var (
	// Policy is globally public accessible object for table authz_policies operations.
	Policy = authzPoliciesDao{
		internal.NewAuthzPoliciesDao(),
	}
)

// Fill with you ideas below.
