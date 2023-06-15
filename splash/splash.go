package splash

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/scene"
)

type splash struct {
	bg   color.RGBA
	logo *ebiten.Image
}

func New() (scene.Scene, error) {
	logo, err := assets.Fukuronooshiri()
	if err != nil {
		return splash{}, err
	}

	return splash{
		bg:   color.RGBA{249, 239, 214, 0},
		logo: logo,
	}, nil
}

func (s splash) Draw(screen *ebiten.Image) {
	screen.Fill(s.bg)

	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	lw, lh := s.logo.Bounds().Dx(), s.logo.Bounds().Dy()

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64((w-lw)/2), float64((h-lh)/2))

	screen.DrawImage(s.logo, opt)
}

func (s splash) Update() error {
	return nil
}
