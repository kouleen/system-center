package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/kouleen/common/bootstrap"
	"github.com/kouleen/common/middleware"
	"github.com/kouleen/idl/kitex_gen/rpc"
	"github.com/kouleen/idl/kitex_gen/system/systemservice"
)

func main() {
	// etcd注册中心
	bootstrap.Run(rpc.SYSTEM_RPC_SERVER, func(option ...server.Option) server.Server {
		return systemservice.NewServer(new(SystemServiceImpl), option...)
	},
		bootstrap.WithServerMiddleware(middleware.RpcServerMiddleware),
	)
}
