package tasks

import (
	"github.com/samber/lo"
)

type Tasks struct {
	tasks []taggedTask
}

type taggedTask struct {
	Task
	tags []string
}

func (t *Tasks) Add(task Task, tags ...string) {
	t.tasks = append(t.tasks, taggedTask{
		Task: task,
		tags: tags,
	})
}

func (t *Tasks) Update() error {
	for _, task := range t.tasks {
		if err := task.Update(); err != nil {
			return err
		}
	}

	t.tasks = lo.Reject(t.tasks, func(task taggedTask, _ int) bool {
		return task.Done()
	})

	return nil
}

func (t *Tasks) Cancel(tag string) {
	for _, task := range lo.Filter(t.tasks, func(task taggedTask, _ int) bool {
		return lo.Contains(task.tags, tag)
	}) {
		task.Cancel()
	}

	t.tasks = lo.Reject(t.tasks, func(task taggedTask, _ int) bool {
		return task.Done()
	})
}

type Task interface {
	Done() bool
	Update() error
	Cancel()
}
