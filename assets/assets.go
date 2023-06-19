package assets

import (
	"bytes"
	_ "embed"
	_ "image/jpeg"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var (
	//go:embed fukuronooshiri.jpg
	fukuronooshiri []byte

	//go:embed owl-sound-1.mp3
	owlSound1 []byte

	//go:embed owl-sound-2.mp3
	owlSound2 []byte

	//go:embed Sriracha-Regular.ttf
	srirachaRegular         []byte
	srirachaRegularFont     *sfnt.Font
	loadSrirachaRegularFont sync.Once
)

func Fukuronooshiri() (*ebiten.Image, error) {
	return bytesToImage(fukuronooshiri)
}

func OwlSound1() (*mp3.Stream, error) {
	return mp3.DecodeWithoutResampling(bytes.NewReader(owlSound1))
}

func OwlSound2() (*mp3.Stream, error) {
	return mp3.DecodeWithoutResampling(bytes.NewReader(owlSound2))
}

func SrirachaRegular(opts *opentype.FaceOptions) (face font.Face, err error) {
	loadSrirachaRegularFont.Do(func() {
		srirachaRegularFont, err = opentype.Parse(srirachaRegular)
	})
	if err != nil {
		return nil, err
	}

	face, err = opentype.NewFace(srirachaRegularFont, opts)
	return
}

func bytesToImage(b []byte) (img *ebiten.Image, err error) {
	img, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(b))
	return
}
