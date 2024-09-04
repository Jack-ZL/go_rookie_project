package main

import (
	"github.com/Jack-ZL/go_rookie/rpc"
	"goRookie_project/api/handler"
	"goRookie_project/api/router"
)

func main() {
	config := rpc.DefaultG2rpcClientConfig()
	config.Address = ":9001"
	rpcClient, _ := rpc.NewG2rpcClient(config)
	rpcClient.Conn.Close()

	gr := router.InitRoutes()
	handler.RegisterClients(rpcClient.Conn)

	gr.Run(":9003")
}
