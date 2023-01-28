package objects

import (
	log "ebitenGame/loger"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ObjectType struct {
	Name   string
	Sprite *ebiten.Image
	Color  string
}

func InitTypes() map[string]ObjectType {
	m := make(map[string]ObjectType)
	var spr *ebiten.Image
	var err error
	spr, _, err = ebitenutil.NewImageFromFile("sprites/redTank.png")
	if err != nil {
		log.L.Error(err)
	}
	m["Tank"] = ObjectType{Name: "Tank", Sprite: spr, Color: "RED"}
	spr, _, _ = ebitenutil.NewImageFromFile("sprites/sbase.png")
	if err != nil {
		log.L.Error(err)
	}
	m["SBase"] = ObjectType{Name: "SBase", Sprite: spr, Color: "RED"}
	spr, _, _ = ebitenutil.NewImageFromFile("sprites/sbaseblue.png")
	if err != nil {
		log.L.Error(err)
	}
	m["SBaseBlue"] = ObjectType{Name: "SBaseBlue", Sprite: spr, Color: "BLUE"}

	return m
}
