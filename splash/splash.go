package splash

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/game"
	"github.com/FukuroNoOShiri/psyche/utils"
)

type Scene struct {
	g    *game.Game
	Next game.Scene

	bg   color.RGBA
	logo utils.ImageWithOptions

	ticks       int
	sound1Ticks int
	minTicks    int
	maxTicks    int

	fading         bool
	fadingTicks    int
	fadingMaxTicks int
	fadingOverlay  utils.ImageWithOptions
}

var _ game.Scene = &Scene{}

func (s *Scene) Init(game *game.Game) error {
	s.g = game

	s.bg = color.RGBA{249, 239, 214, 0}

	logo, err := assets.Fukuronooshiri()
	if err != nil {
		return err
	}
	logoW, logoH := logo.Bounds().Dx(), logo.Bounds().Dy()
	logoOpts := &ebiten.DrawImageOptions{}
	logoOpts.GeoM.Translate(float64((1920-logoW)/2), float64((1080-logoH)/2))
	s.logo = utils.ImageWithOptions{Image: logo, Options: logoOpts}

	s.sound1Ticks = ebiten.TPS() / 2
	s.minTicks = ebiten.TPS() * 3
	s.maxTicks = ebiten.TPS() * 5
	s.fadingMaxTicks = ebiten.TPS()

	fadingOverlayOpts := &ebiten.DrawImageOptions{}
	s.fadingOverlay = utils.ImageWithOptions{Image: ebiten.NewImage(1920, 1080), Options: fadingOverlayOpts}

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(s.bg)

	s.logo.Draw(screen)

	if s.fading {
		s.fadingOverlay.Image.Fill(color.RGBA{0, 0, 0, uint8(s.fadingTicks * 255 / s.fadingMaxTicks)})
		s.fadingOverlay.Draw(screen)
	}
}

func (s *Scene) Update() error {
	if s.fading {
		if s.fadingTicks++; s.fadingTicks == s.fadingMaxTicks {
			return s.g.SetScene(s.Next)
		}
		return nil
	}

	if s.ticks == s.sound1Ticks {
		snd, err := assets.OwlSound1()
		if err != nil {
			return err
		}

		p, err := s.g.Audio.NewPlayer(snd)
		if err != nil {
			return err
		}

		p.Play()
	}

	if s.ticks >= s.minTicks {
		if ok, _ := utils.IsSomeKeyJustPressed(ebiten.KeySpace, ebiten.KeyEnter, ebiten.KeyEscape); ok {
			return s.fade()
		}
	}

	if s.ticks++; s.ticks == s.maxTicks {
		return s.fade()
	}

	return nil
}

func (s *Scene) fade() error {
	s.fading = true

	snd, err := assets.OwlSound2()
	if err != nil {
		return err
	}

	p, err := s.g.Audio.NewPlayer(snd)
	if err != nil {
		return err
	}

	p.Play()

	return nil
}

func (s *Scene) Dispose() {
	s.logo.Dispose()
	s.fadingOverlay.Dispose()
}

func (s *Scene) Layout(_, _ int) (int, int) {
	return 1920, 1080
}
