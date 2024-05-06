package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegReq struct {
	g.Meta      `method:"post"`
	Nickname    string `v:"required#nickname字段不能为空"`
	Email       string `v:"required|email#email字段不能为空|email格式不正确"`
	Password    string `v:"required|password|same:Password2#password字段不能为空|password格式不正确|账户密码和确认密码要一致"`
	Password2   string `v:"required|password#password2字段不能为空|password2格式不正确"`
	AccountType string `v:"required|in:regular,auth,admin#账户类型字段可选范围 regular,auth,admin" dc:"regular:普通用户;auth:Auth用户;admin:管理员用户"`
}

type RegRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Id     uint64 `json:"id"`
}

type LoginReq struct {
	g.Meta   `method:"post"`
	Email    string `v:"required|email#email字段不能为空|email格式不正确"`
	Password string `v:"required|password#password字段不能为空|password格式不正确"`
}

type LoginRes struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
