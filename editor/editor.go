package editor

import (
	"ebitenGame/camera"
	c "ebitenGame/config"
	"ebitenGame/location"
	"ebitenGame/tiles"
	u "ebitenGame/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"strconv"
	"time"
)

type Editor struct {
	SelectedTile int
}

type Mouse struct {
	X int
	Y int
}

var (
	mouse Mouse
)

func New() *Editor {
	return &Editor{
		SelectedTile: 20,
	}
}

func (e *Editor) Update(l *location.Location, cam *camera.Camera) *location.Location {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		time.Sleep(10000000)
		e.SelectedTile -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		time.Sleep(10000000)
		e.SelectedTile += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		l.SaveToFile("location/", "test.loc")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		l.LoadFromFile("location/", "test.loc")
	}
	mouse.X, mouse.Y = ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x := (mouse.X + int(cam.X)) / 32
		y := (mouse.Y + int(cam.Y)) / 32
		if u.InRange(0, c.LocationSize, x) && u.InRange(0, c.LocationSize, y) {
			l.Field[x][y].SpriteNumber = e.SelectedTile
		}
	}
	return l
}

func (e *Editor) Draw(screen *ebiten.Image) {

	for i := 0; i < tiles.Count; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i*32-e.SelectedTile*32), 0)
		screen.DrawImage(tiles.ByNumber(i), op)
	}
	ebitenutil.DebugPrint(screen, "\n"+strconv.Itoa(e.SelectedTile))
}
