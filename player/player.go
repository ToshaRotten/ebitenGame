package player

import (
	"ebitenGame/camera"
	"ebitenGame/loger"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
	"math"
)

type Player struct {
	X      float64
	Y      float64
	Rotate float64
	Image  *ebiten.Image
}

type Mouse struct {
	X int
	Y int
}

var (
	mouse Mouse
)

func New() *Player {
	var err error
	var pl Player
	pl.Image, _, err = ebitenutil.NewImageFromFile("sprites/player.png")
	if err != nil {
		logrus.Fatal(err)
	}
	return &pl
}

func (p *Player) Movement() *Player {
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

func (p *Player) MovementToMouse(cam *camera.Camera) *Player {
	mouse.X, mouse.Y = ebiten.CursorPosition()
	relationX := float64(mouse.X) - (p.X - cam.X)
	relationY := float64(mouse.Y) - (p.Y - cam.Y)

	angle := (180 / math.Pi) * -math.Atan2(relationY, relationX)
	loger.L.Trace(angle)
	p.Rotate = math.Atan2(relationY, relationX)
	loger.L.Trace(p.Rotate)
	return p
}

func (p *Player) Draw(screen *ebiten.Image, cam *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(p.Rotate)
	op.GeoM.Translate(p.X-cam.X, p.Y-cam.Y)
	screen.DrawImage(p.Image, op)
}
