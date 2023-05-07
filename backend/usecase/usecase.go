package usecase

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

type (
	TodoRepository interface {
		List() (*domain.Tasks, error)
		Sync(domain.Tasks) error
	}
	TopicRepository interface {
		List() ([]domain.Topic, error)
		Sync([]domain.Topic) error
	}
	SettingsRepository interface {
		Get() (domain.Settings, error)
		Sync(domain.Settings) error
	}
	LogRepository interface {
		Get(dateName Date) ([]Log, error)
		Sync(dateName Date, logs []Log) error
		ListDate() ([]Date, error)
	}
)
