package play

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/game"
	"github.com/FukuroNoOShiri/psyche/tasks"
	"github.com/FukuroNoOShiri/psyche/utils"
)

const (
	fallSpeed float64 = 5

	jumpSpeed    float64 = 10
	maxJumpTicks int     = 10
)

type Scene struct {
	g *game.Game

	space *resolv.Space
	tasks tasks.Tasks

	player     *resolv.Object
	dx         float64
	dy         float64
	facingLeft bool
	jumping    bool
	reversed   bool

	idleImg          *utils.ImageWithOptions
	idleReversedImg  *utils.ImageWithOptions
	greenPlatformImg *utils.ImageWithOptions
	bg               *utils.ImageWithOptions
}

var _ game.Scene = &Scene{}

func (s *Scene) Init(game *game.Game) error {
	s.g = game

	s.space = resolv.NewSpace(1920, 1080, 16, 16)

	s.space.Add(resolv.NewObject(0, 1080-184, 1920, 100, "platform"))
	s.space.Add(resolv.NewObject(0, 0, 16, 1080, "wall"))

	s.player = resolv.NewObject(200, 1080-400, 95, 100)
	s.space.Add(s.player)

	s.dy = fallSpeed

	img, err := assets.Idle()
	if err != nil {
		return err
	}
	s.idleImg = &utils.ImageWithOptions{
		Image:   img,
		Options: &ebiten.DrawImageOptions{},
	}

	var c colorm.ColorM

	c.ChangeHSV(-2.714, 1.44, 1)
	s.idleReversedImg = &utils.ImageWithOptions{
		Image:   ebiten.NewImage(s.idleImg.Image.Bounds().Dx(), s.idleImg.Image.Bounds().Dy()),
		Options: &ebiten.DrawImageOptions{},
	}
	colorm.DrawImage(s.idleReversedImg.Image, s.idleImg.Image, c, nil)

	img, err = assets.GreenPlatformFg()
	if err != nil {
		return err
	}
	s.greenPlatformImg = &utils.ImageWithOptions{
		Image: img,
	}

	img, err = assets.GreenPlatformBg()
	if err != nil {
		return err
	}
	s.bg = &utils.ImageWithOptions{
		Image: img,
	}

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	s.bg.Draw(screen)

	var playerImg *utils.ImageWithOptions
	if s.reversed {
		playerImg = s.idleReversedImg
	} else {
		playerImg = s.idleImg
	}
	playerImg.Options.GeoM.Reset()
	if s.facingLeft {
		playerImg.Options.GeoM.Scale(-1, 1)
		playerImg.Options.GeoM.Translate(95, 0)
	}
	playerImg.Options.GeoM.Translate(s.player.X, s.player.Y)
	playerImg.Draw(screen)

	s.greenPlatformImg.Draw(screen)
}

func (s *Scene) Update() error {
	if err := s.tasks.Update(); err != nil {
		return err
	}

	if l, r := ebiten.IsKeyPressed(ebiten.KeyLeft), ebiten.IsKeyPressed(ebiten.KeyRight); l && !r {
		s.dx = -5.0
		s.facingLeft = true
	} else if r && !l {
		s.dx = 5.0
		s.facingLeft = false
	} else {
		s.dx = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		s.reversed = !s.reversed
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
			s.tasks.Add(tasks.AfterTicks(maxJumpTicks, func() error {
				s.jumping = false
				s.dy = fallSpeed
				return nil
			}))
		}
	}

	s.player.X += s.dx

	s.player.Update()

	return nil
}

func (s *Scene) Dispose() {

}

func (s *Scene) Layout(_, _ int) (int, int) {
	return 1920, 1080
}
