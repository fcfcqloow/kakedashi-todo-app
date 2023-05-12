package controller

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

func feLogs(logs []domain.Log) []string {
	result := []string{}
	for _, v := range logs {
		result = append(result, string(v))
	}

	return result
}

func feDateList(dates []domain.Date) []string {
	result := []string{}
	for _, v := range dates {
		result = append(result, string(v))
	}

	return result
}
