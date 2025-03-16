// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthzUsersDao is the data access object for table authz_users.
type AuthzUsersDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns AuthzUsersColumns // columns contains all the column names of Table for convenient usage.
}

// AuthzUsersColumns defines and stores column names for table authz_users.
type AuthzUsersColumns struct {
	Username     string //
	PasswordHash string //
	CreatedAt    string //
	UpdatedAt    string //
}

// authzUsersColumns holds the columns for table authz_users.
var authzUsersColumns = AuthzUsersColumns{
	Username:     "username",
	PasswordHash: "password_hash",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewAuthzUsersDao creates and returns a new DAO object for table data access.
func NewAuthzUsersDao() *AuthzUsersDao {
	return &AuthzUsersDao{
		group:   "default",
		table:   "authz_users",
		columns: authzUsersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthzUsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthzUsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthzUsersDao) Columns() AuthzUsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthzUsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthzUsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthzUsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
