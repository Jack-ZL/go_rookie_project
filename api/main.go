package main

import (
	"github.com/Jack-ZL/go_rookie"
	"github.com/Jack-ZL/go_rookie/rpc"
	"goRookie_project/api/handler"
)

func main() {
	engine := go_rookie.Default()
	config := rpc.DefaultG2rpcClientConfig()
	config.Address = ":9101" // 起一个9101的客户端
	rpcClient, _ := rpc.NewG2rpcClient(config)
	defer rpcClient.Conn.Close()

	handler.RegisterClients(rpcClient)

	group := engine.Group("admin")
	group.Post("/login", handler.AdminLogin)

	//router.InitRoutes(engine)

	engine.Run(":9103")
}
