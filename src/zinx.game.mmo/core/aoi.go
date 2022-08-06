package core

import "fmt"

const (
	AoiMinX float32 = 85
	AoiMaxX float32 = 410
	AoiCntX int     = 10
	AoiMinY float32 = 75
	AoiMaxY float32 = 400
	AoiCntY int     = 20
)

// AreaOfInterest 兴趣点地图
type AreaOfInterest struct {
	// 地图左边界
	AreaMinX float32
	// 地图右边界
	AreaMaxX float32
	// 地图上边界
	AreaMinY float32
	// 地图下边界
	AreaMaxY float32
	// 每行兴趣点的数量
	CntX int
	// 每列兴趣点的数量
	CntY int
	// 兴趣点
	grids map[int]*Grid
}

// 获取兴趣点长度
func (aoi *AreaOfInterest) getPointLength() float32 {
	return (aoi.AreaMaxX - aoi.AreaMinX) / float32(aoi.CntX)
}

// 获取每个兴趣点宽度
func (aoi *AreaOfInterest) getPointWidth() float32 {
	return (aoi.AreaMaxY - aoi.AreaMinY) / float32(aoi.CntY)
}

// 根据坐标获取格子
func (aoi AreaOfInterest) getGidOfGridByPosition(x, y float32) (gid int) {
	return int(y/aoi.getPointLength())*aoi.CntX + int(x/aoi.getPointWidth())
}

// GetSurroundingGridsByGid  获取周边九宫格
func (aoi AreaOfInterest) GetSurroundingGridsByGid(gid int) (grids []*Grid) {
	// 1. 判断需要获取周边格子的自身是否存在
	if _, ok := aoi.grids[gid]; !ok {
		fmt.Println("this gid doesn't exist")
		return grids
	}
	// 2. 判断左右两侧是否越界
	idx := gid % aoi.CntX
	if idx > 0 {
		grids = append(grids, aoi.grids[gid-1])
	}
	if idx < aoi.CntX-1 {
		grids = append(grids, aoi.grids[gid+1])
	}
	// 3. 将自己添加到集合中
	grids = append(grids, aoi.grids[gid])
	// 4. 获取所有水平集合的 ID
	gidList := make([]int, 0, len(grids))
	for _, grid := range grids {
		gidList = append(gidList, grid.gid)
	}
	// 5. 遍历水平集合看上下是否存在
	for _, gid := range gidList {
		idy := gid / aoi.CntY
		// 5. 判断上下边界是否越界
		if idy > 0 {
			grids = append(grids, aoi.grids[gid-aoi.CntX])
		}
		if idy < aoi.CntY-1 {
			grids = append(grids, aoi.grids[gid+aoi.CntX])
		}
	}
	return grids
}

// GetSurroundingPlayersByPosition 通过玩家的坐标获取周围所有玩家
func (aoi *AreaOfInterest) GetSurroundingPlayersByPosition(x float32, y float32) (players []int) {
	// 1. 根据坐标计算得到所在的格子
	gid := aoi.getGidOfGridByPosition(x, y)
	// 2. 根据所在的格子获取周围的格子
	grids := aoi.GetSurroundingGridsByGid(gid)
	// 3. 获取每个格子中的玩家
	for _, grid := range grids {
		players = append(players, grid.GetPlayerIdList()...)
	}
	return players
}

func (aoi *AreaOfInterest) AddNewPlayerToGridByPid(pid, gid int) {
	// 1. 校验传入的 ID 是否合法
	if pid < 0 || gid < 0 {
		fmt.Println("pid or gid is illegal")
		return
	}
	// 2. 添加到格子中
	aoi.grids[gid].AddNewPlayerToGrid(pid)
}

func (aoi *AreaOfInterest) RemovePlayerFromGridByPid(pid, gid int) {
	// 1. 校验传入的 ID 是否合法
	if pid < 0 || gid < 0 {
		fmt.Println("pid or gid is illegal")
		return
	}
	// 2. 移除格子中的玩家
	aoi.grids[gid].RemovePlayerFromGrid(pid)
}

func (aoi AreaOfInterest) AddNewPlayerToGridByPosition(x, y float32, pid int) {
	// 1. 校验传入的坐标是否合法
	if x < 0 || y < 0 || pid < 0 {
		fmt.Println(" x or y or pid is illegal")
		return
	}
	// 2. 坐标转换为格子所属 ID
	gid := aoi.getGidOfGridByPosition(x, y)
	// 3. 添加玩家到格子中
	aoi.AddNewPlayerToGridByPid(pid, gid)
}

func (aoi AreaOfInterest) RemoverPlayerFromGridByPosition(x, y float32, pid int) {
	if x < 0 || y < 0 || pid < 0 {
		fmt.Println(" x or y or pid is illegal")
		return
	}
	gid := aoi.getGidOfGridByPosition(x, y)
	aoi.RemovePlayerFromGridByPid(pid, gid)
}

func (aoi AreaOfInterest) GetPlayersInGridByGid(gid int) (players []int) {
	if gid < 0 {
		fmt.Println("gid is illegal")
		return players
	}
	return aoi.grids[gid].GetPlayerIdList()
}

func (aoi *AreaOfInterest) String() string {
	str := fmt.Sprintf("aoi: \nAreaMinX: %f, AreaMaxX: %f, CntX: %d, AreaMinY: %f, AreaMaxY: %f, CountOfCol: %d\n",
		aoi.AreaMinX, aoi.AreaMaxX, aoi.CntX, aoi.AreaMinY, aoi.AreaMaxY, aoi.CntY)
	for _, grid := range aoi.grids {
		str += fmt.Sprintln(grid)
	}
	return str
}

func NewAreaOfInterest(minX float32, maxX float32, minY float32, maxY float32, cntX int, cntY int) *AreaOfInterest {
	// 1. 初始化兴趣地图
	area := &AreaOfInterest{
		AreaMinX: minX,
		AreaMaxX: maxX,
		AreaMinY: minY,
		AreaMaxY: maxY,
		CntY:     cntY,
		CntX:     cntX,
		grids:    make(map[int]*Grid),
	}
	// 2. 获取每个兴趣点的长度和宽度
	width := area.getPointWidth()
	length := area.getPointLength()
	// 3. 初始化每个兴趣点的编号
	for y := 0; y < area.CntY; y++ {
		for x := 0; x < area.CntX; x++ {
			// 2.1, 计算兴趣点编号
			pid := y*area.CntX + x
			// 2.2 初始化兴趣点范围
			area.grids[pid] = NewGrid(pid,
				area.AreaMinX+float32(x)*width,
				area.AreaMinX+float32(x+1)*width,
				area.AreaMinY+float32(y)*length,
				area.AreaMinY+float32(y+1)*length)
		}
	}

	return area
}
