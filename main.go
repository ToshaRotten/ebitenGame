package main

import (
	"ebitenGame/camera"
	c "ebitenGame/config"
	"ebitenGame/editor"
	"ebitenGame/location"
	"ebitenGame/loger"
	"ebitenGame/player"
	"ebitenGame/ui"
	"ebitenGame/ui/menu"
	"ebitenGame/unit"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
	"os"
)

type Game struct{}

var (
	edit = editor.New()
	u    = unit.New()
	loc  = location.New(c.LocationSize, c.LocationSize)
	cam  = camera.New()
	pl   = player.New()
	MENU = menu.New()
	L1   = ui.NewDefaultLabel(0, 100, "CTRL + S to save")
	L2   = ui.NewDefaultLabel(0, 150, "CTRL + L to load")
)

func (g *Game) Update() error {
	if menu.Mode == "MENU" {
		menu.MenuMouseCheck(MENU)
	}
	if menu.Mode == "MULTIPLAYER" {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			menu.Mode = "MENU"
		}
	}
	if menu.Mode == "SETTINGS" {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			menu.Mode = "MENU"
		}
	}
	if menu.Mode == "EDITOR" {
		loc = edit.Update(loc, cam)
		loger.L.Trace("Camera: ", cam.X, cam.Y, cam.Scale)
		cam.UpdateCamera()
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			menu.Mode = "MENU"
		}
	}
	if menu.Mode == "START" {
		loger.L.Trace("Player: ", pl.X, pl.Y, pl.Rotate)
		pl.Movement()
		cam.X = pl.X - float64(c.ScreenWidth)/2
		cam.Y = pl.Y - float64(c.ScreenHeight)/2
		pl.MovementToMouse(cam)
		cam.UpdateCamera()
		loger.L.Trace("Camera: ", cam.X, cam.Y, cam.Scale)
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			loc = location.New(c.LocationSize, c.LocationSize)
			loger.L.Debug("regen")
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			menu.Mode = "MENU"
		}
	}
	if menu.Mode == "EXIT" {
		os.Exit(0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if menu.Mode == "MENU" {
		loc.Draw(screen, cam)
		MENU.Draw(screen)
	}
	if menu.Mode == "EDITOR" {
		loc.Draw(screen, cam)
		edit.Draw(screen)
		L1.Draw(screen)
		L2.Draw(screen)
	}
	if menu.Mode == "START" {
		loc.Draw(screen, cam)
		u.Draw(screen, cam)
		pl.Draw(screen, cam)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return c.ScreenWidth, c.ScreenHeight
}

func main() {
	loger.ConfigLoger()
	ebiten.SetWindowSize(c.ScreenWidth, c.ScreenHeight)
	ebiten.SetWindowTitle("")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
