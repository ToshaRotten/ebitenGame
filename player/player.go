package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sirupsen/logrus"
)

type Player struct {
	X     int
	Y     int
	Image *ebiten.Image
}

func New() Player {
	var err error
	var pl Player
	pl.Image, _, err = ebitenutil.NewImageFromFile("sprites/player.png")
	if err != nil {
		logrus.Fatal(err)
	}
	return pl
}

func (p Player) Movment() Player {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		p.Y += 5
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		p.Y -= 5
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		p.X += 5
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		p.X -= 5
	}
	return p
}

func (p Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(p.Image, op)
}
