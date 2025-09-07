package inmemory

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/wardvisual/go-layered/internal/modules/todo/internal/entity"
)

type TaskRepositoryImpl struct {
	mu   sync.RWMutex
	data map[uuid.UUID]entity.Task
}

func NewTaskRepository() *TaskRepositoryImpl {
	return &TaskRepositoryImpl{data: make(map[uuid.UUID]entity.Task)}
}

func (r *TaskRepositoryImpl) Insert(task *entity.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if task.Id == uuid.Nil {
		task.Id = uuid.New()
	}
	r.data[task.Id] = *task
	return nil
}

func (r *TaskRepositoryImpl) Find() ([]entity.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := make([]entity.Task, 0, len(r.data))
	for _, v := range r.data {
		items = append(items, v)
	}
	return items, nil
}

func (r *TaskRepositoryImpl) FindById(id uuid.UUID) (entity.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if v, ok := r.data[id]; ok {
		return v, nil
	}
	return entity.Task{}, errors.New("not found")
}

func (r *TaskRepositoryImpl) Update(task *entity.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[task.Id]; !ok {
		return errors.New("not found")
	}
	r.data[task.Id] = *task
	return nil
}

func (r *TaskRepositoryImpl) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[id]; !ok {
		return errors.New("not found")
	}
	delete(r.data, id)
	return nil
}
