// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MpcContextDao is the data access object for table mpc_context.
type MpcContextDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns MpcContextColumns // columns contains all the column names of Table for convenient usage.
}

// MpcContextColumns defines and stores column names for table mpc_context.
type MpcContextColumns struct {
	UserId    string //
	Context   string //
	UpdatedAt string //
	Request   string //
	Token     string //
	CreatedAt string //
	DeletedAt string //
	PubKey    string //
}

// mpcContextColumns holds the columns for table mpc_context.
var mpcContextColumns = MpcContextColumns{
	UserId:    "user_id",
	Context:   "context",
	UpdatedAt: "updated_at",
	Request:   "request",
	Token:     "token",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
	PubKey:    "pub_key",
}

// NewMpcContextDao creates and returns a new DAO object for table data access.
func NewMpcContextDao() *MpcContextDao {
	return &MpcContextDao{
		group:   "default",
		table:   "mpc_context",
		columns: mpcContextColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MpcContextDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MpcContextDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MpcContextDao) Columns() MpcContextColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MpcContextDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MpcContextDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MpcContextDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
