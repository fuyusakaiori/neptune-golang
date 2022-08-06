package core

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"neptune-go/src/zinx.game.mmo/protobuf/pb"
	"neptune-go/src/zinx/ziface"
	"sync"
)

type Player struct {
	Pid      int
	Conn     ziface.IConnection
	Position *Position
}

var PidGen int = 1
var IdLock sync.RWMutex

// SendMessage 发送数据给客户端
func (player *Player) SendMessage(msgId uint32, message proto.Message) {
	// 1. 序列化
	data, err := proto.Marshal(message)
	if err != nil {
		fmt.Println("serialize message occur err", err)
		return
	}
	// 2. 发送给服务器处理
	if player.Conn == nil {
		fmt.Println("player conn is nil")
		return
	}
	if err := player.Conn.SendMessage(msgId, data); err != nil {
		fmt.Println("player conn send message err", err)
		return
	}

}

// SyncPlayerId 同步玩家 ID
func (player *Player) SyncPlayerId() {
	// 1. 准备发送的消息
	message := &pb.SyncPid{
		Pid: int32(player.Pid),
	}
	// 2. 发送消息
	player.SendMessage(1, message)
}

// BroadCastStartPosition 广播消息
func (player *Player) BroadCastStartPosition() {
	// 1. 准备发送的消息
	message := &pb.BroadCast{
		Pid: int32(player.Pid),
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: player.Position.X,
				Y: player.Position.Y,
				Z: player.Position.Z,
				V: player.Position.V,
			},
		},
	}
	// 2. 发送消息
	player.SendMessage(200, message)
}

// NewPlayer 创建玩家
func NewPlayer(conn ziface.IConnection) *Player {
	// 1. 生成玩家 ID
	IdLock.Lock()
	pid := PidGen
	PidGen++
	IdLock.Unlock()
	// 2. 创建玩家位置
	position := NewPosition()
	// 3. 创建玩家对象
	player := &Player{
		Pid:      pid,
		Conn:     conn,
		Position: position,
	}
	return player
}
