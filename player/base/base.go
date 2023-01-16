package base

import (
	"ebitenGame/camera"
	log "ebitenGame/loger"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SmallBase struct {
	X        float64
	Y        float64
	IsPlayer bool
	Image    *ebiten.Image
}

type MainBase struct {
	HP         int
	X          float64
	Y          float64
	IsWorking  bool
	IsAttacked bool
	IsPlayer   bool
	Image      *ebiten.Image
}

func NewMainBase(player bool) *MainBase {
	var b MainBase
	var err error
	b.HP = 100
	b.X = 0
	b.Y = 0
	b.IsWorking = false
	b.IsAttacked = false
	b.IsPlayer = player
	if b.IsPlayer {
		b.Image, _, err = ebitenutil.NewImageFromFile("sprites/sbase.png")
	}
	if err != nil {
		log.L.Error(err)
	}
	return &b
}

func NewSmallBase(player bool) *SmallBase {
	var b SmallBase
	var err error
	b.X = 0
	b.Y = 0
	b.IsPlayer = player
	if b.IsPlayer {
		b.Image, _, err = ebitenutil.NewImageFromFile("sprites/sbase.png")
	} else {
		b.Image, _, err = ebitenutil.NewImageFromFile("sprites/sbaseblue.png")
	}
	if err != nil {
		log.L.Error(err)
	}
	return &b
}

func (b *MainBase) Draw(screen *ebiten.Image, cam *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X-cam.X, b.Y-cam.Y)
	screen.DrawImage(b.Image, op)
}

func (b *SmallBase) Draw(screen *ebiten.Image, cam *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X-cam.X, b.Y-cam.Y)
	screen.DrawImage(b.Image, op)
}
