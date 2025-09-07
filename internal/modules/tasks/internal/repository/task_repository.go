package repository

import (
	"github.com/google/uuid"
	"github.com/wardvisual/go-layered/internal/features/tasks/internal/entity"
)

type TaskRepository interface {
	Insert(task *entity.Task) error
	Find() ([]entity.Task, error)
	FindById(id uuid.UUID) (entity.Task, error)
	Update(task *entity.Task) error
	Delete(id uuid.UUID) error
}
