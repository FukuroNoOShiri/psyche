package splash

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/game"
	"github.com/FukuroNoOShiri/psyche/utils"
)

type Scene struct {
	g        *game.Game
	Next     game.Scene
	bg       color.RGBA
	logo     *ebiten.Image
	ticks    int
	maxTicks int
}

var _ game.Scene = &Scene{}

func (s *Scene) Init(game *game.Game) error {
	s.g = game

	s.bg = color.RGBA{249, 239, 214, 0}

	logo, err := assets.Fukuronooshiri()
	if err != nil {
		return err
	}
	s.logo = logo

	s.maxTicks = ebiten.TPS() * 5

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(s.bg)

	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	lw, lh := s.logo.Bounds().Dx(), s.logo.Bounds().Dy()

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64((w-lw)/2), float64((h-lh)/2))

	screen.DrawImage(s.logo, opt)
}

func (s *Scene) Update() error {
	if ok, _ := utils.IsSomeKeyJustPressed(ebiten.KeySpace, ebiten.KeyEnter, ebiten.KeyEscape); ok {
		return s.g.SetScene(s.Next)
	}

	s.ticks++

	if s.ticks == s.maxTicks {
		return s.g.SetScene(s.Next)
	}

	return nil
}

func (s *Scene) Layout(_, _ int) (int, int) {
	return 1920, 1080
}
