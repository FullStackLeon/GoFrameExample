package profile

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"GoFrameExample/api/profile/v1"
)

func (c *ControllerV1) AdminProfile(ctx context.Context, req *v1.AdminReq) (res *v1.AdminRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
