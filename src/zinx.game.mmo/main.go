package main

import (
	"fmt"
	"neptune-go/src/zinx.game.mmo/core"
	"neptune-go/src/zinx/ziface"
	"neptune-go/src/zinx/znet"
)

func OnConnectionCreate(conn ziface.IConnection) {
	fmt.Println(conn)
	// 1. 创建玩家对象
	player := core.NewPlayer(conn)
	// 2. 发送同步消息
	player.SyncPlayerId()
	// 3. 发送广播消息
	player.BroadCastStartPosition()
	fmt.Println("player pid -> : ", player.Pid)
}

func main() {
	// 1. 获取服务器对象
	zinx := znet.NewServer()
	// 2. 注册钩子和销毁函数
	zinx.SetOnConnStart(OnConnectionCreate)
	// 3. 注册处理器
	// 4. 启动服务器
	zinx.Serve()

}
