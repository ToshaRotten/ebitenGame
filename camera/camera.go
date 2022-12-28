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

//
//func (c Camera) UpdateZoom() {
//	var scrollY float64
//	if ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyPageDown) {
//		scrollY = -0.25
//	} else if ebiten.IsKeyPressed(ebiten.KeyE) || ebiten.IsKeyPressed(ebiten.KeyPageUp) {
//		scrollY = .25
//	} else {
//		_, scrollY = ebiten.Wheel()
//		if scrollY < -1 {
//			scrollY = -1
//		} else if scrollY > 1 {
//			scrollY = 1
//		}
//	}
//	c.ScaleTo += scrollY * (c.ScaleTo / 7)
//}
//
//func (c Camera) ClampZoomLevel() {
//	if c.ScaleTo < 0.01 {
//		c.ScaleTo = 0.01
//	} else if c.ScaleTo > 100 {
//		c.Scale = 100
//	}
//}
//
//func (c Camera) SmoothZoom() {
//	div := 10.0
//	if c.ScaleTo > c.Scale {
//		c.Scale += (c.ScaleTo - c.Scale) / div
//	} else if c.ScaleTo < c.Scale {
//		c.Scale -= (c.Scale - c.ScaleTo) / div
//	}
//}
