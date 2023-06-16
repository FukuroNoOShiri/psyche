package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
)

func IsSomeKeyJustPressed(keys ...ebiten.Key) (ok bool, key ebiten.Key) {
	key, ok = lo.Find(keys, func(key ebiten.Key) bool {
		return inpututil.IsKeyJustPressed(key)
	})
	return
}
