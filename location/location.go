package location

import (
	"ebitenGame/camera"
	"ebitenGame/config"
	log "ebitenGame/loger"
	"ebitenGame/tiles"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"os"
	"strconv"
	s "strings"
)

type Location struct {
	Width  int
	Height int
	Field  [][]tiles.Tile
}

func New(width int, height int) *Location {
	var l Location
	l.Width = width
	l.Height = height
	if l.Width*tiles.Size > config.ScreenWidth || l.Height*tiles.Size > config.ScreenHeight {
		logrus.Warn("size of number of tiles is bigger than screen sizes")
	}
	l.Field = make([][]tiles.Tile, width)
	for i := 0; i < height; i++ {
		l.Field[i] = make([]tiles.Tile, height)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			random := rand.Intn(50)
			if random < 1 || random > 9 {
				l.Field[i][j].SpriteNumber = random
			} else {
				l.Field[i][j].SpriteNumber = 0
			}

		}
	}
	if config.LogLevel == logrus.Level(6) {
		l.DebugPrint()
	}
	return &l
}

func (l Location) Draw(screen *ebiten.Image, cam *camera.Camera) {
	step := tiles.Size
	tiles.InitTiles()
	for i := 0; i < l.Height; i++ {
		for j := 0; j < l.Width; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*step)-cam.X, float64(j*step)-cam.Y)
			screen.DrawImage(tiles.ByNumber(l.Field[i][j].SpriteNumber), op)
		}
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
	}
}

func (l Location) DebugPrint() {
	for i := 0; i < l.Height; i++ {
		for j := 0; j < l.Width; j++ {
			fmt.Print(l.Field[j][i].SpriteNumber)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (l *Location) SaveToFile(path string, filename string) {
	file, err := os.Create(path + filename)
	if err != nil {
		log.L.Error("Unable to create file")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.L.Error(err)
		}
	}()
	for i := 0; i < l.Height; i++ {
		for j := 0; j < l.Width; j++ {
			_, err = file.WriteString(strconv.Itoa(l.Field[j][i].SpriteNumber) + ",")
			if err != nil {
				log.L.Error(err)
			}
		}
		_, err = file.WriteString("\n")
		if err != nil {
			log.L.Error(err)
		}
	}
	log.L.Info("Location is saved")
}
func (l *Location) LoadFromFile(path string, filename string) *Location {
	file, err := os.Open(path + filename)
	if err != nil {
		log.L.Error("Unable to open file")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.L.Error(err)
		}
	}()
	var res []string
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		res = append(res, string(data[:n]))
	}
	m := s.Join(res, "")
	m = s.ReplaceAll(m, ",", " ")
	m = s.ReplaceAll(m, "\n", "")
	n := s.Split(m, " ")

	for i := 0; i < l.Height; i++ {
		for j := 0; j < l.Width; j++ {
			l.Field[j][i].SpriteNumber, err = strconv.Atoi(n[j+(i*config.LocationSize)])
		}
	}
	log.L.Info("Location is loaded")
	return l
}
