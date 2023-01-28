package base

import (
	G "ebitenGame/globals"
	"ebitenGame/ui/elements"
	V "ebitenGame/variables"
	"golang.org/x/image/colornames"
)

func CREATE() { V.PL.TakeUnit = true; RETURN() }
func RETURN() { G.UI_MODE = "GAME" }

func New() *elements.UI {
	var m elements.UI
	CreateBtn := elements.NewButton(500, 100, float64(G.ScreenWidth/2-250), 0,
		"Create", colornames.Palegreen, colornames.Black, CREATE)
	ExitBtn := elements.NewButton(500, 100, float64(G.ScreenWidth/2-250), 150,
		"Return", colornames.Palegreen, colornames.Black, RETURN)
	m.Buttons = append(m.Buttons, CreateBtn)
	m.Buttons = append(m.Buttons, ExitBtn)
	return &m
}
