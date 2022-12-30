package location

import (
	"ebitenGame/camera"
	"ebitenGame/config"
	"ebitenGame/tiles"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
	"math/rand"
)

type Tile struct {
	Sprite int
}

type Location struct {
	Width  int
	Height int
	Field  [][]Tile
}

func New(width int, height int) Location {
	var l Location
	l.Width = width
	l.Height = height
	if l.Width*tiles.Size > config.ScreenWidth || l.Height*tiles.Size > config.ScreenHeight {
		logrus.Warn("size of number of tiles is bigger than screen sizes")
	}
	l.Field = make([][]Tile, width)
	for i := 0; i < height; i++ {
		l.Field[i] = make([]Tile, height)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			l.Field[i][j].Sprite = rand.Intn(50)
		}
	}
	return l
}

func (l Location) Draw(screen *ebiten.Image, cam *camera.Camera) {
	step := tiles.Size
	tiles.InitTiles()
	for i := 0; i < l.Height; i++ {
		for j := 0; j < l.Width; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*step)-cam.X, float64(j*step)-cam.Y)
			//time.Sleep(100000000)
			screen.DrawImage(tiles.Tile(l.Field[i][j].Sprite), op)
		}
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
	}
}

func (l Location) DebugPrint() {
	for i := 0; i < l.Height; i++ {
		for j := 0; i < l.Width; j++ {
			fmt.Print(l.Field[i][j].Sprite)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
