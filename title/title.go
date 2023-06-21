package title

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/game"
)

type Scene struct {
	g    *game.Game
	Next game.Scene

	textColor color.Color

	titleFace font.Face
	titlePos  image.Point

	startFace font.Face
	startPos  image.Point
}

var _ game.Scene = &Scene{}

func (s *Scene) Init(game *game.Game) error {
	s.g = game

	s.textColor = color.RGBA{84, 137, 169, 0}

	face, err := assets.SrirachaRegular(&opentype.FaceOptions{
		DPI:  72,
		Size: 120,
	})
	if err != nil {
		return err
	}
	s.titleFace = face
	titleBounds := text.BoundString(s.titleFace, "Psyché")
	s.titlePos.X, s.titlePos.Y = (1920-titleBounds.Dx())/2, 400

	face, err = assets.SrirachaRegular(&opentype.FaceOptions{
		DPI:  48,
		Size: 120,
	})
	if err != nil {
		return err
	}
	s.startFace = face
	startBounds := text.BoundString(s.startFace, "Press Enter to start")
	s.startPos.X, s.startPos.Y = (1920-startBounds.Dx())/2, 600

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	text.Draw(screen, "Psyché", s.titleFace, s.titlePos.X, s.titlePos.Y, s.textColor)
	text.Draw(screen, "Press Enter to start", s.startFace, s.startPos.X, s.startPos.Y, s.textColor)
}

func (s *Scene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && !ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
		return s.g.SetScene(s.Next)
	}

	return nil
}

func (s *Scene) Dispose() {
}

func (s *Scene) Layout(_, _ int) (int, int) {
	return 1920, 1080
}
