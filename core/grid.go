package core

import (
	"fmt"
	"sync"
)

/**
aoi 格子
*/

type Grid struct {
	Gid int
	// 格子左边界
	MinX int
	// 格子右边界
	MaxX int
	// 格子上边界
	MinY int
	// 格子下边界
	MaxY int

	playerIds map[int]bool
	pidLock   sync.Mutex
}

// 格子的构造函数
func NewGrid(gid int, minX int, maxX int, minY int, maxY int) *Grid {
	return &Grid{
		Gid:       gid,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIds: make(map[int]bool),
	}
}

// 添加玩家
func (g *Grid) Add(playerId int) {
	g.pidLock.Lock()
	defer g.pidLock.Unlock()
	g.playerIds[playerId] = true
}

// 删除玩家
func (g *Grid) Remove(playerId int) {
	if _, ok := g.playerIds[playerId]; ok {
		delete(g.playerIds, playerId)
	}
}

// 获取当前格子内的所有玩家id
func (g *Grid) GetPlayerIds() (playerIds []int) {
	for k, _ := range g.playerIds {
		playerIds = append(playerIds, k)
	}
	return
}

// 格子的基本属性显示
func (g *Grid) String() string {
	return fmt.Sprintf("gid: %d, minX: %d, maxX: %d, minY: %d, maxY: %d", g.Gid, g.MinX, g.MaxX, g.MinY, g.MaxY)
}
