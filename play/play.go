package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/beam"
	"github.com/FukuroNoOShiri/psyche/game"
	"github.com/FukuroNoOShiri/psyche/tasks"
	"github.com/FukuroNoOShiri/psyche/utils"
)

const (
	fallSpeed float64 = 5

	jumpSpeed    float64 = 10
	maxJumpTicks int     = 10
)

type scene struct {
	space *resolv.Space

	player   *resolv.Object
	dx       float64
	dy       float64
	jumping  bool
	reversed bool

	beam             *beam.Beam
	greenPlatformImg *utils.ImageWithOptions
	bg               *utils.ImageWithOptions
}

var Scene = &scene{}
var _ game.Scene = Scene

func (s *scene) Init() error {
	s.space = resolv.NewSpace(1920, 1080, 16, 16)

	s.space.Add(resolv.NewObject(0, 1080-184, 1920, 100, "platform"))
	s.space.Add(resolv.NewObject(0, 0, 16, 1080, "wall"))

	s.player = resolv.NewObject(200, 1080-400, 95, 100)
	s.space.Add(s.player)

	s.dy = fallSpeed

	beam, err := beam.New()
	if err != nil {
		return err
	}
	s.beam = beam

	img, err := assets.Image("GreenPlatform1-grass.png")
	if err != nil {
		return err
	}
	s.greenPlatformImg = &utils.ImageWithOptions{
		Image: img,
	}

	s.bg = &utils.ImageWithOptions{
		Image: ebiten.NewImage(1920, 1080),
	}
	img, err = assets.Image("GreenPlatform1-sky.png")
	if err != nil {
		return err
	}
	s.bg.Image.DrawImage(img, nil)
	img, err = assets.Image("GreenPlatform1-clouds.png")
	if err != nil {
		return err
	}
	s.bg.Image.DrawImage(img, nil)
	img, err = assets.Image("GreenPlatform1-trees.png")
	if err != nil {
		return err
	}
	s.bg.Image.DrawImage(img, nil)

	return nil
}

func (s *scene) Draw(screen *ebiten.Image) {
	s.bg.Draw(screen)

	s.beam.Options.GeoM.Reset()
	s.beam.Options.GeoM.Scale(0.53, 0.53)
	s.beam.Options.GeoM.Translate(s.player.X, s.player.Y)
	s.beam.Draw(screen)

	s.greenPlatformImg.Draw(screen)
}

func (s *scene) Update() error {
	if l, r := ebiten.IsKeyPressed(ebiten.KeyLeft), ebiten.IsKeyPressed(ebiten.KeyRight); l && !r {
		s.dx = -5.0
		s.beam.SetFacingLeft(true)
	} else if r && !l {
		s.dx = 5.0
		s.beam.SetFacingLeft(false)
	} else {
		s.dx = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		s.reversed = !s.reversed
		s.beam.SetReversed(s.reversed)
	}

	if collision := s.player.Check(s.dx, 0, "wall"); collision != nil {
		s.player.X += collision.ContactWithObject(collision.Objects[0]).X()
	} else {
		s.player.X += s.dx
	}

	var onGround bool
	if collision := s.player.Check(0, s.dy, "platform"); collision != nil {
		s.player.Y += collision.ContactWithObject(collision.Objects[0]).Y()
		onGround = true
	} else {
		s.player.Y += s.dy
	}

	if onGround {
		if ok, _ := utils.IsSomeKeyJustPressed(ebiten.KeyArrowUp, ebiten.KeySpace); ok {
			s.jumping = true
			s.dy = -jumpSpeed
			game.Tasks.Add(tasks.AfterTicks(maxJumpTicks, func() error {
				s.jumping = false
				s.dy = fallSpeed
				return nil
			}))
		}
	}

	s.player.X += s.dx

	s.player.Update()
	if err := s.beam.Update(); err != nil {
		return err
	}

	return nil
}

func (s *scene) Dispose() {
}
