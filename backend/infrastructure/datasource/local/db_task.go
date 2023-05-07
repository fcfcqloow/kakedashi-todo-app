package ldb

import (
	"sync"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
	"github.com/fcfcqloow/first-todo-list/backend/usecase"
)

var (
	lock     sync.Mutex
	cacheKey = "tasks"
)

type (
	db struct {
		localTaskPath string
	}
)

func NewTodoRepository(localTaskPath string) usecase.TodoRepository {
	instance := &db{
		localTaskPath: localTaskPath,
	}
	return instance
}

func (d *db) List() (*domain.Tasks, error) {
	tasks, err := load[domain.Tasks](d.localTaskPath, cacheKey)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (d *db) Sync(tasks domain.Tasks) error {
	if err := save(d.localTaskPath, tasks, cacheKey); err != nil {
		return err
	}

	return nil
}
