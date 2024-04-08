// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package profile

import (
	"context"

	"GoFrameExample/api/profile/v1"
)

type IProfileV1 interface {
	RegularProfile(ctx context.Context, req *v1.RegularProfileReq) (res *v1.RegularProfileRes, err error)
	AuthProfile(ctx context.Context, req *v1.AuthProfileReq) (res *v1.AuthProfileRes, err error)
	AdminProfile(ctx context.Context, req *v1.AdminProfileReq) (res *v1.AdminProfileRes, err error)
}
