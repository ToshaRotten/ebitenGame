package main

import (
	G "ebitenGame/globals"
	"ebitenGame/loger"
	"ebitenGame/objects"
	GUI "ebitenGame/ui"
	V "ebitenGame/variables"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
)

type Game struct{}

func (g *Game) Update() error {
	if G.UI_MODE == "MENU" {
		GUI.MouseCheck(GUI.MENU)
	}

	if G.UI_MODE == "BASE" {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			G.UI_MODE = "MENU"
		}
		GUI.MouseCheck(GUI.BASE)
	}

	if G.UI_MODE == "MULTIPLAYER" {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			G.UI_MODE = "MENU"
		}
	}

	if G.UI_MODE == "SETTINGS" {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			G.UI_MODE = "MENU"
		}

	}

	if G.UI_MODE == "EDITOR" {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			G.UI_MODE = "MENU"
		}
		GUI.MouseCheck(GUI.EDITOR)
		GUI.UpdateScroll(GUI.EDITOR)
		V.CAM.UpdateCamera()
	}

	if G.UI_MODE == "GAME" {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			G.UI_MODE = "MENU"
		}
		if V.PL.ColisionWithInstance(V.BASE) {
			G.UI_MODE = "BASE"
			V.PL.X += 64
			V.PL.Y += 64
		}
		if V.PL.PlaceUnit() {
			temp := objects.NewInstance(V.TYPES["Tank"])
			temp.X = V.PL.X
			temp.Y = V.PL.Y
			V.UNITS.Add(temp)
			V.PL.TakeUnit = false
		}
		V.PL.MovementToMouse(V.CAM)
		V.PL.Movement()
		V.CAM.X = V.PL.X - float64(G.ScreenWidth/2)
		V.CAM.Y = V.PL.Y - float64(G.ScreenHeight/2)
	}
	if G.UI_MODE == "EXIT" {
		os.Exit(0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if G.UI_MODE == "MENU" {
		V.LOC.Draw(screen, V.CAM)
		GUI.MENU.Draw(screen)
	}
	if G.UI_MODE == "EDITOR" {
		V.LOC.Draw(screen, V.CAM)
		GUI.EDITOR.Draw(screen)
	}
	if G.UI_MODE == "GAME" {
		V.LOC.Draw(screen, V.CAM)
		V.PL.Draw(screen, V.CAM)
		V.BASE.Draw(screen, V.CAM)
		V.ENEMY_BASE.Draw(screen, V.CAM)
		V.UNITS.Draw(screen, V.CAM)
		GUI.GAME.Draw(screen)
	}
	if G.UI_MODE == "BASE" {
		V.LOC.Draw(screen, V.CAM)
		V.PL.Draw(screen, V.CAM)
		V.BASE.Draw(screen, V.CAM)
		V.ENEMY_BASE.Draw(screen, V.CAM)
		V.UNITS.Draw(screen, V.CAM)
		GUI.BASE.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return G.ScreenWidth, G.ScreenHeight
}

func main() {
	loger.ConfigLoger()
	ebiten.SetWindowSize(G.ScreenWidth, G.ScreenHeight)
	ebiten.SetWindowTitle("")
	V.InitVars()
	if err := ebiten.RunGame(&Game{}); err != nil {
		loger.L.Error(err)
	}
}
