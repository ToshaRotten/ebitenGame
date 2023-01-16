package globals

import (
	"github.com/sirupsen/logrus"
)

type Mouse struct {
	X  int
	Y  int
	WX float64
	WY float64
}

var (
	LocationSize = 50
	ScreenWidth  = 1280
	ScreenHeight = 720
	LogLevel     = logrus.Level(5)
	ObjectDebug  = 0
	TileSize     = 32

	MOUSE     Mouse
	UI_MODE   = "MENU"
	GAME_MODE = ""
)
