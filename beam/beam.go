package beam

import (
	"image"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/FukuroNoOShiri/psyche/assets"
	"github.com/FukuroNoOShiri/psyche/tasks"
)

type Beam struct {
	standingHead *ebiten.Image
	standingTail []*ebiten.Image
	standingOpts *ebiten.DrawImageOptions

	tasks tasks.Tasks

	tailIndex  int
	facingLeft bool
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

	b.standingOpts = &ebiten.DrawImageOptions{}

	b.tasks.Add(tasks.Tick(10, func() error {
		b.tailIndex++
		b.tailIndex %= 11
		b.dirty = true
		return nil
	}))

	b.dirty = true
	b.img = ebiten.NewImage(244, 235)

	b.Options = &ebiten.DrawImageOptions{}

	return b, nil
}

func (b *Beam) Draw(dst *ebiten.Image) {
	if b.dirty {
		b.img.Clear()

		b.standingOpts.GeoM.Reset()
		if b.facingLeft {
			b.standingOpts.GeoM.Scale(-1, 1)
			b.standingOpts.GeoM.Translate(244, 0)
		}

		b.img.DrawImage(b.standingHead, b.standingOpts)
		b.img.DrawImage(b.standingTail[b.tailIndex], b.standingOpts)

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
	b.dirty = true
}

func (b *Beam) Bounds() image.Rectangle {
	return b.img.Bounds()
}
