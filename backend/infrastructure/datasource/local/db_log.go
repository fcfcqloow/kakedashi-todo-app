package ldb

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

func filePath(dateName domain.Date) string {
	result := filepath.Join(LogDir, string(dateName))
	if !strings.Contains(string(dateName), ".log") {
		result += ".log"
	}
	safeCreateFile(result, ptrString(""))
	return result
}

func (l *localRepository) SyncLogs(dateName domain.Date, logs []domain.Log) error {
	stringArr := []string{}
	for _, log := range logs {
		stringArr = append(stringArr, string(log))
	}

	if err := os.WriteFile(filePath(dateName), []byte(strings.Join(stringArr, "\n")), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (l *localRepository) GetLogs(path domain.Date) ([]domain.Log, error) {
	byts, err := os.ReadFile(filePath(path))
	if err != nil {
		return nil, err
	}

	result := []domain.Log{}
	for _, v := range strings.Split(string(byts), "\n") {
		result = append(result, domain.Log(v))
	}

	return result, nil
}

func (l *localRepository) ListDate() ([]domain.Date, error) {
	files, err := listFiles(LogDir)
	if err != nil {
		return nil, err
	}

	result := []domain.Date{}
	for _, v := range files {
		if strings.Contains(v, ".log") {
			result = append(result, domain.Date(v))
		}
	}

	return result, nil
}
