package domain_test

import (
	"testing"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
	domainerr "github.com/fcfcqloow/first-todo-list/backend/domain/errors"
	"github.com/stretchr/testify/assert"
)

func TestTasks_Pick(t *testing.T) {
	t.Parallel()
	type (
		input struct {
			key   domain.PickKey
			tasks *domain.Tasks
		}
		want struct {
			err   error
			tasks *domain.Tasks
		}
	)

	cases := map[string]struct {
		input input
		want  want
	}{
		"Success: Pick Todo": {
			input: input{
				tasks: &domain.Tasks{
					Todo:       []domain.Task{{Id: "1"}},
					Processing: []domain.Task{{Id: "2"}},
					Pending:    []domain.Task{{Id: "3"}},
					Done:       []domain.Task{{Id: "4"}},
				},
				key: domain.TodoPickKey,
			},
			want: want{
				err:   nil,
				tasks: &domain.Tasks{Todo: []domain.Task{{Id: "1"}}},
			},
		},
		"Success: Pick Processing": {
			input: input{
				tasks: &domain.Tasks{
					Todo:       []domain.Task{{Id: "1"}},
					Processing: []domain.Task{{Id: "2"}},
					Pending:    []domain.Task{{Id: "3"}},
					Done:       []domain.Task{{Id: "4"}},
				},
				key: domain.ProcessingPickKey,
			},
			want: want{
				err:   nil,
				tasks: &domain.Tasks{Processing: []domain.Task{{Id: "2"}}},
			},
		},
		"Success: Pick `Pending`": {
			input: input{
				tasks: &domain.Tasks{
					Todo:       []domain.Task{{Id: "1"}},
					Processing: []domain.Task{{Id: "2"}},
					Pending:    []domain.Task{{Id: "3"}},
					Done:       []domain.Task{{Id: "4"}},
				},
				key: domain.PendingPickKey,
			},
			want: want{
				err:   nil,
				tasks: &domain.Tasks{Pending: []domain.Task{{Id: "3"}}},
			},
		},
		"Success: Pick `Done`": {
			input: input{
				tasks: &domain.Tasks{
					Todo:       []domain.Task{{Id: "1"}},
					Processing: []domain.Task{{Id: "2"}},
					Pending:    []domain.Task{{Id: "3"}},
					Done:       []domain.Task{{Id: "4"}},
				},
				key: domain.DonePickKey,
			},
			want: want{
				err:   nil,
				tasks: &domain.Tasks{Done: []domain.Task{{Id: "4"}}},
			},
		},
		"Error: Not Found Error": {
			input: input{
				tasks: &domain.Tasks{
					Todo:       []domain.Task{{Id: "1"}},
					Processing: []domain.Task{{Id: "2"}},
					Pending:    []domain.Task{{Id: "3"}},
					Done:       []domain.Task{{Id: "4"}},
				},
				key: domain.PickKey(""),
			},
			want: want{
				err: &domainerr.NotFoundError{},
				tasks: &domain.Tasks{
					Todo:       []domain.Task{{Id: "1"}},
					Processing: []domain.Task{{Id: "2"}},
					Pending:    []domain.Task{{Id: "3"}},
					Done:       []domain.Task{{Id: "4"}},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var tmp *domain.Tasks
			if tc.input.tasks != nil {
				tmp = &domain.Tasks{}
				*tmp = *tc.input.tasks
			}

			err := tmp.Pick(tc.input.key)
			assert.Equal(t, tc.want.tasks, tmp)
			testError(t, tc.want.err, err)
		})
	}
}
func TestTasks_SortPriority(t *testing.T) {
	t.Parallel()
	type (
		input struct {
			tasks *domain.Tasks
		}
		want struct {
			err   error
			tasks *domain.Tasks
		}
	)

	cases := map[string]struct {
		input input
		want  want
	}{
		"Success: Empty Task": {
			input: input{
				tasks: &domain.Tasks{},
			},
			want: want{
				err:   nil,
				tasks: &domain.Tasks{},
			},
		},
		"Success: Tasks Sorted Correctly": {
			input: input{
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Priority: domain.NullableInteger{Value: 1}, DeadLine: 1620000000000}, // 2021-05-03 00:00:00
						{Id: "2", Priority: domain.NullableInteger{Value: 3}, DeadLine: 1620144000000}, // 2021-05-04 12:00:00
						{Id: "3", Priority: domain.NullableInteger{Value: 2}, DeadLine: 1620176400000}, // 2021-05-05 08:00:00
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Priority: domain.NullableInteger{Value: 1}, DeadLine: 1620000000000}, // 2021-05-03 00:00:00
						{Id: "2", Priority: domain.NullableInteger{Value: 3}, DeadLine: 1620144000000}, // 2021-05-04 12:00:00
						{Id: "3", Priority: domain.NullableInteger{Value: 2}, DeadLine: 1620176400000}, // 2021-05-05 08:00:00
					},
				},
			},
		},
		"Success: Multiple Tasks with Same Priority": {
			input: input{
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "5", Priority: domain.NullableInteger{Value: 3}, DeadLine: 1620332400000}, // 2021-05-07 08:00:00
						{Id: "2", Priority: domain.NullableInteger{Value: 1}, DeadLine: 1620144000000}, // 2021-05-04 12:00:00
						{Id: "4", Priority: domain.NullableInteger{Value: 1}, DeadLine: 1620208800000}, // 2021-05-05 17:00:00
						{Id: "3", Priority: domain.NullableInteger{Value: 2}, DeadLine: 1620176400000}, // 2021-05-05 08:00:00
						{Id: "1", Priority: domain.NullableInteger{Value: 2}, DeadLine: 1620000000000}, // 2021-05-03 00:00:00
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Priority: domain.NullableInteger{Value: 2}, DeadLine: 1620000000000}, // 2021-05-03 00:00:00
						{Id: "2", Priority: domain.NullableInteger{Value: 1}, DeadLine: 1620144000000}, // 2021-05-04 12:00:00
						{Id: "3", Priority: domain.NullableInteger{Value: 2}, DeadLine: 1620176400000}, // 2021-05-05 08:00:00
						{Id: "4", Priority: domain.NullableInteger{Value: 1}, DeadLine: 1620208800000}, // 2021-05-05 17:00:00
						{Id: "5", Priority: domain.NullableInteger{Value: 3}, DeadLine: 1620332400000}, // 2021-05-07 08:00:00
					},
				},
			},
		},
		"Success: Tasks with Null Priority": {
			input: input{
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "3", Priority: domain.NullableInteger{Value: 2}, DeadLine: 1620176400000},     // 2021-05-05 08:00:00
						{Id: "1", Priority: domain.NullableInteger{IsNull: true}, DeadLine: 1620000000000}, // 2021-05-03 00:00:00
						{Id: "2", Priority: domain.NullableInteger{Value: 1}, DeadLine: 1620144000000},     // 2021-05-04 12:00:00
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Priority: domain.NullableInteger{IsNull: true}, DeadLine: 1620000000000}, // 2021-05-03 00:00:00
						{Id: "2", Priority: domain.NullableInteger{Value: 1}, DeadLine: 1620144000000},     // 2021-05-04 12:00:00
						{Id: "3", Priority: domain.NullableInteger{Value: 2}, DeadLine: 1620176400000},     // 2021-05-05 08:00:00
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var tmp *domain.Tasks
			if tc.input.tasks != nil {
				tmp = &domain.Tasks{}
				*tmp = *tc.input.tasks
			}

			err := tmp.SortPriority()
			assert.Equal(t, tc.want.tasks, tmp)
			testError(t, tc.want.err, err)
		})
	}
}

