package tasks

type TickTask struct {
	ticks  int
	modulo int
	done   bool
	task   func() error
}

func Tick(modulo int, task func() error) Task {
	return &TickTask{
		modulo: modulo,
		task:   task,
	}
}

func (t *TickTask) Done() bool {
	return t.done
}

func (t *TickTask) Update() error {
	if t.done {
		return nil
	}
	t.ticks++
	if t.ticks%t.modulo == 0 {
		return t.task()
	}
	return nil
}

func (t *TickTask) Cancel() {
	t.done = true
}
