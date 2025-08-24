package common

import "github.com/med-sidmed/snake-game/entity"

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	GridSize     = 20
)

var (
	DirUp    = entity.Point{X: 0, Y: -GridSize}
	DirDown  = entity.Point{X: 0, Y: GridSize}
	DirLeft  = entity.Point{X: -GridSize, Y: 0}
	DirRight = entity.Point{X: GridSize, Y: 0}
)
