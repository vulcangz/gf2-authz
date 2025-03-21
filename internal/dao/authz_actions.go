// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/vulcangz/gf2-authz/internal/dao/internal"
)

// internalAuthzActionsDao is internal type for wrapping internal DAO implements.
type internalAuthzActionsDao = *internal.AuthzActionsDao

// authzActionsDao is the data access object for table authz_actions.
// You can define custom methods on it to extend its functionality as you wish.
type authzActionsDao struct {
	internalAuthzActionsDao
}

var (
	// Action is globally public accessible object for table authz_actions operations.
	Action = authzActionsDao{
		internal.NewAuthzActionsDao(),
	}
)

// Fill with you ideas below.
