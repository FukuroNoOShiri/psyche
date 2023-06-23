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

type scene struct {
	space *resolv.Space

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
	rock1Img          *utils.ImageWithOptions
}

var Scene = &scene{}
var _ game.Scene = Scene

func (s *scene) Init() error {
	s.space = resolv.NewSpace(1920, 1080, 16, 16)

	s.space.Add(resolv.NewObject(0, 1080-184, 1920-773, 100, "platform"))
	s.space.Add(resolv.NewObject(0, 0, 16, 1080, "wall"))

	s.player = resolv.NewObject(200, 1080-400, 95, 100)
	s.space.Add(s.player)

	s.dy = fallSpeed

	img, err := assets.Image("Idle.png")
	if err != nil {
		return err
	}
	s.idleImg = &utils.ImageWithOptions{
		Image:   img,
		Options: &ebiten.DrawImageOptions{},
	}

	var c colorm.ColorM

	c.ChangeHSV(-2.8, 1.5, 1)
	s.idleReversedImg = &utils.ImageWithOptions{
		Image:   ebiten.NewImage(s.idleImg.Image.Bounds().Dx(), s.idleImg.Image.Bounds().Dy()),
		Options: &ebiten.DrawImageOptions{},
	}
	colorm.DrawImage(s.idleReversedImg.Image, s.idleImg.Image, c, nil)

	img, err = assets.Image("GreenPlatform1-grass1.png")
	if err != nil {
		return err
	}
	s.greenPlatformImg = &utils.ImageWithOptions{
		Image: img,
		Options: &ebiten.DrawImageOptions{},
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
	img, err = assets.Image("BGtrees.png")
	if err != nil {
		return err
	}
	s.bg.Image.DrawImage(img, nil)

    img, err = assets.Image("rock1.png")
	if err != nil {
		return err
	}
	s.rock1Img = &utils.ImageWithOptions{
		Image: img,
		Options: &ebiten.DrawImageOptions{},
	}

	return nil
}

func (s *scene) Draw(screen *ebiten.Image) {
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
    s.rock1Img.Options.GeoM.Apply(1080-171, 70)
	s.greenPlatformImg.Draw(screen)
	s.rock1Img.Draw(screen)
}

func (s *scene) Update() error {
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
			game.Tasks.Add(tasks.AfterTicks(maxJumpTicks, func() error {
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

func (s *scene) Dispose() {
}
