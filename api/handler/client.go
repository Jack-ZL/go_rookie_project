package handler

import (
	"github.com/Jack-ZL/go_rookie/rpc"
	adminProto "goRookie_project/server/proto/admin"
)

var (
	adminClient adminProto.AdminServiceClient
)

func RegisterClients(con *rpc.G2rpcClient) {
	adminClient = adminProto.NewAdminServiceClient(con.Conn)
}
