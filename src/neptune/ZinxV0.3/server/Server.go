package main

import (
	"neptune-go/src/neptune/ZinxV0.3/server/router"
	"neptune-go/src/zinx/znet"
)

func main() {
	// 1. 创建服务器对象
	server := znet.NewServer("[ZinxV0.3]")
	// 2. 调用服务器方法
	server.AddRouter(&router.PingRouter{})
	server.Serve()
}
