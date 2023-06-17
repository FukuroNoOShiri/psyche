package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageWithOptions struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions
}

func (iwo ImageWithOptions) Draw(dst *ebiten.Image) {
	dst.DrawImage(iwo.Image, iwo.Options)
}
