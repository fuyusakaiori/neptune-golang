package core

import (
	"fmt"
	"sync"
)

// Grid 兴趣点
type Grid struct {
	// 兴趣点 ID
	gid int
	// 兴趣点左边边界
	minX float32
	// 兴趣点右边边界
	maxX float32
	// 兴趣点上边边界
	minY float32
	// 兴趣点下边边界
	maxY float32
	// 兴趣点内玩家数量: key-玩家 ID, value-玩家是否存在
	players map[int]bool
	// 保护兴趣点的锁
	gridLock sync.RWMutex
}

func (grid *Grid) GetGridGid() int {
	return grid.gid
}

// GetPlayerIdList 获取兴趣点内所有玩家
func (grid *Grid) GetPlayerIdList() (players []int) {
	// 1. 获取读锁
	grid.gridLock.RLock()
	defer grid.gridLock.RUnlock()
	// 2. 获取所有玩家
	for player, _ := range grid.players {
		players = append(players, player)
	}
	return players
}

// AddNewPlayerToGrid 向兴趣点中添加玩家
func (grid *Grid) AddNewPlayerToGrid(playerID int) {
	// 1. 获取写锁
	grid.gridLock.Lock()
	defer grid.gridLock.Unlock()
	// 2. 添加新玩家
	grid.players[playerID] = true
}

// RemovePlayerFromGrid 将玩家从兴趣点中移除
func (grid *Grid) RemovePlayerFromGrid(playerID int) {
	// 1. 获取写锁
	grid.gridLock.Lock()
	defer grid.gridLock.Unlock()
	// 2. 判断玩家是否存在
	if _, ok := grid.players[playerID]; !ok {
		fmt.Println("this player doesn't exist")
	}
	// 3. 移除玩家
	delete(grid.players, playerID)
}

// NewGrid 创建新的兴趣点
func NewGrid(gid int, minX float32, maxX float32, minY float32, maxY float32) (grid *Grid) {
	grid = &Grid{
		gid:     gid,
		minX:    minX,
		maxX:    maxX,
		minY:    minY,
		maxY:    maxY,
		players: make(map[int]bool),
	}
	return grid
}

func (grid *Grid) String() string {
	return fmt.Sprintf("grid id: %d, minX: %f, maxX: %f, minY: %f, maxY: %f, players: %v",
		grid.gid, grid.minX, grid.maxX, grid.minY, grid.maxY, grid.players)
}
