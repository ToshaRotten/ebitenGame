package base

import (
	G "ebitenGame/globals"
	"ebitenGame/ui/elements"
	u "ebitenGame/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Mouse struct {
	X int
	Y int
}

func CREATE() {}
func EXIT()   { G.UI_MODE = "EXIT" }

func MenuMouseCheck(m *elements.UI) {
	G.MOUSE.X, G.MOUSE.Y = ebiten.CursorPosition()
	for i := 0; i < len(m.Buttons); i++ {
		if u.InRange(int(m.Buttons[i].X), int(m.Buttons[i].X)+m.Buttons[i].Width, G.MOUSE.X) &&
			u.InRange(int(m.Buttons[i].Y), int(m.Buttons[i].Y)+m.Buttons[i].Height, G.MOUSE.Y) {
			m.Buttons[i].BackgroundColor = colornames.Azure
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				m.Buttons[i].OnClick()
			}
		} else {
			m.Buttons[i].BackgroundColor = colornames.Palegreen
		}
	}

}
func New() *elements.UI {
	var m elements.UI
	CreateBtn := elements.NewButton(500, 100, float64(G.ScreenWidth/2-250), 0,
		"Play", colornames.Palegreen, colornames.Black, CREATE)
	ExitBtn := elements.NewButton(500, 100, float64(G.ScreenWidth/2-250), 150,
		"Exit", colornames.Palegreen, colornames.Black, EXIT)
	m.Buttons = append(m.Buttons, CreateBtn)
	m.Buttons = append(m.Buttons, ExitBtn)
	return &m
}
