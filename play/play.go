package play

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/game"
)

type Scene struct {
	g *game.Game

	space *resolv.Space

	player     *resolv.Object
	dx         float64
	facingLeft bool

	idleImg *ebiten.Image
}

var _ game.Scene = &Scene{}

func (s *Scene) Init(game *game.Game) error {
	s.g = game

	s.space = resolv.NewSpace(1920, 1080, 16, 16)
	s.space.Add(resolv.NewObject(0, 1080-100, 1920, 100))

	s.player = resolv.NewObject(200, 1080-200, 95, 100)
	s.space.Add(s.player)

	img, err := assets.Idle()
	if err != nil {
		return err
	}
	s.idleImg = img

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	opts := &ebiten.DrawImageOptions{}
	if s.facingLeft {
		opts.GeoM.Scale(-1, 1)
		opts.GeoM.Translate(95, 0)
	}
	opts.GeoM.Translate(s.player.X, s.player.Y)
	screen.DrawImage(s.idleImg, opts)
}

func (s *Scene) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		s.dx = 5.0
		s.facingLeft = false
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		s.dx = 0.0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		s.dx = -5.0
		s.facingLeft = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		s.dx = 0.0
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
