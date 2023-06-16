package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Init(*Game) error
	Update() error
	Draw(*ebiten.Image)
	Layout(int, int) (int, int)
}
