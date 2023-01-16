package game

import (
	"ebitenGame/ui/elements"
)

func New() *elements.UI {
	var m elements.UI

	L1 := elements.NewDefaultLabel(0, 100, "PLAYER HP")
	L2 := elements.NewDefaultLabel(0, 150, "BASE HP")

	m.Labels = append(m.Labels, L1)
	m.Labels = append(m.Labels, L2)
	return &m
}
