package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/urfave/cli/v2"

	"github.com/FukuroNoOShiri/psyche/game"
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
			&cli.BoolFlag{
				Name:  "skip-splash",
				Usage: "skips splash screen",
			},
			&cli.BoolFlag{
				Name:  "skip-title",
				Usage: "skips title screen",
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
	titleScene := &title.Scene{
		Next: playScene,
	}
	splashScene := &splash.Scene{
		Next: titleScene,
	}

	var firstScene game.Scene = splashScene

	if skipSplash, skipTitle := ctx.Bool("skip-splash"), ctx.Bool("skip-title"); skipSplash {
		if skipTitle {
			firstScene = playScene
		} else {
			firstScene = titleScene
		}
	} else if skipTitle {
		splashScene.Next = playScene
	}

	game := &game.Game{
		Audio: audio.NewContext(48000),
	}

	if err := game.SetScene(firstScene); err != nil {
		return err
	}

	return ebiten.RunGame(game)
}
