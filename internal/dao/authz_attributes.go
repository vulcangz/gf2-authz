// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/vulcangz/gf2-authz/internal/dao/internal"
)

// internalAuthzAttributesDao is internal type for wrapping internal DAO implements.
type internalAuthzAttributesDao = *internal.AuthzAttributesDao

// authzAttributesDao is the data access object for table authz_attributes.
// You can define custom methods on it to extend its functionality as you wish.
type authzAttributesDao struct {
	internalAuthzAttributesDao
}

var (
	// Attribute is globally public accessible object for table authz_attributes operations.
	Attribute = authzAttributesDao{
		internal.NewAuthzAttributesDao(),
	}
)

// Fill with you ideas below.
