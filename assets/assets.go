package assets

import (
	"bytes"
	"embed"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	//go:embed *.jpg *.png *.mp3 *.ttf flying standing
	assets embed.FS
)

func Image(name string) (*ebiten.Image, error) {
	b, err := fs.ReadFile(assets, name)
	if err != nil {
		return nil, err
	}
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(b))
	return img, err
}

func Mp3Stream(name string) (*mp3.Stream, error) {
	b, err := fs.ReadFile(assets, name)
	if err != nil {
		return nil, err
	}
	return mp3.DecodeWithoutResampling(bytes.NewReader(b))
}

func FontFace(name string, opts *opentype.FaceOptions) (font.Face, error) {
	b, err := fs.ReadFile(assets, name)
	if err != nil {
		return nil, err
	}
	font, err := opentype.Parse(b)
	if err != nil {
		return nil, err
	}
	return opentype.NewFace(font, opts)
}
