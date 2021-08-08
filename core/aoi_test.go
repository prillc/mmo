package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	aoiMgr := NewAOIManager(5, 0, 250, 5, 0, 250)
	fmt.Println(aoiMgr)
}

func TestAOIManager_GetSurroundGirdsByGid(t *testing.T) {
	aoiMgr := NewAOIManager(5, 0, 250, 5, 0, 250)
	for gid, _ := range aoiMgr.grids {
		grids := aoiMgr.GetSurroundGirdsByGid(gid)
		gids := make([]int, 0)
		for _, v := range grids {
			gids = append(gids, v.Gid)
		}
		fmt.Println(gid, gids)
	}
}
