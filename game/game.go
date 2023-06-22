package game

import (
	"image/color"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/FukuroNoOShiri/psyche/tasks"
)

var (
	Game  = game{}
	Tasks tasks.Tasks
	Audio = audio.NewContext(48000)

	scene Scene

	fadeOverlay     = ebiten.NewImage(1920, 1080)
	fadeProgression float64
)

type game struct {
}

var _ ebiten.Game = Game

func (game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) && ebiten.IsKeyPressed(ebiten.KeyControlLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyF4) && ebiten.IsKeyPressed(ebiten.KeyAltLeft) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) ||
		inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if err := Tasks.Update(); err != nil {
		return err
	}

	return scene.Update()
}

func (game) Draw(screen *ebiten.Image) {
	scene.Draw(screen)

	if fadeProgression > 0 {
		fadeOverlay.Fill(color.RGBA{0, 0, 0, uint8(fadeProgression * 255)})
		screen.DrawImage(fadeOverlay, nil)
	}
}

func (game) Layout(_, _ int) (int, int) {
	return 1920, 1080
}

func SetScene(s Scene) error {
	if scene == nil {
		defer FadeIn(1 * time.Second)
		scene = s
		return scene.Init()
	}

	FadeOut(1 * time.Second)
	Tasks.Add(tasks.After(1*time.Second, func() error {
		defer FadeIn(1 * time.Second)
		scene.Dispose()
		scene = s
		return scene.Init()
	}))
	return nil
}

func FadeOut(duration time.Duration) {
	Tasks.Add(tasks.During(duration, func(progression float64) error {
		fadeProgression = progression
		return nil
	}))
}

func FadeIn(duration time.Duration) {
	Tasks.Add(tasks.During(duration, func(progression float64) error {
		fadeProgression = 1 - progression
		return nil
	}))
}
