package httpserver

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (task *TaskDTO) ValidateToCreate() error {
	if task.Title == "" {
		return errors.New("title is required")
	}

	if task.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

type CompleteTaskDTO struct {
	Complete bool
}

type ErrorDTO struct {
	Error string    `json:"error"`
	Time  time.Time `json:"time"`
}

func NewErrorDTO(err error) ErrorDTO {
	return ErrorDTO{
		Error: err.Error(),
		Time:  time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}

	return string(b)
}
