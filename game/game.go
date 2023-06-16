package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Scene
}

var _ ebiten.Game = &Game{}

func (g *Game) SetScene(s Scene) error {
	g.Scene = s
	return g.Scene.Init(g)
}
