// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthzPoliciesDao is the data access object for table authz_policies.
type AuthzPoliciesDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AuthzPoliciesColumns // columns contains all the column names of Table for convenient usage.
}

// AuthzPoliciesColumns defines and stores column names for table authz_policies.
type AuthzPoliciesColumns struct {
	Id             string //
	AttributeRules string //
	CreatedAt      string //
	UpdatedAt      string //
}

// authzPoliciesColumns holds the columns for table authz_policies.
var authzPoliciesColumns = AuthzPoliciesColumns{
	Id:             "id",
	AttributeRules: "attribute_rules",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewAuthzPoliciesDao creates and returns a new DAO object for table data access.
func NewAuthzPoliciesDao() *AuthzPoliciesDao {
	return &AuthzPoliciesDao{
		group:   "default",
		table:   "authz_policies",
		columns: authzPoliciesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthzPoliciesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthzPoliciesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthzPoliciesDao) Columns() AuthzPoliciesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthzPoliciesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthzPoliciesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthzPoliciesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
