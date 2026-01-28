package models

import "time"

type TaskModel struct {
	Id          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
