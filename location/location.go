package location

import (
	"ebitenGame/config"
	"ebitenGame/tiles"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
)

type Cell struct {
	Sprite int
}

type Location struct {
	Width  int
	Height int
	Field  [][]Cell
}

func New(width int, height int) Location {
	var loc Location
	loc.Width = width
	loc.Height = height

	if loc.Width*tiles.Size > config.ScreenWidth || loc.Height*tiles.Size > config.ScreenHeigh {
		logrus.Warn("size of number of tiles is bigger than screen sizes")
	}

	loc.Field = make([][]Cell, width)
	for i := 0; i < height; i++ {
		loc.Field[i] = make([]Cell, height)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			loc.Field[i][j].Sprite = 0
			//loc.Field[i][j].Sprite = rand.Intn(20)
		}
	}
	return loc
}

func (l Location) Draw(screen *ebiten.Image) {
	step := tiles.Size
	tiles.InitTiles()
	for i := 0; i < l.Height; i++ {
		for j := 0; j < l.Width; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*step), float64(j*step))
			//time.Sleep(100000000)
			screen.DrawImage(tiles.Tile(l.Field[i][j].Sprite), op)
		}
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
	}
}
