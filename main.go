package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/FukuroNoOShiri/psyche/game"
	"github.com/FukuroNoOShiri/psyche/play"
	"github.com/FukuroNoOShiri/psyche/splash"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Psyché")

	splashScene := &splash.Scene{}
	playScene := &play.Scene{}

	splashScene.Next = playScene

	game := &game.Game{
		Audio: audio.NewContext(48000),
	}
	if err := game.SetScene(splashScene); err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
