//go:build js && wasm
// +build js,wasm

package wasmapp

import (
	"encoding/json"
	"syscall/js"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
	"github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/wasm/model"
	"github.com/fcfcqloow/first-todo-list/backend/usecase"
	cnv "github.com/fcfcqloow/go-advance/convert"
)

type (
	Controller struct {
		todoUseCase usecase.TodoUseCase
	}
)

func NewController(
	todoUseCase usecase.TodoUseCase,
) *Controller {
	return &Controller{
		todoUseCase: todoUseCase,
	}
}

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
func body[T any](value js.Value) (*T, error) {
	var result T
	if err := json.Unmarshal([]byte(value.String()), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Controller) AddTask(_ js.Value, args []js.Value) any {
	task, err := body[model.Task](args[0])
	if err != nil {
		return err
	}

	tasks, err := c.todoUseCase.AddTodo(domainTask(*task))
	if err != nil {
		return err
	}

	return cnv.MustCompactJson(tasks)
}

func (c *Controller) GetTasks(_ js.Value, args []js.Value) any {
	type Option struct {
		Limit  *int    `json:"limit"`
		Offset *int    `json:"offset"`
		Target *string `json:"target"`
		Sort   *bool   `json:"sort"`
	}

	options, err := body[Option](args[0])
	if err != nil {
		return err
	}

	useCaseOptions := []usecase.TodoOption{}
	if options.Limit != nil {
		useCaseOptions = append(useCaseOptions, usecase.TodoLimit(*options.Limit))
	}
	if options.Offset != nil {
		useCaseOptions = append(useCaseOptions, usecase.TodoOffset(*options.Offset))
	}
	if options.Target != nil {
		useCaseOptions = append(useCaseOptions, usecase.TodoTarget(*options.Target))
	}
	if options.Sort != nil {
		useCaseOptions = append(useCaseOptions, usecase.TodoSort(*options.Sort))
	}

	tasks, err := c.todoUseCase.ListTodo(useCaseOptions...)
	if err != nil {
		return err
	}

	return cnv.MustCompactJson(tasks)

}

func (c *Controller) RemoveTask(_ js.Value, args []js.Value) any {
	task, err := body[model.Task](args[0])
	if err != nil {
		return err
	}

	tasks, err := c.todoUseCase.DeleteTodo(domainTask(*task))
	if err != nil {
		return err
	}

	return cnv.MustCompactJson(tasks)
}

func (c *Controller) UpdateTask(_ js.Value, args []js.Value) any {
	task, err := body[model.Task](args[0])
	if err != nil {
		return err
	}

	tasks, err := c.todoUseCase.UpdateTodo(domainTask(*task))
	if err != nil {
		return err
	}

	return cnv.MustCompactJson(tasks)
}

func (c *Controller) MoveTask(_ js.Value, args []js.Value) any {
	type Body struct {
		From  string     `json:"from"`
		To    string     `json:"to"`
		Index int        `json:"index"`
		Task  model.Task `json:"task"`
	}

	requestBody, err := body[Body](args[0])
	if err != nil {
		return err
	}

	tasks, err := c.todoUseCase.MoveTodo(requestBody.From, requestBody.To, requestBody.Index, domainTask(requestBody.Task))
	if err != nil {
		return err
	}

	return cnv.MustCompactJson(tasks)
}
