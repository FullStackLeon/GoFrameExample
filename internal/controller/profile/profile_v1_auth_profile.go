package profile

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"GoFrameExample/api/profile/v1"
)

func (c *ControllerV1) AuthProfile(ctx context.Context, req *v1.AuthProfileReq) (res *v1.AuthProfileRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
