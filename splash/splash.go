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

type scene struct {
	Next game.Scene

	bg     color.RGBA
	logo   utils.ImageWithOptions
	sound1 *mp3.Stream
	sound2 *mp3.Stream

	canSkip bool
}

var Scene = &scene{}
var _ game.Scene = Scene

func (s *scene) Init() error {
	s.bg = color.RGBA{249, 239, 214, 0}

	logo, err := assets.Image("fukuronooshiri.jpg")
	if err != nil {
		return err
	}
	logoW, logoH := logo.Bounds().Dx(), logo.Bounds().Dy()
	logoOpts := &ebiten.DrawImageOptions{}
	logoOpts.GeoM.Translate(float64((1920-logoW)/2), float64((1080-logoH)/2))
	s.logo = utils.ImageWithOptions{Image: logo, Options: logoOpts}

	sound1, err := assets.Mp3Stream("owl-sound-1.mp3")
	if err != nil {
		return err
	}
	s.sound1 = sound1

	sound2, err := assets.Mp3Stream("owl-sound-2.mp3")
	if err != nil {
		return err
	}
	s.sound2 = sound2

	game.Tasks.Add(tasks.After(500*time.Millisecond, func() error {
		p, err := game.Audio.NewPlayer(sound1)
		if err != nil {
			return err
		}
		p.Play()
		return nil
	}))

	game.Tasks.Add(tasks.After(3*time.Second, func() error {
		s.canSkip = true
		return nil
	}))

	game.Tasks.Add(tasks.After(6*time.Second, s.leave), "leave")

	return nil
}

func (s *scene) Draw(screen *ebiten.Image) {
	screen.Fill(s.bg)
	s.logo.Draw(screen)
}

func (s *scene) Update() error {
	if s.canSkip {
		if ok, _ := utils.IsSomeKeyJustPressed(ebiten.KeySpace, ebiten.KeyEnter, ebiten.KeyEscape); ok {
			game.Tasks.Cancel("leave")
			s.canSkip = false
			if err := s.leave(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *scene) leave() error {
	p, err := game.Audio.NewPlayer(s.sound2)
	if err != nil {
		return err
	}
	p.Play()

	return game.SetScene(s.Next)
}

func (s *scene) Dispose() {
	s.logo.Dispose()
	s.sound1 = nil
	s.sound2 = nil
}
