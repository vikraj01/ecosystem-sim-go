package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
)

func LoadTexture(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	ebitenImg := ebiten.NewImageFromImage(img)
	return ebitenImg, nil
}
