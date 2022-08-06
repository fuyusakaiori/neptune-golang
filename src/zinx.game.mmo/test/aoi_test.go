package test

import (
	"fmt"
	"neptune-go/src/zinx.game.mmo/core"
	"testing"
)

func TestAreaOfInterest(t *testing.T) {

	aoi := core.NewAreaOfInterest(100, 250, 300, 500, 3, 4)

	fmt.Println(aoi)
}

func TestGetSurroundingGridsByGid(t *testing.T) {

	aoi := core.NewAreaOfInterest(0, 250, 0, 250, 5, 5)
	grids := aoi.GetSurroundingGridsByGid(12)
	for _, grid := range grids {
		fmt.Println("grid id: ", grid.GetGridGid())
	}

}
