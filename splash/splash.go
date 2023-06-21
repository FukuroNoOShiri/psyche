package splash

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/game"
	"github.com/FukuroNoOShiri/psyche/tasks"
	"github.com/FukuroNoOShiri/psyche/utils"
)

type Scene struct {
	g    *game.Game
	Next game.Scene

	bg            color.RGBA
	logo          utils.ImageWithOptions
	fadingOverlay utils.ImageWithOptions
	sound1        *mp3.Stream
	sound2        *mp3.Stream

	tasks tasks.Tasks

	canSkip           bool
	fadingProgression float64
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

	sound1, err := assets.OwlSound1()
	if err != nil {
		return err
	}
	s.sound1 = sound1

	sound2, err := assets.OwlSound2()
	if err != nil {
		return err
	}
	s.sound2 = sound2

	s.tasks.Add(tasks.After(500*time.Millisecond, func() error {
		p, err := s.g.Audio.NewPlayer(sound1)
		if err != nil {
			return err
		}
		p.Play()
		return nil
	}))

	s.tasks.Add(tasks.After(3*time.Second, func() error {
		s.canSkip = true
		return nil
	}))

	s.tasks.Add(tasks.After(6*time.Second, s.fade), "fade")

	fadingOverlayOpts := &ebiten.DrawImageOptions{}
	s.fadingOverlay = utils.ImageWithOptions{Image: ebiten.NewImage(1920, 1080), Options: fadingOverlayOpts}

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(s.bg)
	s.logo.Draw(screen)

	if s.fadingProgression > 0 {
		s.fadingOverlay.Image.Fill(color.RGBA{0, 0, 0, uint8(s.fadingProgression * 255)})
		s.fadingOverlay.Draw(screen)
	}
}

func (s *Scene) Update() error {
	if s.canSkip {
		if ok, _ := utils.IsSomeKeyJustPressed(ebiten.KeySpace, ebiten.KeyEnter, ebiten.KeyEscape); ok {
			s.tasks.Cancel("fade")
			if err := s.fade(); err != nil {
				return err
			}
		}
	}

	if err := s.tasks.Update(); err != nil {
		return err
	}

	return nil
}

func (s *Scene) fade() error {
	s.tasks.Add(tasks.During(1*time.Second, func(progression float64) error {
		if progression == 1 {
			return s.g.SetScene(s.Next)
		}
		s.fadingProgression = progression
		return nil
	}))

	p, err := s.g.Audio.NewPlayer(s.sound2)
	if err != nil {
		return err
	}
	p.Play()

	return nil
}

func (s *Scene) Dispose() {
	s.logo.Dispose()
	s.fadingOverlay.Dispose()
	s.sound1 = nil
	s.sound2 = nil
}

func (s *Scene) Layout(_, _ int) (int, int) {
	return 1920, 1080
}
