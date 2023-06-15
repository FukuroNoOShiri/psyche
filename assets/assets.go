package assets

import (
	"bytes"
	_ "embed"
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	//go:embed fukuronooshiri.jpg
	fukuronooshiri []byte
)

func Fukuronooshiri() (*ebiten.Image, error) {
	return bytesToImage(fukuronooshiri)
}

func bytesToImage(b []byte) (img *ebiten.Image, err error) {
	img, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(b))
	return
}
