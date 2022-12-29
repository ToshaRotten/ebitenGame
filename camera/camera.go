package camera

import (
	"ebitenGame/loger"
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	X     float64
	Y     float64
	Scale float64
}

func New() *Camera {
	return &Camera{
		X: 0,
		Y: 0,
	}
}

func (c *Camera) UpdateCamera() *Camera {
	loger.L.Trace("Update camera pos")
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		c.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		c.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		c.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		c.Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyPageDown) {
		c.Scale -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyPageUp) {
		c.Scale += 5
	}
	return c
}
