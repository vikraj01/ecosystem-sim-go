package main

import (
	"ecosystem-sim/internals/game"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ecosystem Simulation - muhahaha")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
