package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"GoFrameExample/api/profile/v1"
	"GoFrameExample/internal/consts"
)

type Profile struct{}

func NewProfile() *Profile {
	return &Profile{}
}

func (p *Profile) RegularProfile(ctx context.Context, req *v1.RegularProfileReq) (res *v1.RegularProfileRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	md := g.Model("users")
	user, err := md.Where(`email = ?`, req.Email).Where(`account_type = ?`, consts.RegularAccountType).One()
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return
	}

	if user.IsEmpty() {
		r.Response.WriteJson(g.Map{
			"code":    consts.ErrCodeAccountNotExist,
			"message": consts.ErrMsgAccountNotExist.Error(),
		})
		return
	}

	res = &v1.RegularProfileRes{
		ProfileRes: v1.ProfileRes{
			Email:       user["email"].String(),
			Username:    user["username"].String(),
			AccountType: user["account_type"].String(),
		},
	}

	return
}

func (p *Profile) AuthProfile(ctx context.Context, req *v1.RegularProfileReq) (res *v1.AuthProfileRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	md := g.Model("users")
	user, err := md.Where(`email = ?`, req.Email).Where(`account_type = ?`, consts.AuthAccountType).One()
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return
	}

	if user.IsEmpty() {
		r.Response.WriteJson(g.Map{
			"code":    consts.ErrCodeAccountNotExist,
			"message": consts.ErrMsgAccountNotExist.Error(),
		})
		return
	}

	var authUserInfo v1.AuthUserInfo

	_ = json.Unmarshal([]byte(user["auth_user_info"].String()), &authUserInfo)
	res = &v1.AuthProfileRes{
		ProfileRes: v1.ProfileRes{
			Email:       user["email"].String(),
			Username:    user["username"].String(),
			AccountType: user["account_type"].String(),
		},
		AuthProvider:    user["auth_provider"].String(),
		AuthUserInfo:    authUserInfo,
		AuthAccessToken: user["auth_access_token"].String(),
	}
	return
}

func (p *Profile) AdminProfile(ctx context.Context, req *v1.RegularProfileReq) (res *v1.AdminProfileRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	md := g.Model("users")
	user, err := md.Where(`email = ?`, req.Email).Where(`account_type = ?`, consts.AdminAccountType).One()
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return
	}

	if user.IsEmpty() {
		r.Response.WriteJson(g.Map{
			"code":    consts.ErrCodeAccountNotExist,
			"message": consts.ErrMsgAccountNotExist.Error(),
		})
		return
	}

	res = &v1.AdminProfileRes{
		ProfileRes: v1.ProfileRes{
			Email:       user["email"].String(),
			Username:    user["username"].String(),
			AccountType: user["account_type"].String(),
		},
	}
	return
}
