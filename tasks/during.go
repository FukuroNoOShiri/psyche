package tasks

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type DuringTask struct {
	ticks    int
	maxTicks int
	done     bool
	task     func(progression float64) error
}

func During(delay time.Duration, task func(progression float64) error) Task {
	return &DuringTask{
		maxTicks: int(delay.Seconds() * float64(ebiten.TPS())),
		task:     task,
	}
}

func (t *DuringTask) Done() bool {
	return t.done
}

func (t *DuringTask) Update() error {
	if t.done {
		return nil
	}
	t.ticks++
	t.done = t.ticks == t.maxTicks
	return t.task(float64(t.ticks) / float64(t.maxTicks))
}

func (t *DuringTask) Cancel() {
	t.done = true
}
