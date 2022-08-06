package main

import (
	"fmt"
	"neptune-go/src/zinx.game.mmo/apis"
	"neptune-go/src/zinx.game.mmo/core"
	"neptune-go/src/zinx/ziface"
	"neptune-go/src/zinx/znet"
)

const (
	Pid = "pid"
)

func OnConnectionCreate(conn ziface.IConnection) {
	fmt.Println(conn)
	// 1. 创建玩家对象
	player := core.NewPlayer(conn)
	// 2. 发送同步消息
	player.SyncPlayerId()
	// 3. 发送广播消息
	player.BroadCastStartPosition()
	// 4. 记录玩家
	core.WorldObject.AddNewPlayerToWorld(player)
	// 5. 绑定属性
	player.Conn.SetConnectionProperty(Pid, player.Pid)

	fmt.Println("player pid -> : ", player.Pid)
}

func main() {
	// 1. 获取服务器对象
	zinx := znet.NewServer()
	// 2. 注册钩子和销毁函数
	zinx.SetOnConnStart(OnConnectionCreate)
	// 3. 注册处理器
	zinx.AddRouter(2, &apis.WorldChatHandler{})
	// 4. 启动服务器
	zinx.Serve()

}
