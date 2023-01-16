package main

import (
	"ebitenGame/camera"
	G "ebitenGame/globals"
	"ebitenGame/location"
	"ebitenGame/loger"
	"ebitenGame/player"
	GUI "ebitenGame/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
)

type Game struct{}

var (
	LOC = location.New(G.LocationSize, G.LocationSize)
	CAM = camera.New()
	PL  = player.New()
)

func (g *Game) Update() error {
	if G.UI_MODE == "MENU" {
		GUI.MouseCheck(GUI.MENU)
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
	}

	if G.UI_MODE == "GAME" {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			G.UI_MODE = "MENU"
		}

	}
	if G.UI_MODE == "EXIT" {
		os.Exit(0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if G.UI_MODE == "MENU" {
		GUI.MENU.Draw(screen)
	}
	if G.UI_MODE == "EDITOR" {
		GUI.EDITOR.Draw(screen)
	}
	if G.UI_MODE == "GAME" {
		LOC.Draw(screen, CAM)
		PL.Draw(screen, CAM)
		GUI.GAME.Draw(screen)
	}
	if G.UI_MODE == "BASE" {

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return G.ScreenWidth, G.ScreenHeight
}

func main() {
	loger.ConfigLoger()
	ebiten.SetWindowSize(G.ScreenWidth, G.ScreenHeight)
	ebiten.SetWindowTitle("")
	if err := ebiten.RunGame(&Game{}); err != nil {
		loger.L.Error(err)
	}
}
