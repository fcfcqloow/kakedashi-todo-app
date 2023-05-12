package usecase

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

type (
	TodoRepository interface {
		ListTasks() (*domain.Tasks, error)
		SyncTasks(domain.Tasks) error
	}
	TopicRepository interface {
		ListTopics() ([]domain.Topic, error)
		SyncTopics([]domain.Topic) error
	}
	SettingsRepository interface {
		GetSettings() (domain.Settings, error)
		SyncSettings(domain.Settings) error
	}
	LogRepository interface {
		GetLogs(dateName domain.Date) ([]domain.Log, error)
		SyncLogs(dateName domain.Date, logs []domain.Log) error
		ListDate() ([]domain.Date, error)
	}
	RestoreRepository interface {
		Restore([]domain.Task) error
	}
	MemoRepository interface {
		GetMemo() (string, error)
		SyncMemo(string) error
	}
)
