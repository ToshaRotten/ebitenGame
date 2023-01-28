package player

import (
	"ebitenGame/camera"
	"ebitenGame/debug"
	c "ebitenGame/globals"
	"ebitenGame/loger"
	"ebitenGame/objects"
	"ebitenGame/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sirupsen/logrus"
	"math"
)

type Player struct {
	X         float64
	Y         float64
	Rotate    float64
	Image     *ebiten.Image
	TakeUnit  bool
	Visible   bool
	Transform bool
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
	var p Player
	p.Visible = true
	p.Image, _, err = ebitenutil.NewImageFromFile("sprites/player.png")
	if err != nil {
		logrus.Fatal(err)
	}
	return &p
}

func (p *Player) Movement() *Player {
	if utils.InRange(0, c.LocationSize*32, int(p.X)) && utils.InRange(0, c.LocationSize*32, int(p.Y)) {
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
	} else {
		p.Y = 100
		p.X = 100
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

func (p *Player) ColisionWithBase(objs *objects.Objects) bool {
	for i := 0; i < len(objs.Instances); i++ {
		if objs.Instances[i].Type.Name == "SBase" {
			if utils.InRange(int(p.X), int(p.X+32), int(objs.Instances[i].X)) && utils.InRange(int(p.Y), int(p.Y+32), int(objs.Instances[i].Y)) {
				return true
			}
		}
	}
	return false
}

func (p *Player) ColisionWithAny(objs *objects.Objects) bool {
	for i := 0; i < len(objs.Instances); i++ {
		if utils.InRange(int(p.X), int(p.X+32), int(objs.Instances[i].X)) && utils.InRange(int(p.Y), int(p.Y+32), int(objs.Instances[i].Y)) {
			return true
		}
	}
	return false
}

func (p *Player) ColisionWithInstance(inst *objects.Instance) bool {
	if utils.InRange(int(p.X), int(p.X+32), int(inst.X)) && utils.InRange(int(p.Y), int(p.Y+32), int(inst.Y)) {
		return true
	}
	return false
}

func (p *Player) PlaceUnit() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && p.TakeUnit == true {
		return true
	}
	return false
}

func (p *Player) Draw(screen *ebiten.Image, cam *camera.Camera) {
	if !p.Visible {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(p.Rotate)
	op.GeoM.Translate(p.X-cam.X, p.Y-cam.Y)
	if c.ObjectDebug == 1 {
		debug.Rect(p.Image, p.X, p.Y)
	}
	screen.DrawImage(p.Image, op)
}
