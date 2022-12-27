package main

import (
	"ebitenGame/config"
	"ebitenGame/location"
	"ebitenGame/player"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
	"log"
)

type Game struct{}

var (
	Log = logrus.New()
	Loc = location.New(5, 5)
	pl  = player.New()
)

func (g *Game) Update() error {
	Log.Trace("Player choords: ", pl.X, pl.Y)
	pl.Movment()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	pl.Draw(screen)
	//Loc.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeigh
}

func main() {
	Log.Level = 5
	//Log.Trace("6")
	//Log.Debug("5")
	//Log.Info("4")
	//Log.Warning("3")
	//Log.Error("2")
	//Log.Fatal("1")
	//Log.Panic("0")

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
