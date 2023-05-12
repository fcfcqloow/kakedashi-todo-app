package ldb

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

var settingsCacheKey = "settings-cache-key"

func (t *localRepository) GetSettings() (domain.Settings, error) {
	settings, err := load[domain.Settings](SettingDbFile, settingsCacheKey)
	if err != nil {
		return domain.Settings{}, err
	}

	return *settings, nil
}

func (t *localRepository) SyncSettings(settings domain.Settings) error {
	if err := save(SettingDbFile, settings, settingsCacheKey); err != nil {
		return err
	}

	return nil
}
