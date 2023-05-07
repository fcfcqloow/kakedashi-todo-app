package controller

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
	"github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/ginapp/model"
)

func domainTask(feTask model.Task) domain.Task {
	return domain.Task{
		Id:          feTask.Id,
		Value:       feTask.Value,
		CreatedDate: feTask.Created,
		DeadLine:    feTask.DeadLine,
		Priority:    domain.PtrToNullInt(feTask.Priority),
		DoneDate:    domain.PtrToNullInt(feTask.DoneDate),
		Topics:      feTask.Topics,
	}
}

func feTask(task domain.Task) model.Task {
	return model.Task{
		Id:       task.Id,
		Value:    task.Value,
		Created:  task.CreatedDate,
		DeadLine: task.DeadLine,
		Priority: &task.Priority.Value,
		DoneDate: &task.DoneDate.Value,
		Topics:   task.Topics,
	}
}

func FeTasks(tasks []domain.Task) []model.Task {
	result := []model.Task{}
	for _, task := range tasks {
		result = append(result, feTask(task))
	}

	return result
}

func feResponse(tasks domain.Tasks) model.TasksResponse {
	return model.TasksResponse{
		Todo:       FeTasks(tasks.Todo),
		Processing: FeTasks(tasks.Processing),
		Pending:    FeTasks(tasks.Pending),
		Done:       FeTasks(tasks.Done),
	}
}
