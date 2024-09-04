package handler

import (
	"github.com/Jack-ZL/go_rookie/rpc"
	goodsProto "goRookie_project/server/proto/goods"
	loginProto "goRookie_project/server/proto/login"
	"google.golang.org/grpc"
)

func RegisterServiceHandler(server *rpc.G2rpcServer) {
	server.Register(func(g *grpc.Server) {
		goodsProto.RegisterGoodsServiceServer(g, &GoodsServer{})
		loginProto.RegisterLoginServiceServer(g, &LoginServer{})
	})

}
