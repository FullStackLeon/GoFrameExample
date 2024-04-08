package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gsession"

	"GoFrameExample/internal/consts"
	"GoFrameExample/utility"
)

// LoginCheck 中间件是登录检测，接口未传入提示未授权，如果用户
func LoginCheck(req *ghttp.Request) {
	SessionKey := req.Cookie.Get(consts.SessionKey)
	storage := gsession.NewStorageRedis(g.Redis())
	sessionMap, err := storage.GetSession(req.Context(), SessionKey.String(), time.Hour*24)
	if err != nil {
		req.Response.WriteStatus(http.StatusInternalServerError)
		return
	}

	if sessionMap == nil {
		req.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	jwtTokenVar := sessionMap.GetVar(consts.JWTToken)
	if jwtTokenVar != nil {
		token, err := utility.VerifyToken(jwtTokenVar.String())
		if err != nil {
			req.Response.WriteStatus(http.StatusForbidden)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		req.SetParam("email", claims["email"])
		req.SetParam("account_type", claims["account_type"])
	}
	req.Middleware.Next()
}

// RouterCheck 中间件作用是路由检查，每个角色可以访问的路径存储在consts.RouterMap中
func RouterCheck(req *ghttp.Request) {
	accountType := req.GetParam("account_type")
	if accountType == nil {
		req.Response.WriteStatus(http.StatusForbidden)
		return
	}
	for _, v := range consts.RouterMap[accountType.String()] {
		if req.URL.Path == v {
			req.Middleware.Next()
			return
		}
	}

	req.Response.WriteStatus(http.StatusForbidden)
	return
}
