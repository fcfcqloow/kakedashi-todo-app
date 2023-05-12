package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

type (
	Summary struct {
		Start int
		End   int
		Logs  []domain.Log
		Date  []domain.Date
	}
	LogUseCase interface {
		Append(text string) ([]domain.Log, error)
		List(dirName domain.Date) ([]domain.Log, error)
		ListDate(...LogOption) ([]domain.Date, error)
		CreateSummary(start, end int) (*Summary, error)
	}
	logUseCase struct {
		nowDateName domain.Date
		repository  LogRepository
	}
)

const YYYYMMDD = "2006/01/02"

func DateName() domain.Date {
	return domain.Date(strings.ReplaceAll(time.Now().Format(YYYYMMDD), "/", "_"))
}

func NewLogUseCase(repository LogRepository) LogUseCase {
	return &logUseCase{
		nowDateName: DateName(),
		repository:  repository,
	}
}

func (l *logUseCase) Append(text string) ([]domain.Log, error) {
	if DateName() != l.nowDateName {
		l.nowDateName = DateName()
	}

	logs, err := l.repository.GetLogs(l.nowDateName)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs for append: %w", err)
	}

	logs = append(logs, domain.Log(strings.ReplaceAll(text, "\n", "")))
	if err := l.repository.SyncLogs(l.nowDateName, logs); err != nil {
		return nil, fmt.Errorf("failed to sync logs for append: %w", err)
	}

	return logs, nil
}
func (l *logUseCase) List(dirName domain.Date) ([]domain.Log, error) {
	logs, err := l.repository.GetLogs(dirName)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	return logs, nil
}

func (l *logUseCase) ListDate(options ...LogOption) ([]domain.Date, error) {
	option := LogOptions{}
	for _, opt := range options {
		opt(&option)
	}

	paths, err := l.repository.ListDate()
	if err != nil {
		return nil, fmt.Errorf("failed to list date: %w", err)
	}

	return domain.FilterYearMonth(paths, option.year, option.month), nil
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
