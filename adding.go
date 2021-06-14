package main

import (
	"log"
	"time"
	"image/color"
	
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"math/rand"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 640
	screenHeight = 480

	dotCount = 4096
)

var (
	runnerImage *ebiten.Image
)

type Dot struct {
	x, y, dx, dy float64
}

func (d *Dot) Init() {
	d.x = rand.Float64() * float64(screenWidth) * 1.2
	d.y = rand.Float64() * float64(screenHeight) * 1.2
	d.dx = 0
	d.dy = 0
}

type Game struct {
	dots [dotCount]Dot
	dotImage *ebiten.Image
}

func NewGame() *Game {
	g := new(Game)
	for i := 0; i < dotCount; i++ {
		g.dots[i].Init()
	}

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    18,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	g.dotImage = ebiten.NewImage(20, 20)
	text.Draw(g.dotImage, "●", mplusNormalFont, 0, 18,
		color.RGBA{0x10, 0x30, 0x40, 0xff})
	return g;
}

func (g *Game) Update() error {
	for i := 0; i < dotCount; i++ {
		g.dots[i].x += g.dots[i].dx / 8
		g.dots[i].y += g.dots[i].dy / 7
		if screenWidth / 2 < g.dots[i].x {
			g.dots[i].dx--
		} else {
			g.dots[i].dx++
		}
		if screenHeight / 2 < g.dots[i].y {
			g.dots[i].dy--
		} else {
			g.dots[i].dy++
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < dotCount; i++ {
		op := &ebiten.DrawImageOptions{}
		op.CompositeMode = ebiten.CompositeModeLighter
		op.GeoM.Translate(g.dots[i].x, g.dots[i].y)
		screen.DrawImage(g.dotImage, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
