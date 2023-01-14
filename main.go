package main

import (
	"ebitenGame/camera"
	"ebitenGame/config"
	"ebitenGame/editor"
	"ebitenGame/location"
	"ebitenGame/loger"
	"ebitenGame/player"
	"ebitenGame/unit"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct{}

var (
	mode = "MAP"
	edit = editor.New()
	u    = unit.New()
	loc  = location.New(config.LocationSize, config.LocationSize)
	cam  = camera.New()
	pl   = player.New()
)

func (g *Game) Update() error {

	if mode == "MENU" {

	}
	if mode == "MAP" {
		loc = edit.Update(loc, cam)
		loger.L.Trace("Camera: ", cam.X, cam.Y, cam.Scale)

		cam.UpdateCamera()

	}
	if mode == "GAME" {
		loger.L.Trace("Player: ", pl.X, pl.Y, pl.Rotate)
		pl.Movement()
		cam.X = pl.X - float64(config.ScreenWidth)/2
		cam.Y = pl.Y - float64(config.ScreenHeight)/2
		pl.MovementToMouse(cam)
		cam.UpdateCamera()
		loger.L.Trace("Camera: ", cam.X, cam.Y, cam.Scale)
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			loc = location.New(config.LocationSize, config.LocationSize)
			loger.L.Debug("regen")
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if mode == "MAP" {
		loc.Draw(screen, cam)
		edit.Draw(screen)
	}
	if mode == "GAME" {
		loc.Draw(screen, cam)
		u.Draw(screen, cam)
		pl.Draw(screen, cam)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	loger.ConfigLoger()
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
