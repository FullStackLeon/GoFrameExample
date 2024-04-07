package user

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/google/uuid"

	"GoFrameExample/api/user/v1"
	"GoFrameExample/internal/consts"
	"GoFrameExample/internal/model/do"
	"GoFrameExample/utility"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) Register(ctx context.Context, req *v1.RegReq) (res *v1.RegRes, err error) {
	md := g.Model("users")
	user := &do.Users{}
	record, err := md.Where("email = ?", req.Email).One()
	if err != nil {
		return nil, consts.ErrMsgQueryUserFailed
	}

	res = &v1.RegRes{}
	if record != nil {
		return nil, consts.ErrMsgUserRegistered
	}
	encodedPassword := utility.HashStr(req.Password)
	user = &do.Users{
		Username:    req.Nickname,
		Email:       req.Email,
		Password:    encodedPassword,
		AccountType: req.AccountType,
	}

	if req.AccountType == consts.AuthAccountType {
		user.AuthProvider = utility.GetRandomProvider()
		user.AuthUserId = uuid.NewString()
		user.AuthAccessToken = utility.GenerateRandomString(32)
		user.AuthRefreshToken = utility.GenerateRandomString(32)
		user.AuthUserInfo = utility.GenerateRandomUserProfile()
	}
	id, err := md.Where("email = ?", req.Email).InsertAndGetId(user)
	if err != nil {
		return nil, consts.ErrMsgRegisterFailed
	}
	res.Id = uint64(id)
	return
}

func (u *User) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	password := utility.HashStr(req.Password)
	model := g.Model("users")
	user, err := model.Fields("username", "email", "account_type").Where("email = ? and password = ?", req.Email, password).One()
	if err != nil {
		return nil, consts.ErrMsgQueryUserFailed
	}

	if user == nil {
		return nil, consts.ErrMsgAccountNotExist
	}

	token, err := utility.GenerateToken(user)
	if err != nil {
		return nil, consts.ErrMsgGenerateTokenFailed
	}

	// 通过Redis缓存Session登录态，解决单点问题，Session中存储JWT Token
	storage := gsession.NewStorageRedis(g.Redis())
	sessionId := consts.SessionKey + utility.EncodeStr(user["email"].String())
	err = storage.SetSession(ctx, sessionId, gmap.NewStrAnyMapFrom(g.Map{consts.JWTToken: token}), time.Hour*24)
	if err != nil {
		return nil, consts.ErrMsgSetSessionFailed
	}

	// 设置Cookie，将SessionId发送到客户端
	r := ghttp.RequestFromCtx(ctx)
	r.Cookie.Set(consts.SessionKey, sessionId)
	res = &v1.LoginRes{
		Token:    token,
		Username: user["username"].String(),
		Email:    user["email"].String(),
	}
	return
}
