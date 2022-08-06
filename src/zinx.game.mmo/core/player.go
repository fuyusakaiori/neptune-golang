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

// Talk 发送消息给所有的玩家
func (player *Player) Talk(content string) {
	// 1. 准备发送的消息
	message := &pb.BroadCast{
		Pid: int32(player.Pid),
		Tp:  1,
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}
	// 2. 获取所有玩家
	players := WorldObject.GetOnlinePlayers()
	// 3. 发送消息
	for _, p := range players {
		p.SendMessage(200, message)
	}
}

// SyncSurrounding 将新加入的玩家同步给其他玩家
func (player *Player) SyncSurrounding() {
	// 1. 获取周边的玩家 ID
	pids := WorldObject.aoi.GetSurroundingPlayersByPosition(player.Position.X, player.Position.Z)
	// 2. 获取周边玩家
	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WorldObject.GetPlayerByPid(pid))
	}
	// 3. 发送消息
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
	// 4. 发送给周边玩家
	for _, player := range players {
		player.SendMessage(200, message)
	}
}

func (player *Player) SyncPlayers() {
	// 1. 获取周边的玩家 ID
	pids := WorldObject.aoi.GetSurroundingPlayersByPosition(player.Position.X, player.Position.Z)
	// 2. 获取周边玩家
	players := make([]*pb.Player, 0, len(pids))
	for _, pid := range pids {
		p := WorldObject.GetPlayerByPid(pid)
		players = append(players, &pb.Player{
			Pid: int32(p.Pid),
			P: &pb.Position{
				X: p.Position.X,
				Y: p.Position.Y,
				Z: p.Position.Z,
				V: p.Position.V,
			},
		})
	}
	// 3. 准备要发送的消息
	message := &pb.SyncPlayers{
		Players: players,
	}
	// 4. 发送
	player.SendMessage(202, message)
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
