// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthzAttributesDao is the data access object for table authz_attributes.
type AuthzAttributesDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns AuthzAttributesColumns // columns contains all the column names of Table for convenient usage.
}

// AuthzAttributesColumns defines and stores column names for table authz_attributes.
type AuthzAttributesColumns struct {
	Id    string //
	Key   string //
	Value string //
}

// authzAttributesColumns holds the columns for table authz_attributes.
var authzAttributesColumns = AuthzAttributesColumns{
	Id:    "id",
	Key:   "key",
	Value: "value",
}

// NewAuthzAttributesDao creates and returns a new DAO object for table data access.
func NewAuthzAttributesDao() *AuthzAttributesDao {
	return &AuthzAttributesDao{
		group:   "default",
		table:   "authz_attributes",
		columns: authzAttributesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthzAttributesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthzAttributesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthzAttributesDao) Columns() AuthzAttributesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthzAttributesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthzAttributesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthzAttributesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
