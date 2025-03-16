// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthzOauthTokensDao is the data access object for table authz_oauth_tokens.
type AuthzOauthTokensDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns AuthzOauthTokensColumns // columns contains all the column names of Table for convenient usage.
}

// AuthzOauthTokensColumns defines and stores column names for table authz_oauth_tokens.
type AuthzOauthTokensColumns struct {
	Id        string //
	Code      string //
	Access    string //
	Refresh   string //
	Data      string //
	ExpiredAt string //
	CreatedAt string //
	UpdatedAt string //
}

// authzOauthTokensColumns holds the columns for table authz_oauth_tokens.
var authzOauthTokensColumns = AuthzOauthTokensColumns{
	Id:        "id",
	Code:      "code",
	Access:    "access",
	Refresh:   "refresh",
	Data:      "data",
	ExpiredAt: "expired_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAuthzOauthTokensDao creates and returns a new DAO object for table data access.
func NewAuthzOauthTokensDao() *AuthzOauthTokensDao {
	return &AuthzOauthTokensDao{
		group:   "default",
		table:   "authz_oauth_tokens",
		columns: authzOauthTokensColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthzOauthTokensDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthzOauthTokensDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthzOauthTokensDao) Columns() AuthzOauthTokensColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthzOauthTokensDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthzOauthTokensDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthzOauthTokensDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
