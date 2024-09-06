package main

import (
	"github.com/Jack-ZL/go_rookie"
	"github.com/Jack-ZL/go_rookie/rpc"
	"goRookie_project/server/handler"
)

func main() {
	engine := go_rookie.Default()

	server, err := rpc.NewG2rpcServer(":9101")
	if err != nil {
		panic(err)
	}
	handler.RegisterServiceHandler(server)
	err = server.Run()
	if err != nil {
		panic(err)
	}
	engine.Run(":9102")
}
