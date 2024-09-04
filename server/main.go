package main

import (
	"github.com/Jack-ZL/go_rookie"
	"github.com/Jack-ZL/go_rookie/rpc"
	"goRookie_project/server/handler"
)

func main() {
	engine := go_rookie.Default()

	server, err := rpc.NewG2rpcServer(":9001")
	if err != nil {
		panic(err)
	}
	handler.RegisterServiceHandler(server)
	server.Run()
	engine.Run(":9002")
}
