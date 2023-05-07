package usecase

import (
	"fmt"
	"strings"
	"time"
)

type (
	Date    string
	Log     string
	Summary struct {
		Start int
		End   int
		Logs  []Log
		Date  []Date
	}
	LogUseCase interface {
		Append(text string) ([]Log, error)
		List(dirName Date) ([]Log, error)
		ListDate() ([]Date, error)
		CreateSummary(start, end int) (*Summary, error)
	}
	logUseCase struct {
		nowDateName Date
		repository  LogRepository
	}
)

const YYYYMMDD = "2006/01/02"

func DateName() Date {
	return Date(strings.ReplaceAll(time.Now().Format(YYYYMMDD), "/", "_"))
}

func NewLogUseCase(repository LogRepository) LogUseCase {
	return &logUseCase{
		nowDateName: DateName(),
		repository:  repository,
	}
}

func (l *logUseCase) Append(text string) ([]Log, error) {
	if DateName() != l.nowDateName {
		l.nowDateName = DateName()
	}

	logs, err := l.repository.Get(l.nowDateName)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs for append: %w", err)
	}

	logs = append(logs, Log(strings.ReplaceAll(text, "\n", "")))
	if err := l.repository.Sync(l.nowDateName, logs); err != nil {
		return nil, fmt.Errorf("failed to sync logs for append: %w", err)
	}

	return logs, nil
}
func (l *logUseCase) List(dirName Date) ([]Log, error) {
	logs, err := l.repository.Get(dirName)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	return logs, nil
}

func (l *logUseCase) ListDate() ([]Date, error) {
	paths, err := l.repository.ListDate()
	if err != nil {
		return nil, fmt.Errorf("failed to list date: %w", err)
	}

	return paths, nil
}

func (l *logUseCase) CreateSummary(start, end int) (*Summary, error) {
	result := &Summary{Start: start, End: end}

	dates, err := l.ListDate()
	if err != nil {
		return nil, err
	}

	for _, date := range dates {
		parsedDate, err := time.Parse(YYYYMMDD, strings.ReplaceAll(string(date), "_", "/"))
		if err != nil {
			return nil, fmt.Errorf("failed to parse of time: %w", err)
		}

		if start < int(parsedDate.UnixMilli()) && int(parsedDate.UnixMilli()) < end {
			result.Date = append(result.Date, date)

			logs, err := l.List(date)
			if err != nil {
				return nil, err
			}

			result.Logs = append(result.Logs, logs...)
		}
	}

	return result, nil
}
