package unit

import (
	"ebitenGame/camera"
	"ebitenGame/config"
	"ebitenGame/debug"
	"ebitenGame/loger"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type UnitType struct {
	name   string
	Sprite *ebiten.Image
}

type Unit struct {
	HP   int
	X    float64
	Y    float64
	Type UnitType
}

func New() *Unit {
	var u Unit
	var err error
	u.Type.Sprite, _, err = ebitenutil.NewImageFromFile("sprites/redTank.png")
	if err != nil {
		loger.L.Error(err)
	}
	u.HP = 100
	u.X = 0
	u.Y = 0

	return &u
}

func (u *Unit) Movement() *Unit {
	return u
}

func (u *Unit) Draw(screen *ebiten.Image, cam *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(u.X-cam.X, u.Y-cam.Y)
	if config.ObjectDebug == 1 {
		debug.Rect(u.Type.Sprite, u.X, u.Y)
	}
	screen.DrawImage(u.Type.Sprite, op)
}
