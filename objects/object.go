package objects

import (
	"ebitenGame/camera"
	"ebitenGame/utils"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Objects struct {
	Instances []*Instance
}

type InstanceI interface {
	New()
	Create()
	Destroy()
	ColisionWith()
}

type Instance struct {
	ID   int
	HP   int
	X    float64
	Y    float64
	Type ObjectType
}

var (
	ObjectCounter = 0
)

func NewObjects() *Objects {
	var o Objects
	return &o
}

func (o *Objects) Draw(screen *ebiten.Image, cam *camera.Camera) {
	for i := 0; i < len(o.Instances); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(o.Instances[i].X-cam.X, o.Instances[i].Y-cam.Y)
		screen.DrawImage(o.Instances[i].Type.Sprite, op)
	}
}

func (o *Objects) Add(instance *Instance) {
	fmt.Println(instance)
	o.Instances = append(o.Instances, instance)
}

func NewInstance(Type ObjectType) *Instance {
	ObjectCounter++
	return &Instance{
		ID:   ObjectCounter,
		X:    0,
		Y:    0,
		HP:   100,
		Type: Type,
	}
}

func (o *Instance) CollisionWith(secondObject *Instance) bool {
	if utils.InRange(int(o.X), int(o.X+32), int(secondObject.X)) && utils.InRange(int(o.Y), int(o.Y+32), int(secondObject.Y)) {
		return true
	}
	return false
}

func (o *Instance) Destroy() {

}

func (o *Objects) GetByType(t ObjectType) []*Instance {
	var inst []*Instance
	for i := 0; i < len(o.Instances); i++ {
		if o.Instances[i].Type == t {
			inst = append(inst, o.Instances[i])
		}
	}
	return inst
}

func (o *Instance) Draw(screen *ebiten.Image, cam *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.X-cam.X, o.Y-cam.Y)
	screen.DrawImage(o.Type.Sprite, op)
}
