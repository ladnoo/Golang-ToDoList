package task

type List struct {
	Tasks map[string]Task
}

func NewList() *List {
	return &List{
		Tasks: make(map[string]Task),
	}
}

func (list *List) AddTask(task Task) error {
	if _, ok := list.Tasks[task.Title]; ok {
		return ErrTaskAlreadyExists
	}

	list.Tasks[task.Title] = task
	return nil
}

func (list *List) GetTasks() map[string]Task {
	template := make(map[string]Task, len(list.Tasks))

	for title, task := range list.Tasks {
		template[title] = task
	}

	return template
}

func (list *List) GetNotCompletedTasks() map[string]Task {
	template := make(map[string]Task)
	for title, task := range list.Tasks {
		if task.Completed {
			template[title] = task
		}
	}

	return template
}

func (list *List) CompleteTask(title string) error {
	task, ok := list.Tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	task.Complete()

	list.Tasks[task.Title] = task

	return nil
}

func (list *List) DeleteTask(title string) error {
	_, ok := list.Tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(list.Tasks, title)
	return nil
}
