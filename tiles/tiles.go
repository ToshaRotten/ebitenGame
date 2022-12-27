package tiles

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"io/ioutil"
	"log"
)

const (
	Size = 32
)

var (
	Image *ebiten.Image
)

func InitTiles() {
	file, err := ioutil.ReadFile("tiles/src/locationTileset.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}
	Image = ebiten.NewImageFromImage(img)
}

func toXY(number int) (int, int) {
	y := number / 10
	x := number % 10
	return x, y
}

func Tile(number int) *ebiten.Image {
	sx, sy := toXY(number)
	tile := Image.SubImage(image.Rect(sx*Size, sy*Size, sx*Size+Size, sy*Size+Size)).(*ebiten.Image)
	return tile
}
