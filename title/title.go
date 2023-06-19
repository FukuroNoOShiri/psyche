package title

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/game"
)

type Scene struct {
	g    *game.Game
	Next game.Scene

	titleFont   font.Face
	titleColor  color.Color
	titleBounds image.Rectangle
}

var _ game.Scene = &Scene{}

func (s *Scene) Init(game *game.Game) error {
	s.g = game

	face, err := assets.SrirachaRegular(&opentype.FaceOptions{
		DPI:  72,
		Size: 120,
	})
	if err != nil {
		return err
	}
	s.titleFont = face

	s.titleColor = color.RGBA{84, 137, 169, 0}

	s.titleBounds = text.BoundString(s.titleFont, "Psyché")

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	text.Draw(screen, "Psyché", s.titleFont, (1920-s.titleBounds.Dx())/2, (1080-s.titleBounds.Dy())/2, s.titleColor)
}

func (s *Scene) Update() error {
	return nil
}

func (s *Scene) Dispose() {

}

func (s *Scene) Layout(_, _ int) (int, int) {
	return 1920, 1080
}
