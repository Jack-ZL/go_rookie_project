package handler

import (
	"context"
	goodsProto "goRookie_project/server/proto/goods"
)

type GoodsServer struct {
}

func (g GoodsServer) ListGoods(ctx context.Context, req *goodsProto.ListGoodsReq) (*goodsProto.ListGoodsRsp, error) {
	return &goodsProto.ListGoodsRsp{
		Code: 200,
		Msg:  "success",
		List: nil,
	}, nil
}
