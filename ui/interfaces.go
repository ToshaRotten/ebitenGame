package ui

import (
	G "ebitenGame/globals"
	"ebitenGame/ui/base"
	"ebitenGame/ui/editor"
	"ebitenGame/ui/elements"
	"ebitenGame/ui/game"
	"ebitenGame/ui/menu"
	u "ebitenGame/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

var (
	MENU   = menu.New()
	EDITOR = editor.New()
	BASE   = base.New()
	GAME   = game.New()
)

func MouseCheck(m *elements.UI) {
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

func UpdateScroll(e *elements.UI) {
	G.MOUSE.X, G.MOUSE.Y = ebiten.CursorPosition()
	for i := 0; i < len(e.Scrolls); i++ {
		if u.InRange(int(e.Scrolls[i].X), int(e.Scrolls[i].X+e.Scrolls[i].Width), G.MOUSE.X) &&
			u.InRange(int(e.Scrolls[i].Y), int(e.Scrolls[i].Y+e.Scrolls[i].Height), G.MOUSE.Y) {
			e.Scrolls[i].BackgroundColor = colornames.Azure
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				e.Scrolls[i].OnClick()
				e.Scrolls[i].Update()
			}
		} else {
			e.Scrolls[i].BackgroundColor = colornames.Palegreen
		}
	}
}
