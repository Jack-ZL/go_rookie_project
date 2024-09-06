package handler

import (
	"github.com/Jack-ZL/go_rookie/rpc"
	adminProto "goRookie_project/server/proto/admin"
	userProto "goRookie_project/server/proto/user"
	"google.golang.org/grpc"
)

func RegisterServiceHandler(server *rpc.G2rpcServer) {
	server.Register(func(g *grpc.Server) {
		adminProto.RegisterAdminServiceServer(g, &AdminServer{})
		userProto.RegisterUserServiceServer(g, &UserServer{})
	})

}
