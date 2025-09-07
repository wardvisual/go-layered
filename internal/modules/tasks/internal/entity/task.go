package entity

import "github.com/google/uuid"

type Task struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
	Done  bool      `json:"done" db:"done"`
}
