package core

import "fmt"

/*
	aoi 兴趣点管理模块
*/
type AOIManager struct {
	// x轴水平格子的数量
	CountsX int
	// 区域的左边界坐标
	MinX int
	// 区域的有边界坐标
	MaxX int

	// y轴的格子数量
	CountsY int
	// 区域的上边界坐标
	MinY int
	// 区域的下边界坐标
	MaxY int

	grids map[int]*Grid
}

// 构造函数
func NewAOIManager(countsX int, minX int, maxX int, countsY int, minY int, maxY int) *AOIManager {
	m := &AOIManager{
		CountsX: countsX,
		MinX:    minX,
		MaxX:    maxX,
		CountsY: countsY,
		MinY:    minY,
		MaxY:    maxY,
		grids:   make(map[int]*Grid),
	}

	for y := 0; y < countsY; y++ {
		for x := 0; x < countsX; x++ {
			gid := y*countsX + x
			m.grids[gid] = NewGrid(gid,
				m.MinX+m.gridWidth()*x,
				m.MinX+m.gridWidth()*(x+1),
				m.MinY+m.gridHeight()*y,
				m.MinY+m.gridHeight()*(y+1),
			)
		}
	}
	return m
}

// 每个格子的宽度
func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CountsX
}

// 每个格子的高度
func (m *AOIManager) gridHeight() int {
	return (m.MaxY - m.MinY) / m.CountsY
}

// 打印的数据
func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager:\nminX:%d, CountsX:%d, maxX:%d, countsY:%d, minY:%d, maxY:%d\n", m.CountsX, m.MinX, m.MaxX, m.CountsY, m.MinY, m.MaxY)
	fmt.Printf(s)
	for _, g := range m.grids {
		s += fmt.Sprintln(g)
	}
	return s
}

// 根据gid得到周边九宫格格子的信息
func (m *AOIManager) GetSurroundGirdsByGid(gid int) []*Grid {
	grids := make([]*Grid, 0)
	// 如果不在m.Grids里，直接返回空
	if _, ok := m.grids[gid]; !ok {
		return grids
	}

	// 自己加入到九宫格中
	grids = append(grids, m.grids[gid])

	idx := gid % m.CountsX
	// 左边是否有格子
	if idx > 0 {
		grids = append(grids, m.grids[idx-1])
	}
	// 右边是否有格子
	if idx < m.CountsX-1 {
		grids = append(grids, m.grids[idx+1])
	}

	gidList := make([]int, 0)
	for _, g := range grids {
		gidList = append(gidList, g.Gid)
	}

	for _, v := range gidList {
		idy := v / m.CountsX
		if idy > 0 {
			// 上方的格子
			grids = append(grids, m.grids[v-m.CountsX])
		}
		if idy < m.CountsY-1 {
			// 下方的格子
			grids = append(grids, m.grids[v+m.CountsX])
		}
	}

	return grids
}

// 根据玩家的x，y获取指定的格子编号
func (m *AOIManager) GetGidByPos(x, y float32) int {
	idx := (int(x) - m.MinX) / m.gridWidth()
	idy := (int(y) - m.MinY) / m.gridHeight()
	return idy*m.CountsX + idx
}

// 根据玩家的x，y坐标，获取九宫格内的所有玩家id
func (m *AOIManager) GetPlayerIdsByPos(x, y float32) (playerIds []int) {
	gid := m.GetGidByPos(x, y)
	for _, grid := range m.GetSurroundGirdsByGid(gid) {
		playerIds = append(playerIds, grid.GetPlayerIds()...)
	}
	return
}

// 添加玩家到一个格子中去
func (m *AOIManager) AddPlayerIdToGrid(playerId, gid int) {
	m.grids[gid].Add(playerId)
}

// 移除玩家格子中的玩家id
func (m *AOIManager) RemovePlayerIdFromGrid(playerId, gid int) {
	m.grids[gid].Remove(playerId)
}

// 根据玩家坐标，把玩家添加到对应格子中去
func (m *AOIManager) AddPlayerIdToGridByPos(playerId int, x, y float32) {
	m.AddPlayerIdToGrid(playerId, m.GetGidByPos(x, y))
}

// 根据玩家的坐标，把玩家从指定的格子中移除
func (m *AOIManager) RemovePlayerIdFromGridByPos(playerId int, x, y float32) {
	m.RemovePlayerIdFromGrid(playerId, m.GetGidByPos(x, y))
}
