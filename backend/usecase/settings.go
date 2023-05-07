package usecase

import (
	"fmt"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

type (
	SettingsUseCase interface {
		GetSettings() (domain.Settings, error)
		UpdateSettings(domain.Settings) (domain.Settings, error)
	}
	settingsUseCase struct {
		repository SettingsRepository
	}
)

func NewSettingsUseCase(repository SettingsRepository) SettingsUseCase {
	return &settingsUseCase{
		repository: repository,
	}
}

func (s *settingsUseCase) GetSettings() (domain.Settings, error) {
	settings, err := s.repository.Get()
	if err != nil {
		return domain.Settings{}, fmt.Errorf("failed to get settings: %w", err)
	}

	return settings, nil
}

func (s *settingsUseCase) UpdateSettings(settings domain.Settings) (domain.Settings, error) {
	err := s.repository.Sync(settings)
	if err != nil {
		return domain.Settings{}, fmt.Errorf("failed to update settings: %w", err)
	}

	return settings, nil
}
