package game

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Scene
	Audio *audio.Context
}

var _ ebiten.Game = &Game{}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) && ebiten.IsKeyPressed(ebiten.KeyControlLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyF4) && ebiten.IsKeyPressed(ebiten.KeyAltLeft) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) ||
		inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	return g.Scene.Update()
}

func (g *Game) SetScene(s Scene) error {
	if g.Scene != nil {
		g.Scene.Dispose()
	}
	g.Scene = s
	return g.Scene.Init(g)
}
