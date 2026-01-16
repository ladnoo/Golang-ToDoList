package todo

import "time"

type Task struct {
	Title           string
	Description     string
	Completed       bool
	CreateData      time.Time
	CompleteData    *time.Time
	ErrTaskNotFound error
}

func NewTask(title, description string) Task {
	return Task{
		Title:        title,
		Description:  description,
		Completed:    false,
		CreateData:   time.Now(),
		CompleteData: nil,
	}
}

func (task Task) Complete() {
	finishData := time.Now()

	task.Completed = true
	task.CompleteData = &finishData
}

func (task Task) Uncomplete() {
	task.Completed = false
	task.CompleteData = nil
}
