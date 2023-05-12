package dsdummy

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
	"github.com/fcfcqloow/first-todo-list/backend/usecase"
)

type (
	dummyTodoRepository struct {
		getData  func() (*domain.Tasks, error)
		saveData func(domain.Tasks) error
	}
)

func NewTodoRepository(
	getData func() (*domain.Tasks, error),
	saveData func(domain.Tasks) error,
) usecase.TodoRepository {
	instance := &dummyTodoRepository{
		getData:  getData,
		saveData: saveData,
	}
	return instance
}

func (d *dummyTodoRepository) ListTasks() (*domain.Tasks, error) {
	return d.getData()
}

func (d *dummyTodoRepository) SyncTasks(tasks domain.Tasks) error {
	return d.saveData(tasks)
}
