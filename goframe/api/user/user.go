// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"GoFrameExample/api/user/v1"
)

type IUserV1 interface {
	Reg(ctx context.Context, req *v1.RegReq) (res *v1.RegRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
}
