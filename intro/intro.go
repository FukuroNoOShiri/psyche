package intro

import (
	"image"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/game"
	"github.com/FukuroNoOShiri/psyche/tasks"
	"github.com/FukuroNoOShiri/psyche/utils"
)

const beamScale = 4

type scene struct {
	Next game.Scene

	background *ebiten.Image
	beam       *utils.ImageWithOptions

	textColor color.Color
	textFace  font.Face

	text        string
	textPos     image.Point
	visibleText string
}

var Scene = &scene{}
var _ game.Scene = Scene

func (s *scene) Init() error {
	s.background = ebiten.NewImage(1920, 1080)
	img, err := assets.Image("GreenPlatform1-sky.png")
	if err != nil {
		return err
	}
	s.background.DrawImage(img, nil)
	img, err = assets.Image("Intro-psyche.png")
	if err != nil {
		return err
	}
	s.background.DrawImage(img, nil)

	img, err = assets.Image("Idle.png")
	if err != nil {
		return err
	}
	s.beam = &utils.ImageWithOptions{
		Image:   img,
		Options: &ebiten.DrawImageOptions{},
	}
	s.beam.Options.GeoM.Scale(beamScale, beamScale)
	s.beam.Options.GeoM.Translate((1920.0-float64(s.beam.Image.Bounds().Dx())*beamScale)/2, (1080.0-float64(s.beam.Image.Bounds().Dy())*beamScale)/2)

	s.textColor = color.RGBA{84, 137, 169, 0}

	face, err := assets.FontFace("Sriracha-Regular.ttf", &opentype.FaceOptions{
		DPI:  36,
		Size: 120,
	})
	if err != nil {
		return err
	}
	s.textFace = face

	game.Tasks.Add(tasks.After(2*time.Second, func() error {
		return s.write("My...")
	}), "intro")

	game.Tasks.Add(tasks.After(5*time.Second, func() error {
		return s.write("...name.")
	}), "intro")

	game.Tasks.Add(tasks.After(10*time.Second, func() error {
		return s.write("I can't remember it...")
	}), "intro")

	game.Tasks.Add(tasks.After(15*time.Second, func() error {
		return s.write("I lost my name!")
	}), "intro")

	game.Tasks.Add(tasks.After(20*time.Second, func() error {
		return game.SetScene(s.Next)
	}), "intro")

	return nil
}

func (s *scene) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.background, nil)
	s.beam.Draw(screen)
	text.Draw(screen, s.visibleText, s.textFace, s.textPos.X, s.textPos.Y, s.textColor)
}

func (s *scene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		game.Tasks.Cancel("intro")
		return game.SetScene(s.Next)
	}

	return nil
}

func (s *scene) Dispose() {
	s.textFace = nil
}

func (s *scene) write(txt string) error {
	s.text = txt
	s.visibleText = ""
	bnd := text.BoundString(s.textFace, txt)
	s.textPos = image.Point{
		X: (1920 - bnd.Dx()) / 2,
		Y: 960 + bnd.Dy(),
	}
	if txt == "" {
		return nil
	}
	return s.addLetter()
}

func (s *scene) addLetter() error {
	s.visibleText = s.text[:len(s.visibleText)+1]
	if s.text != s.visibleText {
		game.Tasks.Add(tasks.AfterTicks(4, s.addLetter), "intro")
	}
	return nil
}
