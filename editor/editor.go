package editor

import (
	"ebitenGame/camera"
	G "ebitenGame/globals"
	"ebitenGame/location"
	"ebitenGame/tiles"
	u "ebitenGame/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Editor struct {
	SelectedTile int
}

var (
	EDITOR = New()
)

func New() *Editor {
	return &Editor{
		SelectedTile: 20,
	}
}

func (e *Editor) Update(l *location.Location, cam *camera.Camera) *location.Location {
	if inpututil.IsKeyJustPressed(ebiten.KeyS) && ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
		l.SaveToFile("location/", "test.loc")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyL) && ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
		l.LoadFromFile("location/", "test.loc")
	}
	G.MOUSE.X, G.MOUSE.Y = ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x := (G.MOUSE.X + int(cam.X)) / 32
		y := (G.MOUSE.Y + int(cam.Y)) / 32
		if u.InRange(0, G.LocationSize-1, x) && u.InRange(0, G.LocationSize-1, y) {
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
}
