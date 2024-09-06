package main

import (
	"github.com/Jack-ZL/go_rookie"
	"github.com/Jack-ZL/go_rookie/rpc"
	"goRookie_project/server/handler"
)

func main() {
	engine := go_rookie.Default()
	// 监听9101端口的客户端
	server, err := rpc.NewG2rpcServer(":9101")
	if err != nil {
		panic(err)
	}
	// 注册服务
	handler.RegisterServiceHandler(server)
	err = server.Run()
	if err != nil {
		panic(err)
	}
	engine.Run(":9102")
}
