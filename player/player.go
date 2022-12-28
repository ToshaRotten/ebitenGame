package player

import (
	"ebitenGame/camera"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
)

type Player struct {
	X     int
	Y     int
	Image *ebiten.Image
}

func New() *Player {
	var err error
	var pl Player
	pl.Image, _, err = ebitenutil.NewImageFromFile("sprites/player.png")
	if err != nil {
		logrus.Fatal(err)
	}
	return &pl
}

func (p *Player) Movment() *Player {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= 5
	}
	return p
}

func (p *Player) Draw(screen *ebiten.Image, cam *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X)-float64(cam.X), float64(p.Y)-float64(cam.Y))
	screen.DrawImage(p.Image, op)
}
