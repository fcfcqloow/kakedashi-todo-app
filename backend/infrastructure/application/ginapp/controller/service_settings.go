package controller

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
	"github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/ginapp/model"
)

func feSettingsResponse(settings domain.Settings) model.SettingsGetResponse {
	return model.SettingsGetResponse{
		TaskLimit:               settings.TaskLimit,
		Sortable:                settings.Sortable,
		TopicColor:              settings.TopicColor,
		Mode:                    settings.Mode,
		TopicListType:           settings.TopicListType,
		NotificationIntervalSec: settings.NotificationIntervalSec,
	}
}

func domainSettings(settingsJson model.SettingsUpdateRequest) domain.Settings {
	return domain.Settings{
		TaskLimit:               settingsJson.TaskLimit,
		Sortable:                settingsJson.Sortable,
		TopicColor:              settingsJson.TopicColor,
		Mode:                    settingsJson.Mode,
		TopicListType:           settingsJson.TopicListType,
		NotificationIntervalSec: settingsJson.NotificationIntervalSec,
	}
}
