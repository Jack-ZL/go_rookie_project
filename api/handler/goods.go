package handler

import (
	"context"
	"github.com/Jack-ZL/go_rookie"
	"goRookie_project/common"
	goodsProto "goRookie_project/server/proto/goods"
	"net/http"
)

func ListGoods(c *go_rookie.Context) {
	ctx := context.Background()
	req := &goodsProto.ListGoodsReq{}
	rsp, err := goodsClient.ListGoods(ctx, req)
	if err != nil {
		c.Fail(int(rsp.Code), rsp.Msg)
		return
	}

	common.RefineNil(rsp)
	c.JSON(http.StatusOK, rsp)
}
