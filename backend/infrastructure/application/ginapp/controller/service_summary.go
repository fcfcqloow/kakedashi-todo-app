package controller

import (
	"github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/ginapp/model"
	"github.com/fcfcqloow/first-todo-list/backend/usecase"
)

func feSummary(summary usecase.Summary) model.Summary {
	return model.Summary{
		Start: summary.Start,
		End:   summary.End,
		Date:  feDateList(summary.Date),
		Logs:  feLogs(summary.Logs),
	}
}
