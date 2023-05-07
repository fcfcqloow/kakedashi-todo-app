package domain

import (
	"errors"
	"sort"

	domainerr "github.com/fcfcqloow/first-todo-list/backend/domain/errors"
)

type (
	Tasks   struct{ Todo, Pending, Processing, Done []Task }
	PickKey string
)

const (
	TodoPickKey       PickKey = "todo"
	ProcessingPickKey PickKey = "processing"
	PendingPickKey    PickKey = "pending"
	DonePickKey       PickKey = "done"
)

func (tasks *Tasks) overwriteTask(isTarget func(task Task) bool, higherOrder func(task Task) *Task) error {
	for _, taskArr := range []*[]Task{&tasks.Todo, &tasks.Processing, &tasks.Pending, &tasks.Done} {
		for j, v := range *taskArr {
			if isTarget(v) {
				newTask := higherOrder((*taskArr)[j])
				if newTask != nil {
					(*taskArr)[j] = *newTask
				} else {
					(*taskArr) = remove(*taskArr, j)
				}
				return nil
			}
		}
	}
	return &domainerr.NotFoundError{}
}
func (tasks *Tasks) overwriteTaskById(id string, higherOrder func(task Task) *Task) error {
	return tasks.overwriteTask(
		func(task Task) bool { return task.Id == id },
		higherOrder,
	)
}

func (tasks *Tasks) ReplaceTask(id string, newTask Task) error {
	return tasks.overwriteTaskById(id, func(task Task) *Task {
		return &newTask
	})
}

func (tasks *Tasks) DeleteTask(id string) error {
	return tasks.overwriteTaskById(id, func(task Task) *Task {
		return nil
	})
}

func (tasks *Tasks) SortPriority() error {
	sort.Slice(tasks.Todo, func(i, j int) bool {
		return tasks.Todo[i].Priority.Value > tasks.Todo[j].Priority.Value
	})

	sort.SliceStable(tasks.Todo, func(i, j int) bool {
		return tasks.Todo[i].DeadLine/(1000*60*60*8) < tasks.Todo[j].DeadLine/(1000*60*60*8)
	})

	return nil
}

func (tasks *Tasks) MoveTask(from, to PickKey, index int, task Task) error {
	getList := func(key PickKey) (*[]Task, error) {
		list, ok := map[PickKey]*[]Task{
			TodoPickKey:       &tasks.Todo,
			ProcessingPickKey: &tasks.Processing,
			PendingPickKey:    &tasks.Pending,
			DonePickKey:       &tasks.Done,
		}[key]
		if !ok {
			return nil, &domainerr.NotFoundError{}
		}

		return list, nil
	}

	fromList, err := getList(from)
	if err != nil {
		return err
	}

	if to != from {
		toList, err := getList(to)
		if err != nil {
			return err
		}

		for i, v := range *fromList {
			if v.Id == task.Id {
				*fromList = remove(*fromList, i)
				break
			}
		}

		if len(*toList) == 0 {
			*toList = append(*toList, task)
			return nil
		}

		if len(*toList) <= index {
			return &domainerr.ListOverError{}
		}

		*toList = insert(*toList, index, task)
		return nil
	}

	fromIndex := findIndex(*fromList, func(v Task) bool { return v.Id == task.Id })
	if fromIndex == -1 {
		return errors.New("from same to index")
	}

	if len(*fromList) <= index {
		return &domainerr.ListOverError{}
	}

	(*fromList)[fromIndex], (*fromList)[index] = (*fromList)[index], (*fromList)[fromIndex]

	return nil
}

func (tasks *Tasks) Pick(key PickKey) error {
	var tmp Tasks
	switch key {
	case TodoPickKey:
		tmp.Todo = tasks.Todo
	case ProcessingPickKey:
		tmp.Processing = tasks.Processing
	case PendingPickKey:
		tmp.Pending = tasks.Pending
	case DonePickKey:
		tmp.Done = tasks.Done
	default:
		return &domainerr.NotFoundError{}
	}

	*tasks = tmp

	return nil
}

func (tasks *Tasks) ClearTopic(topicId string) error {
	return tasks.overwriteTask(
		func(task Task) bool {
			return findIndex(task.Topics, func(_topicId string) bool { return _topicId == topicId }) >= 0
		},
		func(task Task) *Task {
			index := findIndex(task.Topics, func(_topicId string) bool { return _topicId == topicId })
			task.Topics = remove(task.Topics, index)
			return &task
		},
	)
}
