package ldb

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
	"github.com/fcfcqloow/first-todo-list/backend/usecase"
)

var settingsCacheKey = "settings-cache-key"

type settingsDB struct {
	localPath string
}

func NewSettingsRepository(localPath string) usecase.SettingsRepository {
	return &settingsDB{
		localPath: localPath,
	}
}

func (t *settingsDB) Get() (domain.Settings, error) {
	settings, err := load[domain.Settings](t.localPath, settingsCacheKey)
	if err != nil {
		return domain.Settings{}, err
	}

	return *settings, nil
}

func (t *settingsDB) Sync(settings domain.Settings) error {
	if err := save(t.localPath, settings, settingsCacheKey); err != nil {
		return err
	}

	return nil
}
