package debug

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

func Rect(img *ebiten.Image, x float64, y float64) {
	ebitenutil.DrawRect(img, x, y, 32, 32, colornames.Red50)
}
