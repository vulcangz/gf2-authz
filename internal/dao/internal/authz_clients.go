// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthzClientsDao is the data access object for table authz_clients.
type AuthzClientsDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns AuthzClientsColumns // columns contains all the column names of Table for convenient usage.
}

// AuthzClientsColumns defines and stores column names for table authz_clients.
type AuthzClientsColumns struct {
	Id        string //
	Secret    string //
	Name      string //
	Domain    string //
	Data      string //
	CreatedAt string //
	UpdatedAt string //
}

// authzClientsColumns holds the columns for table authz_clients.
var authzClientsColumns = AuthzClientsColumns{
	Id:        "id",
	Secret:    "secret",
	Name:      "name",
	Domain:    "domain",
	Data:      "data",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAuthzClientsDao creates and returns a new DAO object for table data access.
func NewAuthzClientsDao() *AuthzClientsDao {
	return &AuthzClientsDao{
		group:   "default",
		table:   "authz_clients",
		columns: authzClientsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthzClientsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthzClientsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthzClientsDao) Columns() AuthzClientsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthzClientsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthzClientsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthzClientsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
