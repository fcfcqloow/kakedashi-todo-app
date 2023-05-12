package usecase

import (
	"fmt"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

type (
	TodoUseCase interface {
		ListTodo(options ...TodoOption) (*domain.Tasks, error)
		AddTodo(task domain.Task) (*domain.Tasks, error)
		UpdateTodo(task domain.Task) (*domain.Tasks, error)
		DeleteTodo(task domain.Task) (*domain.Tasks, error)
		MoveTodo(from, to string, index int, task domain.Task) (*domain.Tasks, error)
		Sort() (*domain.Tasks, error)
		Restore(tasks []domain.Task) (*domain.Tasks, error)
		RestoreDone() (*domain.Tasks, error)
	}
	todoUseCase struct {
		todoRepository    TodoRepository
		restoreRepository RestoreRepository
	}
)

func NewTodoUseCase(todoRepository TodoRepository, restoreRepository RestoreRepository) TodoUseCase {
	return &todoUseCase{
		todoRepository:    todoRepository,
		restoreRepository: restoreRepository,
	}
}

func slice[T any](target []T, offset *int, limit *int) []T {
	result := target[:]

	if offset != nil {
		result = target[*offset:]
	}

	if limit != nil && len(result) > *limit {
		result = result[:*limit]
	}

	return result
}

func (u *todoUseCase) ListTodo(optionItems ...TodoOption) (*domain.Tasks, error) {
	options := applyTodoOptions(optionItems...)
	tasks, err := u.todoRepository.ListTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to get todolist: %w", err)
	}

	if options.limit != nil || options.offset != nil {
		tasks.Todo = slice(tasks.Todo, options.offset, options.limit)
		tasks.Processing = slice(tasks.Processing, options.offset, options.limit)
		tasks.Pending = slice(tasks.Pending, options.offset, options.limit)
		tasks.Done = slice(tasks.Done, options.offset, options.limit)
	}

	if options.target != nil {
		if err := tasks.Pick(domain.PickKey(*options.target)); err != nil {
			return nil, fmt.Errorf("failed to pick task ["+*options.target+"]: %w", err)
		}
	}

	return tasks, nil
}

func (u *todoUseCase) AddTodo(task domain.Task) (*domain.Tasks, error) {
	tasks, err := u.ListTodo()
	if err != nil {
		return nil, err
	}

	tasks.Todo = append(tasks.Todo, task)

	if err := u.todoRepository.SyncTasks(*tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (u *todoUseCase) UpdateTodo(task domain.Task) (*domain.Tasks, error) {
	tasks, err := u.ListTodo()
	if err != nil {
		return nil, fmt.Errorf("failed to get todolist on update: %w", err)
	}

	if err := tasks.ReplaceTask(task.Id, task); err != nil {
		return nil, fmt.Errorf("failed to update: %w", err)
	}

	if err := u.todoRepository.SyncTasks(*tasks); err != nil {
		return nil, fmt.Errorf("failed to sync on update: %w", err)
	}

	return tasks, nil
}

func (u *todoUseCase) DeleteTodo(argTask domain.Task) (*domain.Tasks, error) {
	tasks, err := u.ListTodo()
	if err != nil {
		return nil, fmt.Errorf("failed to get todolist on delete: %w", err)
	}

	if err := tasks.DeleteTask(argTask.Id); err != nil {
		return nil, fmt.Errorf("failed to delete: %w", err)
	}

	if err := u.todoRepository.SyncTasks(*tasks); err != nil {
		return nil, fmt.Errorf("failed to sync on delete : %w", err)
	}

	return tasks, nil
}

func (u *todoUseCase) MoveTodo(from, to string, index int, task domain.Task) (*domain.Tasks, error) {
	tasks, err := u.ListTodo()
	if err != nil {
		return nil, fmt.Errorf("failed to get todolist on move: %w", err)
	}

	if err := tasks.MoveTask(domain.PickKey(from), domain.PickKey(to), index, task); err != nil {
		return nil, fmt.Errorf("failed to move: %w", err)
	}

	if err := u.todoRepository.SyncTasks(*tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (u *todoUseCase) Sort() (*domain.Tasks, error) {
	tasks, err := u.todoRepository.ListTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to get todolist for sort: %w", err)
	}

	if err := tasks.SortPriority(); err != nil {
		return nil, fmt.Errorf("failed to sort tasks: %w", err)
	}

	if err := u.todoRepository.SyncTasks(*tasks); err != nil {
		return nil, fmt.Errorf("failed to sync tasks: %w", err)
	}

	return tasks, nil
}

func (t *todoUseCase) Restore(tasks []domain.Task) (*domain.Tasks, error) {
	if err := t.restoreRepository.Restore(tasks); err != nil {
		return nil, fmt.Errorf("failed to restore: %w", err)
	}

	var result *domain.Tasks
	var err error
	for _, task := range tasks {
		if result, err = t.DeleteTodo(task); err != nil {
			return nil, fmt.Errorf("failed to delete todolist of restore: %w", err)
		}

	}

	return result, nil
}

func (t *todoUseCase) RestoreDone() (*domain.Tasks, error) {
	tasks, err := t.ListTodo(TodoTarget("done"))
	if err != nil {
		return nil, fmt.Errorf("failed to get done list of restore: %w", err)
	}

	return t.Restore(tasks.Done)
}
