package handler

import (
	"github.com/Jack-ZL/go_rookie/rpc"
)

var (
// gclient goodsProto.GoodsApiClient
)

func RegisterClient() {
	config := rpc.DefaultG2rpcClientConfig()
	config.Address = "localhost:9001"
	//rpcClient, _ := rpc.NewG2rpcClient(config)
	//defer rpcClient.Conn.Close()
	//
	//gclient = goodsProto.NewGoodsApiClient(rpcClient.Conn)
}
