// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	UserId           int         `json:"userId"           ` //
	Username         string      `json:"username"         ` //
	Email            string      `json:"email"            ` //
	Password         string      `json:"password"         ` //
	AccountType      string      `json:"accountType"      ` //
	AuthProvider     string      `json:"authProvider"     ` //
	AuthUserId       string      `json:"authUserId"       ` //
	AuthAccessToken  string      `json:"authAccessToken"  ` //
	AuthRefreshToken string      `json:"authRefreshToken" ` //
	AuthUserInfo     string      `json:"authUserInfo"     ` //
	Balance          float64     `json:"balance"          ` //
	Source           string      `json:"source"           ` //
	Status           string      `json:"status"           ` //
	GoogleAuthKey    string      `json:"googleAuthKey"    ` //
	TwoFactorAuth    string      `json:"twoFactorAuth"    ` //
	KycStatus        string      `json:"kycStatus"        ` //
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` //
	DeletedAt        *gtime.Time `json:"deletedAt"        ` //
}
