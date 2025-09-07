package usecase

import (
	"github.com/google/uuid"
	"github.com/wardvisual/go-layered/internal/features/tasks/internal/entity"
	"github.com/wardvisual/go-layered/internal/features/tasks/internal/repository"
)

type TaskUseCase struct {
	Repository repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{Repository: repo}
}

func (u *TaskUseCase) Create(t *entity.Task) error {
	if t.Id == uuid.Nil {
		t.Id = uuid.New()
	}
	return u.Repository.Insert(t)
}

func (u *TaskUseCase) List() ([]entity.Task, error) {
	return u.Repository.Find()
}

func (u *TaskUseCase) Get(id uuid.UUID) (entity.Task, error) {
	return u.Repository.FindById(id)
}

func (u *TaskUseCase) Update(t *entity.Task) error {
	return u.Repository.Update(t)
}

func (u *TaskUseCase) Delete(id uuid.UUID) error {
	return u.Repository.Delete(id)
}
