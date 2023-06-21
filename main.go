package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/urfave/cli/v2"

	"github.com/FukuroNoOShiri/psyche/game"
	"github.com/FukuroNoOShiri/psyche/intro"
	"github.com/FukuroNoOShiri/psyche/play"
	"github.com/FukuroNoOShiri/psyche/splash"
	"github.com/FukuroNoOShiri/psyche/title"
)

func main() {
	app := &cli.App{
		Name:   "psyche",
		Usage:  "Play Psyché",
		Action: run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "skip-to",
				Usage: "skips to specific scene",
			},
			&cli.BoolFlag{
				Name:  "full-screen",
				Usage: "starts in full screen mode",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(ctx *cli.Context) error {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Psyché")

	if ctx.Bool("full-screen") {
		ebiten.SetFullscreen(true)
	}

	playScene := &play.Scene{}
	introScene := &intro.Scene{
		Next: playScene,
	}
	titleScene := &title.Scene{
		Next: introScene,
	}
	splashScene := &splash.Scene{
		Next: titleScene,
	}

	var firstScene game.Scene = splashScene

	switch ctx.String("skip-to") {
	case "play":
		firstScene = playScene
	case "title":
		firstScene = titleScene
	case "intro":
		firstScene = introScene
	}

	game := &game.Game{
		Audio: audio.NewContext(48000),
	}

	if err := game.SetScene(firstScene); err != nil {
		return err
	}

	return ebiten.RunGame(game)
}
