package task

import "time"

type Task struct {
	Title        string
	Description  string
	Completed    bool
	CreateData   time.Time
	CompleteData *time.Time
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
