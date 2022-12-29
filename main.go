package main

import (
	"ebitenGame/camera"
	"ebitenGame/config"
	"ebitenGame/location"
	"ebitenGame/loger"
	"ebitenGame/player"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct{}

var (
	loc = location.New(10, 10)
	cam = camera.New()
	pl  = player.New()
)

func (g *Game) Update() error {
	loger.L.Trace("Player choords: ", pl.X, pl.Y)
	pl.Movment()
	cam.UpdateCamera()
	loger.L.Trace("Camera: ", cam.X, cam.Y, cam.Scale)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	loc.Draw(screen, cam)
	pl.Draw(screen, cam)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeigh
}

func main() {
	loger.ConfigLoger()
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeigh)
	ebiten.SetWindowTitle("")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
