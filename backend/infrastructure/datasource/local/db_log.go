package ldb

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fcfcqloow/first-todo-list/backend/usecase"
)

type (
	logdb struct{}
)

func NewLogRepository() usecase.LogRepository {
	return &logdb{}
}

func filePath(dateName usecase.Date) string {
	result := filepath.Join(LogDir, string(dateName))
	if !strings.Contains(string(dateName), ".log") {
		result += ".log"
	}
	safeCreateFile(result, ptrString(""))
	return result
}

func (l *logdb) Sync(dateName usecase.Date, logs []usecase.Log) error {
	stringArr := []string{}
	for _, log := range logs {
		stringArr = append(stringArr, string(log))
	}

	if err := os.WriteFile(filePath(dateName), []byte(strings.Join(stringArr, "\n")), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (l *logdb) Get(path usecase.Date) ([]usecase.Log, error) {
	byts, err := os.ReadFile(filePath(path))
	if err != nil {
		return nil, err
	}

	result := []usecase.Log{}
	for _, v := range strings.Split(string(byts), "\n") {
		result = append(result, usecase.Log(v))
	}

	return result, nil
}

func (l *logdb) ListDate() ([]usecase.Date, error) {
	files, err := listFiles(LogDir)
	if err != nil {
		return nil, err
	}

	result := []usecase.Date{}
	for _, v := range files {
		if strings.Contains(v, ".log") {
			result = append(result, usecase.Date(v))
		}
	}

	return result, nil
}
