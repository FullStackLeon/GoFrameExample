package v1

import "github.com/gogf/gf/v2/frame/g"

type ProfileReq struct {
	Email string `v:"required|email#email字段不能为空|email格式不正确"`
}
type RegularReq struct {
	g.Meta `method:"get"`
	ProfileReq
}

type AuthReq struct {
	g.Meta `method:"get"`
	ProfileReq
}

type AdminReq struct {
	g.Meta `method:"get"`
	ProfileReq
}

type AuthUserInfo struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}
type ProfileRes struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccountType string `json:"account_type"`
}

type RegularRes struct {
	ProfileRes
}

type AuthRes struct {
	ProfileRes
	AuthProvider    string       `json:"auth_provider"`
	AuthUserInfo    AuthUserInfo `json:"auth_user_profile"`
	AuthAccessToken string       `json:"auth_access_token"`
}

type AdminRes struct {
	ProfileRes
}
