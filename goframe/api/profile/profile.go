// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package profile

import (
	"context"

	"GoFrameExample/api/profile/v1"
)

type IProfileV1 interface {
	Regular(ctx context.Context, req *v1.RegularReq) (res *v1.RegularRes, err error)
	Auth(ctx context.Context, req *v1.AuthReq) (res *v1.AuthRes, err error)
	Admin(ctx context.Context, req *v1.AdminReq) (res *v1.AdminRes, err error)
}
