package beam

import (
	"image"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/tasks"
)

type Beam struct {
	standingHead *ebiten.Image
	standingTail []*ebiten.Image
	standingOpts *colorm.DrawImageOptions
	colorm       *colorm.ColorM

	tasks tasks.Tasks

	tailIndex  int
	facingLeft bool
	reversed   bool
	dirty      bool

	img *ebiten.Image

	Options *ebiten.DrawImageOptions
}

func New() (*Beam, error) {
	b := &Beam{}

	img, err := assets.Image("standing/Standing_Beam-head.png")
	if err != nil {
		return nil, err
	}
	b.standingHead = img

	b.standingTail = make([]*ebiten.Image, 11)
	for i := 0; i < 11; i++ {
		img, err := assets.Image("standing/Standing_Beam-" + strconv.Itoa(i+1) + ".png")
		if err != nil {
			return nil, err
		}
		b.standingTail[i] = img
	}

	b.standingOpts = &colorm.DrawImageOptions{}

	b.colorm = &colorm.ColorM{}

	b.tasks.Add(tasks.Tick(10, func() error {
		b.tailIndex++
		b.tailIndex %= 11
		b.dirty = true
		return nil
	}))

	b.dirty = true
	b.img = ebiten.NewImage(129, 125)

	b.Options = &ebiten.DrawImageOptions{}

	return b, nil
}

func (b *Beam) Draw(dst *ebiten.Image) {
	if b.dirty {
		b.img.Clear()

		colorm.DrawImage(b.img, b.standingHead, *b.colorm, b.standingOpts)
		colorm.DrawImage(b.img, b.standingTail[b.tailIndex], *b.colorm, b.standingOpts)

		b.dirty = false
	}

	dst.DrawImage(b.img, b.Options)
}

func (b *Beam) Update() error {
	if err := b.tasks.Update(); err != nil {
		return err
	}

	return nil
}

func (b *Beam) SetFacingLeft(facingLeft bool) {
	if facingLeft == b.facingLeft {
		return
	}
	b.facingLeft = facingLeft

	if b.facingLeft {
		b.standingOpts.GeoM.Scale(-1, 1)
		b.standingOpts.GeoM.Translate(244, 0)
	} else {
		b.standingOpts.GeoM.Reset()
	}

	b.dirty = true
}

func (b *Beam) SetReversed(reversed bool) {
	if reversed == b.reversed {
		return
	}
	b.reversed = reversed

	if b.reversed {
		b.colorm.ChangeHSV(-2.8, 1.5, 1)
	} else {
		b.colorm.Reset()
	}

	b.dirty = true
}

func (b *Beam) Bounds() image.Rectangle {
	return b.img.Bounds()
}
