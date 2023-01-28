package variables

import (
	"ebitenGame/camera"
	G "ebitenGame/globals"
	"ebitenGame/location"
	"ebitenGame/objects"
	"ebitenGame/player"
)

var (
	LOC        = location.New(G.LocationSize, G.LocationSize)
	CAM        = camera.New()
	PL         = player.New()
	TYPES      = objects.InitTypes()
	UNITS      = objects.NewObjects()
	BASE       = objects.NewInstance(TYPES["SBase"])
	ENEMY_BASE = objects.NewInstance(TYPES["SBaseBlue"])
)

func InitVars() {
	BASE.X = 128
	BASE.Y = 128
	ENEMY_BASE.X = float64(G.ScreenWidth)
	ENEMY_BASE.Y = float64(G.ScreenWidth)
}
