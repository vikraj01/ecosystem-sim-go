package assets

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"github.com/hajimehoshi/ebiten/v2"
)

func LoadTexture(path string) (*ebiten.Image, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	log.Printf("Loaded image with format: %s", format)

	ebitenImg := ebiten.NewImageFromImage(img)

	return ebitenImg, nil
}
