package usecase

import "fmt"

type (
	MemoUseCase interface {
		GetMemo() (string, error)
		UpdateMemo(value string) (string, error)
	}
	memo struct {
		repository MemoRepository
	}
)

func NewMemoUseCase(repository MemoRepository) MemoUseCase {
	return &memo{
		repository: repository,
	}
}

func (m *memo) GetMemo() (string, error) {
	value, err := m.repository.GetMemo()
	if err != nil {
		return "", fmt.Errorf("failed to get memo: %w", err)
	}

	return value, nil
}

func (m *memo) UpdateMemo(value string) (string, error) {
	err := m.repository.SyncMemo(value)
	if err != nil {
		return "", fmt.Errorf("failed to update memo: %w", err)
	}

	return value, nil
}
