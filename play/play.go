package play

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/game"
	"github.com/FukuroNoOShiri/psyche/utils"
)

type Scene struct {
	g *game.Game

	space *resolv.Space

	player     *resolv.Object
	dx         float64
	dy         float64
	facingLeft bool

	idleImg          *utils.ImageWithOptions
	greenPlatformImg *utils.ImageWithOptions
}

var _ game.Scene = &Scene{}

func (s *Scene) Init(game *game.Game) error {
	s.g = game

	s.space = resolv.NewSpace(1920, 1080, 16, 16)

	s.space.Add(resolv.NewObject(0, 1080-100, 1920, 100, "platform"))
	s.space.Add(resolv.NewObject(0, 0, 16, 1080, "wall"))

	s.player = resolv.NewObject(200, 1080-300, 95, 100)
	s.space.Add(s.player)

	s.dy = 5

	img, err := assets.Idle()
	if err != nil {
		return err
	}
	s.idleImg = &utils.ImageWithOptions{
		Image:   img,
		Options: &ebiten.DrawImageOptions{},
	}

	img, err = assets.GreenPlatform()
	if err != nil {
		return err
	}
	s.greenPlatformImg = &utils.ImageWithOptions{
		Image:   img,
		Options: &ebiten.DrawImageOptions{},
	}
	s.greenPlatformImg.Options.GeoM.Translate(0, 1080-100)

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	s.idleImg.Options.GeoM.Reset()
	if s.facingLeft {
		s.idleImg.Options.GeoM.Scale(-1, 1)
		s.idleImg.Options.GeoM.Translate(95, 0)
	}
	s.idleImg.Options.GeoM.Translate(s.player.X, s.player.Y)
	s.idleImg.Draw(screen)

	s.greenPlatformImg.Draw(screen)
}

func (s *Scene) Update() error {

	if l, r := ebiten.IsKeyPressed(ebiten.KeyLeft), ebiten.IsKeyPressed(ebiten.KeyRight); l && !r {
		s.dx = -5.0
		s.facingLeft = true
	} else if r && !l {
		s.dx = 5.0
		s.facingLeft = false
	} else {
		s.dx = 0
	}

	if collision := s.player.Check(s.dx, 0, "wall"); collision != nil {
		s.player.X += collision.ContactWithObject(collision.Objects[0]).X()
	} else {
		s.player.X += s.dx
	}

	if collision := s.player.Check(0, s.dy, "platform"); collision != nil {
		s.player.Y += collision.ContactWithObject(collision.Objects[0]).Y()
	} else {
		s.player.Y += s.dy
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
