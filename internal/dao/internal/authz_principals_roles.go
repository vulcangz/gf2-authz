// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthzPrincipalsRolesDao is the data access object for table authz_principals_roles.
type AuthzPrincipalsRolesDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns AuthzPrincipalsRolesColumns // columns contains all the column names of Table for convenient usage.
}

// AuthzPrincipalsRolesColumns defines and stores column names for table authz_principals_roles.
type AuthzPrincipalsRolesColumns struct {
	RoleId      string //
	PrincipalId string //
}

// authzPrincipalsRolesColumns holds the columns for table authz_principals_roles.
var authzPrincipalsRolesColumns = AuthzPrincipalsRolesColumns{
	RoleId:      "role_id",
	PrincipalId: "principal_id",
}

// NewAuthzPrincipalsRolesDao creates and returns a new DAO object for table data access.
func NewAuthzPrincipalsRolesDao() *AuthzPrincipalsRolesDao {
	return &AuthzPrincipalsRolesDao{
		group:   "default",
		table:   "authz_principals_roles",
		columns: authzPrincipalsRolesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthzPrincipalsRolesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthzPrincipalsRolesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthzPrincipalsRolesDao) Columns() AuthzPrincipalsRolesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthzPrincipalsRolesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthzPrincipalsRolesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthzPrincipalsRolesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
