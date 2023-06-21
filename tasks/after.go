package tasks

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type AfterTask struct {
	ticks    int
	maxTicks int
	done     bool
	task     func() error
}

func After(delay time.Duration, task func() error) Task {
	return &AfterTask{
		maxTicks: int(delay.Seconds() * float64(ebiten.TPS())),
		task:     task,
	}
}

func (t *AfterTask) Done() bool {
	return t.done
}

func (t *AfterTask) Update() error {
	if t.done {
		return nil
	}
	t.ticks++
	if t.done = t.ticks == t.maxTicks; t.done {
		return t.task()
	}
	return nil
}

func (t *AfterTask) Cancel() {
	t.done = true
}
