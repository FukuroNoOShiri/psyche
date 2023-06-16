package play

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/FukuroNoOShiri/psyche/game"
)

type Scene struct {
	g *game.Game
}

var _ game.Scene = &Scene{}

func (s *Scene) Init(game *game.Game) error {
	s.g = game
	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
}

func (s *Scene) Update() error {
	return nil
}

func (s *Scene) Layout(_, _ int) (int, int) {
	return 1920, 1080
}
