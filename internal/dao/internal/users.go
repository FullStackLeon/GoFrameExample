// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for table users.
type UsersDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns UsersColumns // columns contains all the column names of Table for convenient usage.
}

// UsersColumns defines and stores column names for table users.
type UsersColumns struct {
	UserId           string //
	Username         string //
	Email            string //
	Password         string //
	AccountType      string //
	AuthProvider     string //
	AuthUserId       string //
	AuthAccessToken  string //
	AuthRefreshToken string //
	AuthUserInfo     string //
	Balance          string //
	Source           string //
	Status           string //
	GoogleAuthKey    string //
	TwoFactorAuth    string //
	KycStatus        string //
	CreatedAt        string //
	UpdatedAt        string //
	DeletedAt        string //
}

// usersColumns holds the columns for table users.
var usersColumns = UsersColumns{
	UserId:           "user_id",
	Username:         "username",
	Email:            "email",
	Password:         "password",
	AccountType:      "account_type",
	AuthProvider:     "auth_provider",
	AuthUserId:       "auth_user_id",
	AuthAccessToken:  "auth_access_token",
	AuthRefreshToken: "auth_refresh_token",
	AuthUserInfo:     "auth_user_info",
	Balance:          "balance",
	Source:           "source",
	Status:           "status",
	GoogleAuthKey:    "google_auth_key",
	TwoFactorAuth:    "two_factor_auth",
	KycStatus:        "kyc_status",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   "default",
		table:   "users",
		columns: usersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
