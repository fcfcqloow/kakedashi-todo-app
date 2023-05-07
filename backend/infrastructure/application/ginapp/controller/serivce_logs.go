package controller

import "github.com/fcfcqloow/first-todo-list/backend/usecase"

func feLogs(logs []usecase.Log) []string {
	result := []string{}
	for _, v := range logs {
		result = append(result, string(v))
	}

	return result
}

func feDateList(dates []usecase.Date) []string {
	result := []string{}
	for _, v := range dates {
		result = append(result, string(v))
	}

	return result
}
