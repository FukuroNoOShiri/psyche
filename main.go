package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/FukuroNoOShiri/psyche/scene"
	"github.com/FukuroNoOShiri/psyche/splash"
)

type Game struct {
	scene scene.Scene
}

func (g *Game) Update() error {
	return g.scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Psyché")

	splash, err := splash.New()
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		scene: splash,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
