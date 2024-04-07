package user

import (
	"encoding/json"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"GoFrameExample/api"
	v1 "GoFrameExample/api/user/v1"
	"GoFrameExample/internal/consts"
)

type Profile struct{}

func NewProfile() *Profile {
	return &Profile{}
}

func (p *Profile) RegularProfile(req *ghttp.Request) {
	md := g.Model("users")
	user, err := md.Where(`email = ?`, req.Get("email")).Where(`account_type = ?`, consts.RegularAccountType).One()
	if err != nil {
		req.Response.WriteStatus(http.StatusInternalServerError)
		return
	}

	if user.IsEmpty() {
		req.Response.WriteJson(g.Map{
			"code":    consts.ErrCodeAccountNotExist,
			"message": consts.ErrMsgAccountNotExist.Error(),
		})
		return
	}

	res := api.Res{
		Code:    consts.ErrCodeSuccess,
		Message: consts.ErrMsgSuccess.Error(),
		Data: &v1.RegularUserProfile{
			Email:       user["email"].String(),
			Username:    user["username"].String(),
			AccountType: user["account_type"].String(),
		},
	}
	req.Response.WriteJson(res)
	return
}

func (p *Profile) AuthProfile(req *ghttp.Request) {
	md := g.Model("users")
	user, err := md.Where(`email = ?`, req.Get("email")).Where(`account_type = ?`, consts.AuthAccountType).One()
	if err != nil {
		req.Response.WriteStatus(http.StatusInternalServerError)
		return
	}

	if user.IsEmpty() {
		req.Response.WriteJson(g.Map{
			"code":    consts.ErrCodeAccountNotExist,
			"message": consts.ErrMsgAccountNotExist.Error(),
		})
		return
	}

	var authUserInfo v1.AuthUserInfo

	_ = json.Unmarshal([]byte(user["auth_user_info"].String()), &authUserInfo)
	res := api.Res{
		Code:    consts.ErrCodeSuccess,
		Message: consts.ErrMsgSuccess.Error(),
		Data: &v1.AuthUserProfile{
			Email:           user["email"].String(),
			Username:        user["username"].String(),
			AccountType:     user["account_type"].String(),
			AuthProvider:    user["auth_provider"].String(),
			AuthUserInfo:    authUserInfo,
			AuthAccessToken: user["auth_access_token"].String(),
		},
	}
	req.Response.WriteJson(res)
	return
}

func (p *Profile) AdminProfile(req *ghttp.Request) {
	md := g.Model("users")
	user, err := md.Where(`email = ?`, req.Get("email")).Where(`account_type = ?`, consts.AdminAccountType).One()
	if err != nil {
		req.Response.WriteStatus(http.StatusInternalServerError)
		return
	}

	if user.IsEmpty() {
		req.Response.WriteJson(g.Map{
			"code":    consts.ErrCodeAccountNotExist,
			"message": consts.ErrMsgAccountNotExist.Error(),
		})
		return
	}

	res := api.Res{
		Code:    consts.ErrCodeSuccess,
		Message: consts.ErrMsgSuccess.Error(),
		Data: &v1.AdminUserProfile{
			Email:       user["email"].String(),
			Username:    user["username"].String(),
			AccountType: user["account_type"].String(),
		},
	}
	req.Response.WriteJson(res)
	return
}
