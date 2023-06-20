package assets

import (
	"bytes"
	_ "embed"
	_ "image/jpeg"
	_ "image/png"
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

	//go:embed Idle.png
	idle     []byte
	loadIdle sync.Once
	idleImg  *ebiten.Image

	//go:embed GreenPlatform1-grass.png
	greenPlatformGrass  []byte
	loadGreenPlatformFg sync.Once
	greenPlatformFg     *ebiten.Image

	//go:embed GreenPlatform1-sky.png
	greenPlatformSky []byte
	//go:embed GreenPlatform1-clouds.png
	greenPlatformClouds []byte
	//go:embed GreenPlatform1-trees.png
	greenPlatformTrees  []byte
	loadGreenPlatformBg sync.Once
	greenPlatformBg     *ebiten.Image
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

func Idle() (img *ebiten.Image, err error) {
	loadIdle.Do(func() {
		idleImg, err = bytesToImage(idle)
	})
	if err != nil {
		return nil, err
	}

	img = idleImg

	return
}

func GreenPlatformFg() (img *ebiten.Image, err error) {
	loadGreenPlatformFg.Do(func() {
		greenPlatformFg, err = bytesToImage(greenPlatformGrass)
	})
	if err != nil {
		return nil, err
	}

	img = greenPlatformFg

	return
}

func GreenPlatformBg() (img *ebiten.Image, err error) {
	loadGreenPlatformBg.Do(func() {
		greenPlatformBg = ebiten.NewImage(1920, 1080)

		var img *ebiten.Image

		img, err = bytesToImage(greenPlatformSky)
		if err != nil {
			return
		}
		greenPlatformBg.DrawImage(img, nil)

		img, err = bytesToImage(greenPlatformClouds)
		if err != nil {
			return
		}
		greenPlatformBg.DrawImage(img, nil)

		img, err = bytesToImage(greenPlatformTrees)
		if err != nil {
			return
		}
		greenPlatformBg.DrawImage(img, nil)
	})
	if err != nil {
		return nil, err
	}

	img = greenPlatformBg

	return
}

func bytesToImage(b []byte) (img *ebiten.Image, err error) {
	img, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(b))
	return
}
