package profile

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

func (p *Profile) Regular(ctx context.Context, req *v1.RegularReq) (res *v1.RegularRes, err error) {
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

	res = &v1.RegularRes{
		ProfileRes: v1.ProfileRes{
			Email:       user["email"].String(),
			Username:    user["username"].String(),
			AccountType: user["account_type"].String(),
		},
	}

	return
}

func (p *Profile) Auth(ctx context.Context, req *v1.AuthReq) (res *v1.AuthRes, err error) {
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
	res = &v1.AuthRes{
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

func (p *Profile) Admin(ctx context.Context, req *v1.AdminReq) (res *v1.AdminRes, err error) {
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

	res = &v1.AdminRes{
		ProfileRes: v1.ProfileRes{
			Email:       user["email"].String(),
			Username:    user["username"].String(),
			AccountType: user["account_type"].String(),
		},
	}
	return
}
