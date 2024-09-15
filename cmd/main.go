package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	x float64
	y float64
}

func (g *Game) Update() error {
	g.x += 1
	g.y += 1
	if g.x > screenWidth {
		g.x = 0
	}
	if g.y > screenHeight {
		g.y = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	green := color.RGBA{0, 128, 0, 255}
	screen.Fill(green)
	
	organisam := color.RGBA{200, 200, 200, 255}
	vector.DrawFilledRect(screen, float32(g.x), float32(g.y), 50, 50, organisam, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple Animation With Ebiten")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
