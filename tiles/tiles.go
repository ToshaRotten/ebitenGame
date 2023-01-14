package tiles

import (
	"bytes"
	"ebitenGame/loger"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"io/ioutil"
	"log"
)

type Tile struct {
	SpriteNumber int
	Image        *ebiten.Image
}

const (
	Size = 32
)

var (
	Count = 50
	Image *ebiten.Image
)

func InitTiles() {
	file, err := ioutil.ReadFile("tiles/src/Tiles.png")
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

func ByNumber(number int) *ebiten.Image {
	if number > Count || number < 0 {
		loger.L.Warning("Wrong tile")
		number = 0
	} else {
		sx, sy := toXY(number)
		tile := Image.SubImage(image.Rect(sx*Size, sy*Size, sx*Size+Size, sy*Size+Size)).(*ebiten.Image)
		return tile
	}
	return nil
}
