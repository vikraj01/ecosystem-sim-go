package game

import (
	"ecosystem-sim/internals/assets"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	// my image from github?
)

type Game struct {
	x       float64
	y       float64
	texture *ebiten.Image
}

func (g *Game) Update() error {
	g.x += 0.3
	g.y += 0.5
	if g.x > 640 {
		g.x = 0
	}
	if g.y > 480 {
		g.y = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	greenBgColor := color.RGBA{0, 128, 0, 255}
	screen.Fill(greenBgColor)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)

	scaleX, scaleY := 0.2, 0.2
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(g.texture, op)
}

func (g *Game) Layout(outsideHeight, outsideWidth int) (int, int) {
	return 640, 480
}

func NewGame() (*Game, error) {
	spritePath := "assets/textures/gopher.png"
	texture, err := assets.LoadTexture(spritePath)
	if err != nil {
		return nil, err
	}
	return &Game{texture: texture}, nil
}
