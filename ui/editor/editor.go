package editor

import (
	"ebitenGame/globals"
	"ebitenGame/ui/elements"
	"golang.org/x/image/colornames"
)

func CREATE_UNIT() {}
func MENU()        { globals.UI_MODE = "MENU" }
func SAVE() {
	//location.SaveToFile("location/", "test.loc")
}
func LOAD() {
	//location.LoadFromFile("location/", "test.loc")
}

func New() *elements.UI {
	var m elements.UI

	Scroll := elements.NewScroll()

	L1 := elements.NewDefaultLabel(0, 100, "CTRL + S to save")
	L2 := elements.NewDefaultLabel(0, 150, "CTRL + L to load")
	ReturnButton := elements.NewButton(150, 100, float64(globals.ScreenWidth-150), float64(globals.ScreenHeight-150),
		"Back", colornames.Palegreen, colornames.Black, MENU)
	SaveButton := elements.NewButton(150, 100, 0, float64(globals.ScreenHeight-150),
		"Save", colornames.Palegreen, colornames.Black, SAVE)
	LoadButton := elements.NewButton(150, 100, 0, float64(globals.ScreenHeight-300),
		"Load", colornames.Palegreen, colornames.Black, LOAD)

	m.Scrolls = append(m.Scrolls, Scroll)
	m.Buttons = append(m.Buttons, ReturnButton)
	m.Buttons = append(m.Buttons, SaveButton)
	m.Buttons = append(m.Buttons, LoadButton)
	m.Labels = append(m.Labels, L1)
	m.Labels = append(m.Labels, L2)
	return &m
}
