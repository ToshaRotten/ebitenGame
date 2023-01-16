package menu

import (
	c "ebitenGame/config"
	"ebitenGame/ui"
	u "ebitenGame/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Mouse struct {
	X int
	Y int
}

type Menu struct {
	Buttons []*ui.Button
}

var (
	mouse Mouse
	Mode  = "MENU"
)

func START()       { Mode = "START" }
func MULTIPLAYER() { Mode = "MULTIPLAYER" }
func SETTINGS()    { Mode = "SETTINGS" }
func EDITOR()      { Mode = "EDITOR" }
func EXIT()        { Mode = "EXIT" }

func MenuMouseCheck(m *Menu) {
	mouse.X, mouse.Y = ebiten.CursorPosition()
	//log.L.Debug(mouse.X, mouse.Y)
	//log.L.Debug(u.InRange(int(m.Buttons[1].X), m.Buttons[1].Width, mouse.X) &&
	//	u.InRange(int(m.Buttons[1].Y), m.Buttons[1].Height, mouse.Y))
	for i := 0; i < len(m.Buttons); i++ {
		if u.InRange(int(m.Buttons[i].X), int(m.Buttons[i].X)+m.Buttons[i].Width, mouse.X) &&
			u.InRange(int(m.Buttons[i].Y), int(m.Buttons[i].Y)+m.Buttons[i].Height, mouse.Y) {
			m.Buttons[i].BackgroundColor = colornames.Azure
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				m.Buttons[i].OnClick()
			}
		} else {
			m.Buttons[i].BackgroundColor = colornames.Palegreen
		}
	}

}
func New() *Menu {
	var m Menu
	PlayBtn := ui.NewButton(500, 100, float64(c.ScreenWidth/2-250), 0,
		"Play", colornames.Palegreen, colornames.Black, START)
	MultiplayerBtn := ui.NewButton(500, 100, float64(c.ScreenWidth/2-250), 150,
		"Multiplayer", colornames.Palegreen, colornames.Black, MULTIPLAYER)
	SettingsBtn := ui.NewButton(500, 100, float64(c.ScreenWidth/2-250), 300,
		"Settings", colornames.Palegreen, colornames.Black, SETTINGS)
	EditorBtn := ui.NewButton(500, 100, float64(c.ScreenWidth/2-250), 450,
		"Editor", colornames.Palegreen, colornames.Black, EDITOR)
	ExitBtn := ui.NewButton(500, 100, float64(c.ScreenWidth/2-250), 600,
		"Exit", colornames.Palegreen, colornames.Black, EXIT)
	m.Buttons = append(m.Buttons, PlayBtn)
	m.Buttons = append(m.Buttons, MultiplayerBtn)
	m.Buttons = append(m.Buttons, SettingsBtn)
	m.Buttons = append(m.Buttons, EditorBtn)
	m.Buttons = append(m.Buttons, ExitBtn)
	return &m
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for i := 0; i < len(m.Buttons); i++ {
		m.Buttons[i].Draw(screen)
	}
}
