package handler

import (
	"context"
	"github.com/Jack-ZL/go_rookie"
	adminProto "goRookie_project/server/proto/admin"
	"net/http"
)

func AdminLogin(c *go_rookie.Context) {
	ctx := context.Background()
	req := &adminProto.AdminLoginReq{}
	err := c.BindJson(req)
	if err != nil {
		c.Fail(http.StatusBadRequest, err.Error())
		return
	}

	rsp, err := adminClient.AdminLogin(ctx, req)
	if err != nil {
		c.Fail(int(rsp.Code), rsp.Msg)
		return
	}

	//common.RefineNil(rsp)
	c.JSON(http.StatusOK, rsp)
}
