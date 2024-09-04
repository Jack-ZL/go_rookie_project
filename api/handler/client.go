package handler

import (
	goodsProto "goRookie_project/server/proto/goods"
	"google.golang.org/grpc"
)

var (
	goodsClient goodsProto.GoodsServiceClient
)

func RegisterClients(con *grpc.ClientConn) {
	goodsClient = goodsProto.NewGoodsServiceClient(con)
}
