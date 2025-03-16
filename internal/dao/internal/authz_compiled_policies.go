// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthzCompiledPoliciesDao is the data access object for table authz_compiled_policies.
type AuthzCompiledPoliciesDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns AuthzCompiledPoliciesColumns // columns contains all the column names of Table for convenient usage.
}

// AuthzCompiledPoliciesColumns defines and stores column names for table authz_compiled_policies.
type AuthzCompiledPoliciesColumns struct {
	PolicyId      string //
	PrincipalId   string //
	ResourceKind  string //
	ResourceValue string //
	ActionId      string //
	Version       string //
	CreatedAt     string //
	UpdatedAt     string //
}

// authzCompiledPoliciesColumns holds the columns for table authz_compiled_policies.
var authzCompiledPoliciesColumns = AuthzCompiledPoliciesColumns{
	PolicyId:      "policy_id",
	PrincipalId:   "principal_id",
	ResourceKind:  "resource_kind",
	ResourceValue: "resource_value",
	ActionId:      "action_id",
	Version:       "version",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewAuthzCompiledPoliciesDao creates and returns a new DAO object for table data access.
func NewAuthzCompiledPoliciesDao() *AuthzCompiledPoliciesDao {
	return &AuthzCompiledPoliciesDao{
		group:   "default",
		table:   "authz_compiled_policies",
		columns: authzCompiledPoliciesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthzCompiledPoliciesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthzCompiledPoliciesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthzCompiledPoliciesDao) Columns() AuthzCompiledPoliciesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthzCompiledPoliciesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthzCompiledPoliciesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthzCompiledPoliciesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
