package main

import (
	_ "image/jpeg"
	"log"

	"flag"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"io/ioutil"
	"time"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

var (
	ebitenImage *ebiten.Image
	textObject  font.Face
	endTime     time.Time
	sec         int
	min         int
)

func textObj() {
	f, err := ebitenutil.OpenFile("./HigashiOme-Gothic-C-1.3.ttf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := truetype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 36

	textObject = truetype.NewFace(tt, &truetype.Options{
		Size:    512,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	t := endTime.Sub(time.Now())
	sec = int(t / 1000 / 1000 / 1000 % 60)
	min = int(t / 1000 / 1000 / 1000 / 60)
	// const layout = "15:04:05"

	screen.Fill(color.NRGBA{0xff, 0xff, 0xff, 0xff})

	op := &ebiten.DrawImageOptions{}
	// op.GeoM.Scale(4, 4)
	// op.GeoM.Translate(64, 64)
	screen.DrawImage(ebitenImage, op)
	text.Draw(screen, fmt.Sprintf("%02d:%02d", min, sec), textObject, 600, 600, color.NRGBA{0xe3, 0x49, 0x39, 0xff})

	time.Sleep(100 * time.Millisecond)
	return nil
}

func main() {
	m := flag.Int("m", 0, "minutes")
	path := flag.String("p", "./test.jpg", "image path")
	flag.Parse()
	var err error
	ebitenImage, _, err = ebitenutil.NewImageFromFile(*path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	textObj()
	endTime = time.Now()
	endTime = endTime.Add(time.Duration(*m) * time.Minute)
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Filter (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
