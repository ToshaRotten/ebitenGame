package elements

import (
	l "ebitenGame/loger"
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"io"
)

//go:embed font.ttf
var ui embed.FS
var (
	FontSize = 18
	FileName = "font.ttf"
)

type UI struct {
	Buttons []*Button
	Labels  []*Label
	Inputs  []*Input
	Scrolls []*Scroll
}

func (m *UI) Draw(screen *ebiten.Image) {
	for i := 0; i < len(m.Buttons); i++ {
		m.Buttons[i].Draw(screen)
	}
	for i := 0; i < len(m.Labels); i++ {
		m.Labels[i].Draw(screen)
	}
}

type Button struct {
	Width           int
	Height          int
	X               float64
	Y               float64
	Text            string
	Color           color.RGBA
	BackgroundColor color.RGBA
	FontFace        font.Face
	Visible         bool
	OnClick         func()
}

func NewButton(width int, height int, x float64, y float64, text string, bgcolor color.RGBA, color color.RGBA, onclick func()) *Button {
	return &Button{
		Width:           width,
		Height:          height,
		X:               x,
		Y:               y,
		Text:            text,
		Visible:         true,
		FontFace:        GetFont(),
		BackgroundColor: bgcolor,
		Color:           color,
		OnClick:         onclick,
	}
}
func (b *Button) Draw(screen *ebiten.Image) {
	if !b.Visible {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.X), float64(b.Y))
	ebitenutil.DrawRect(screen, float64(b.X), float64(b.Y), float64(b.Width), float64(b.Height), b.BackgroundColor)

	posX := int(b.Width/2) + int(b.X) - FontSize*(len(b.Text)-1)
	posY := int(b.Height/2) + int(b.Y) + FontSize
	text.Draw(screen, b.Text, b.FontFace, posX, posY, b.Color)
}

type Label struct {
	Width           int
	Height          int
	X               float64
	Y               float64
	Text            string
	Color           color.RGBA
	BackgroundColor color.RGBA
	FontFace        font.Face
	Visible         bool
	FontSize        int
}

func NewLabel(width int, height int, x float64, y float64, text string, fontsize int, color color.RGBA) *Label {
	return &Label{
		Width:    width,
		Height:   height,
		X:        x,
		Y:        y,
		Text:     text,
		Visible:  true,
		FontFace: GetFont(),
		Color:    color,
		FontSize: fontsize,
	}
}
func NewDefaultLabel(x float64, y float64, text string) *Label {
	return &Label{
		Width:    500,
		Height:   500,
		X:        x,
		Y:        y,
		Text:     text,
		Visible:  true,
		FontFace: GetFont(),
		Color:    color.RGBA{255, 255, 255, 255},
	}
}
func (l *Label) Draw(screen *ebiten.Image) {
	if !l.Visible {
		return
	}
	posX := int(l.X)
	posY := int(l.Y)
	text.Draw(screen, l.Text, l.FontFace, posX, posY, l.Color)
}

type Input struct {
	Width           int
	Height          int
	X               float64
	Y               float64
	Text            string
	Color           color.RGBA
	BackgroundColor color.RGBA
	FontFace        font.Face
	Visible         bool
	OnClick         func()
}

type Scroll struct {
	X               float64
	Y               float64
	Width           float64
	Height          float64
	BackgroundColor color.RGBA
	Items           []*ebiten.Image
	Selector        int
	Count           int
	OnClick         func()
}

func NewScroll() *Scroll {
	return &Scroll{}
}

func (s *Scroll) Update() *Scroll {
	_, wy := ebiten.Wheel()
	if wy > 0 {
		s.Selector += 1
	}
	if wy < 0 {
		s.Selector -= 1
	}
	return s
}

func (s *Scroll) Draw(screen *ebiten.Image) {
	for i := 0; i < s.Count; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i*32-s.Count*32), 0)
		screen.DrawImage(s.Items[i], op)
	}
}

func NewInput(width int, height int, x float64, y float64, text string, bgcolor color.RGBA, color color.RGBA, onclick func()) *Input {
	return &Input{
		Width:           width,
		Height:          height,
		X:               x,
		Y:               y,
		Text:            text,
		Visible:         true,
		FontFace:        GetFont(),
		Color:           color,
		BackgroundColor: bgcolor,
		OnClick:         onclick,
	}
}
func (i *Input) EnterText() {

}
func (i *Input) Draw(screen *ebiten.Image) {
	if !i.Visible {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(i.X), float64(i.Y))
	ebitenutil.DrawRect(screen, float64(i.X), float64(i.Y), float64(i.Width), float64(i.Height), i.BackgroundColor)
	posX := int(i.X)
	posY := int(i.Y)
	text.Draw(screen, i.Text, i.FontFace, posX, posY, i.Color)
}

func GetFont() font.Face {
	r, err := ui.Open(FileName)
	if err != nil {
		l.L.Error(err)
	}

	fontData, err := io.ReadAll(r)
	if err != nil {
		l.L.Error(err)
	}
	tt, err := opentype.Parse(fontData)
	if err != nil {
		l.L.Error(err)
	}
	fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(FontSize),
		DPI:     96,
		Hinting: font.HintingFull,
	})
	if err != nil {
		l.L.Error(err)
	}
	return fontFace
}
