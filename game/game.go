package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Game struct {
	Scene
	Audio *audio.Context
}

var _ ebiten.Game = &Game{}

func (g *Game) SetScene(s Scene) error {
	if g.Scene != nil {
		g.Scene.Dispose()
	}
	g.Scene = s
	return g.Scene.Init(g)
}
