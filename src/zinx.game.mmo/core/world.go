package core

import (
	"fmt"
	"sync"
)

type World struct {
	// 兴趣点地图
	aoi *AreaOfInterest
	// 所有在线玩家
	onlinePlayers map[int]*Player
	// 保护集合的锁
	wordLock sync.RWMutex
}

var WorldObject *World

// 全局初始化
func init() {
	WorldObject = &World{
		aoi:           NewAreaOfInterest(AoiMinX, AoiMaxX, AoiMinY, AoiMaxY, AoiCntX, AoiCntY),
		onlinePlayers: make(map[int]*Player),
	}
}

func (world *World) AddNewPlayerToWorld(player *Player) {
	if player == nil {
		fmt.Println("player is nil")
		return
	}
	world.wordLock.Lock()
	defer world.wordLock.Unlock()
	world.onlinePlayers[player.Pid] = player
	world.aoi.AddNewPlayerToGridByPosition(player.Position.X, player.Position.Z, player.Pid)
}

func (world *World) RemovePlayerFromWorld(pid int) {
	if pid < 0 {
		fmt.Println("pid is illegal")
		return
	}
	// TODO 考虑加锁的方式
	world.wordLock.Lock()
	defer world.wordLock.Unlock()
	player, ok := world.onlinePlayers[pid]
	if !ok {
		fmt.Println("player doesn't exist")
		return
	}
	world.aoi.RemoverPlayerFromGridByPosition(player.Position.X, player.Position.Z, pid)
	delete(world.onlinePlayers, pid)
}

func (world *World) GetPlayerByPid(pid int) (player *Player) {
	if pid < 0 {
		fmt.Println("pid is illegal")
		return
	}
	world.wordLock.RLock()
	defer world.wordLock.RUnlock()
	return world.onlinePlayers[pid]
}

func (world *World) GetOnlinePlayers() (players []*Player) {
	world.wordLock.RLock()
	defer world.wordLock.RUnlock()

	for _, player := range world.onlinePlayers {
		players = append(players, player)
	}
	return players
}
