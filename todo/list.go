package todo

import (
	"sync"
)

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (list *List) AddTask(task Task) error {
	list.mtx.Lock()
	defer list.mtx.Unlock()

	if _, ok := list.tasks[task.Title]; ok {

		return ErrTaskAlreadyExists
	}

	list.tasks[task.Title] = task
	return nil
}

func (list *List) GetTask(title string) (Task, error) {
	list.mtx.RLock()
	defer list.mtx.RUnlock()
	task, ok := list.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return task, nil
}

func (list *List) GetTasks() map[string]Task {
	list.mtx.RLock()
	defer list.mtx.RUnlock()
	template := make(map[string]Task, len(list.tasks))

	for title, task := range list.tasks {
		template[title] = task
	}

	return template
}

func (list *List) GetUncompletedTasks() map[string]Task {
	list.mtx.RLock()
	defer list.mtx.RUnlock()

	template := make(map[string]Task)
	for title, task := range list.tasks {
		if task.Completed == false {
			template[title] = task
		}
	}

	return template
}

func (list *List) CompleteTask(title string) (Task, error) {
	list.mtx.Lock()
	defer list.mtx.Unlock()

	task, ok := list.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.Complete()

	list.tasks[task.Title] = task

	return list.tasks[title], nil
}

func (list *List) UncompleteTask(title string) (Task, error) {
	list.mtx.Lock()
	defer list.mtx.Unlock()

	task, ok := list.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.Uncomplete()

	list.tasks[task.Title] = task

	return list.tasks[title], nil
}

func (list *List) DeleteTask(title string) error {
	list.mtx.Lock()
	defer list.mtx.Unlock()

	_, ok := list.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(list.tasks, title)
	return nil
}
