// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthzStatsDao is the data access object for table authz_stats.
type AuthzStatsDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns AuthzStatsColumns // columns contains all the column names of Table for convenient usage.
}

// AuthzStatsColumns defines and stores column names for table authz_stats.
type AuthzStatsColumns struct {
	Id                  string //
	Date                string //
	ChecksAllowedNumber string //
	ChecksDeniedNumber  string //
}

// authzStatsColumns holds the columns for table authz_stats.
var authzStatsColumns = AuthzStatsColumns{
	Id:                  "id",
	Date:                "date",
	ChecksAllowedNumber: "checks_allowed_number",
	ChecksDeniedNumber:  "checks_denied_number",
}

// NewAuthzStatsDao creates and returns a new DAO object for table data access.
func NewAuthzStatsDao() *AuthzStatsDao {
	return &AuthzStatsDao{
		group:   "default",
		table:   "authz_stats",
		columns: authzStatsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthzStatsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthzStatsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthzStatsDao) Columns() AuthzStatsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthzStatsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthzStatsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthzStatsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