func TestTasks_ReplaceTask(t *testing.T) {
	t.Parallel()
	type (
		input struct {
			tasks   *domain.Tasks
			newTask *domain.Task
			id      string
		}
		want struct {
			err   error
			tasks *domain.Tasks
		}
	)
	cases := map[string]struct {
		input input
		want  want
	}{
		"Success: Replace Task": {
			input: input{
				id: "1",
				newTask: &domain.Task{
					Id:    "1",
					Value: "Updated Task",
				},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Updated Task"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
		},
		"Success: Replace Task in Pending Tasks": {
			input: input{
				id: "2",
				newTask: &domain.Task{
					Id:    "2",
					Value: "Updated Task",
				},
				tasks: &domain.Tasks{
					Pending: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Pending: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Updated Task"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
		},
		"Success: Replace Task in Processing Tasks": {
			input: input{
				id: "3",
				newTask: &domain.Task{
					Id:    "3",
					Value: "Updated Task",
				},
				tasks: &domain.Tasks{
					Processing: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Processing: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Updated Task"},
					},
				},
			},
		},
		"Error: Task not found": {
			input: input{
				id: "4",
				newTask: &domain.Task{
					Id:    "4",
					Value: "New Task",
				},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: &domainerr.NotFoundError{},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
		},
		"Error: Empty ID": {
			input: input{
				id: "",
				newTask: &domain.Task{
					Id:    "4",
					Value: "New Task",
				},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: &domainerr.NotFoundError{},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
		},
		"Error: Empty Task": {
			input: input{
				id:      "",
				newTask: &domain.Task{},
				tasks:   &domain.Tasks{},
			},
			want: want{
				err:   &domainerr.NotFoundError{},
				tasks: &domain.Tasks{},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var tmp *domain.Tasks
			if tc.input.tasks != nil {
				tmp = &domain.Tasks{}
				*tmp = *tc.input.tasks
			}

			err := tmp.ReplaceTask(tc.input.id, *tc.input.newTask)
			assert.Equal(t, tc.want.tasks, tmp)
			testError(t, tc.want.err, err)
		})
	}
}

func TestTasks_DeleteTask(t *testing.T) {
	t.Parallel()
	type (
		input struct {
			tasks *domain.Tasks
			id    string
		}
		want struct {
			err   error
			tasks *domain.Tasks
		}
	)
	cases := map[string]struct {
		input input
		want  want
	}{
		"Success: Delete Task from Todo": {
			input: input{
				id: "1",
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
		},
		"Success: Delete Task from Pending Tasks": {
			input: input{
				id: "2",
				tasks: &domain.Tasks{
					Pending: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Pending: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
		},
		"Success: Delete Task from Processing Tasks": {
			input: input{
				id: "3",
				tasks: &domain.Tasks{
					Processing: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Processing: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
					},
				},
			},
		},
		"Success: Task not found in any task list (Todo, Pending, Processing, Completed)": {
			input: input{
				id: "4",
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
					Pending: []domain.Task{
						{Id: "4", Value: "Task 4"},
					},
					Processing: []domain.Task{},
					Done:       []domain.Task{},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
					Pending:    []domain.Task{},
					Processing: []domain.Task{},
					Done:       []domain.Task{},
				},
			},
		},
		"Success: Delete Task from Completed Tasks": {
			input: input{
				id: "3",
				tasks: &domain.Tasks{
					Done: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Done: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
					},
				},
			},
		},
		"Error: Task not found": {
			input: input{
				id: "4",
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: &domainerr.NotFoundError{},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
		},
		"Error: Empty ID provided": {
			input: input{
				id: "",
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: &domainerr.NotFoundError{},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var tmp *domain.Tasks
			if tc.input.tasks != nil {
				tmp = &domain.Tasks{}
				*tmp = *tc.input.tasks
			}

			err := tmp.DeleteTask(tc.input.id)
			assert.Equal(t, tc.want.tasks, tmp)
			testError(t, tc.want.err, err)
		})
	}
}

func TestTasks_MoveTask(t *testing.T) {
	type input struct {
		tasks *domain.Tasks
		from  domain.PickKey
		to    domain.PickKey
		index int
		task  domain.Task
	}

	type want struct {
		err   error
		tasks *domain.Tasks
	}

	cases := map[string]struct {
		input input
		want  want
	}{
		"Success: Move Task from Todo to Processing": {
			input: input{
				from:  domain.TodoPickKey,
				to:    domain.ProcessingPickKey,
				index: 1,
				task: domain.Task{
					Id:    "2",
					Value: "Task 2",
				},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
					Processing: []domain.Task{},
					Pending:    []domain.Task{},
					Done:       []domain.Task{},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "3", Value: "Task 3"},
					},
					Processing: []domain.Task{
						{Id: "2", Value: "Task 2"},
					},
					Pending: []domain.Task{},
					Done:    []domain.Task{},
				},
			},
		},
		"Success: Move Task from Processing to Pending": {
			input: input{
				from:  domain.ProcessingPickKey,
				to:    domain.PendingPickKey,
				index: 0,
				task: domain.Task{
					Id:    "2",
					Value: "Task 2",
				},
				tasks: &domain.Tasks{
					Todo: []domain.Task{},
					Processing: []domain.Task{
						{Id: "2", Value: "Task 2"},
					},
					Pending: []domain.Task{},
					Done:    []domain.Task{},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Todo:       []domain.Task{},
					Processing: []domain.Task{},
					Pending: []domain.Task{
						{Id: "2", Value: "Task 2"},
					},
					Done: []domain.Task{},
				},
			},
		},
		"Success: Move Task from Done to Todo": {
			input: input{
				from:  domain.DonePickKey,
				to:    domain.TodoPickKey,
				index: 0,
				task:  domain.Task{Id: "1", Value: "Task 1"},
				tasks: &domain.Tasks{
					Todo:       []domain.Task{},
					Processing: []domain.Task{},
					Pending:    []domain.Task{},
					Done: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Todo:       []domain.Task{{Id: "1", Value: "Task 1"}},
					Processing: []domain.Task{},
					Pending:    []domain.Task{},
					Done: []domain.Task{
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
				},
			},
		},
		"Success: Move Task within Todo": {
			input: input{
				from:  domain.TodoPickKey,
				to:    domain.TodoPickKey,
				index: 2,
				task: domain.Task{
					Id:    "1",
					Value: "Task 1",
				},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "3", Value: "Task 3"},
						{Id: "2", Value: "Task 2"},
						{Id: "1", Value: "Task 1"},
					},
					Processing: []domain.Task{},
					Pending:    []domain.Task{},
					Done:       []domain.Task{},
				},
			},
			want: want{
				err: nil,
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "3", Value: "Task 3"},
						{Id: "2", Value: "Task 2"},
						{Id: "1", Value: "Task 1"},
					},
					Processing: []domain.Task{},
					Pending:    []domain.Task{},
					Done:       []domain.Task{},
				},
			},
		},
		"Failure: Invalid from PickKey": {
			input: input{
				from:  "InvalidPickKey",
				to:    domain.ProcessingPickKey,
				index: 0,
				task: domain.Task{
					Id:    "1",
					Value: "Task 1",
				},
				tasks: &domain.Tasks{},
			},
			want: want{
				err:   &domainerr.NotFoundError{},
				tasks: &domain.Tasks{},
			},
		},
		"Failure: Invalid task index": {
			input: input{
				from:  domain.TodoPickKey,
				to:    domain.ProcessingPickKey,
				index: 10,
				task:  domain.Task{Id: "1", Value: "Task 1"},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "1", Value: "Task 1"},
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
					Processing: []domain.Task{},
					Pending:    []domain.Task{},
					Done:       []domain.Task{},
				},
			},
			want: want{
				err: &domainerr.ListOverError{},
				tasks: &domain.Tasks{
					Todo: []domain.Task{
						{Id: "2", Value: "Task 2"},
						{Id: "3", Value: "Task 3"},
					},
					Processing: []domain.Task{{Id: "1", Value: "Task 1"}},
					Pending:    []domain.Task{},
					Done:       []domain.Task{},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var tmp *domain.Tasks
			if tc.input.tasks != nil {
				tmp = &domain.Tasks{}
				*tmp = *tc.input.tasks
			}

			err := tmp.MoveTask(tc.input.from, tc.input.to, tc.input.index, tc.input.task)
			assert.Equal(t, tc.want.tasks, tmp)
			testError(t, tc.want.err, err)
		})
	}
}
