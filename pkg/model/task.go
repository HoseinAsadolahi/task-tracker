package model

import "time"

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(id int, desc string) *Task {
	return &Task{
		Id:          id,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
