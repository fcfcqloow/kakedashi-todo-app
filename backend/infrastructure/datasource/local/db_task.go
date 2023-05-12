package ldb

import (
	"sync"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

var (
	lock     sync.Mutex
	cacheKey = "tasks"
)

func (d *localRepository) ListTasks() (*domain.Tasks, error) {
	tasks, err := load[domain.Tasks](TodoDbFile, cacheKey)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (d *localRepository) SyncTasks(tasks domain.Tasks) error {
	if err := save(TodoDbFile, tasks, cacheKey); err != nil {
		return err
	}

	return nil
}
