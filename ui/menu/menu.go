package menu

import (
	"ebitenGame/globals"
	"ebitenGame/ui/elements"
	"golang.org/x/image/colornames"
)

func GAME()        { globals.UI_MODE = "GAME" }
func MULTIPLAYER() { globals.UI_MODE = "MULTIPLAYER" }
func SETTINGS()    { globals.UI_MODE = "SETTINGS" }
func EDITOR()      { globals.UI_MODE = "EDITOR" }
func EXIT()        { globals.UI_MODE = "EXIT" }

func New() *elements.UI {
	var m elements.UI
	PlayBtn := elements.NewButton(500, 100, float64(globals.ScreenWidth/2-250), 0,
		"Play", colornames.Palegreen, colornames.Black, GAME)
	MultiplayerBtn := elements.NewButton(500, 100, float64(globals.ScreenWidth/2-250), 150,
		"Multiplayer", colornames.Palegreen, colornames.Black, MULTIPLAYER)
	SettingsBtn := elements.NewButton(500, 100, float64(globals.ScreenWidth/2-250), 300,
		"Settings", colornames.Palegreen, colornames.Black, SETTINGS)
	EditorBtn := elements.NewButton(500, 100, float64(globals.ScreenWidth/2-250), 450,
		"Editor", colornames.Palegreen, colornames.Black, EDITOR)
	ExitBtn := elements.NewButton(500, 100, float64(globals.ScreenWidth/2-250), 600,
		"Exit", colornames.Palegreen, colornames.Black, EXIT)

	m.Buttons = append(m.Buttons, PlayBtn)
	m.Buttons = append(m.Buttons, MultiplayerBtn)
	m.Buttons = append(m.Buttons, SettingsBtn)
	m.Buttons = append(m.Buttons, EditorBtn)
	m.Buttons = append(m.Buttons, ExitBtn)
	return &m
}
