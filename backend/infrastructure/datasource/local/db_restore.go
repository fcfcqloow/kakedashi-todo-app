package ldb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

func restoreFileName(files []string) string {
	now := time.Now()
	return filepath.Join(RestoreDir, fmt.Sprintf("%d_%d_%d_%d.json", len(files), now.Year(), now.Month(), now.Day()))
}

func (t *localRepository) Restore(tasks []domain.Task) error {
	files, err := listFiles(RestoreDir)
	if err != nil {
		return err
	}

	byts, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(restoreFileName(files), byts, os.ModePerm)
}

func (t *localRepository) ListRestore() ([]string, error) {
	return listFiles(RestoreDir)
}

func (t *localRepository) GetRestore(path string) ([]byte, error) {
	return os.ReadFile(path)
}
