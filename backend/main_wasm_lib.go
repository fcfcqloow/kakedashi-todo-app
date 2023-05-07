//go:build ignore

package main

import (
	"encoding/json"
	"syscall/js"

	wasmapp "github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/wasm"
	dsdummy "github.com/fcfcqloow/first-todo-list/backend/infrastructure/datasource/dummy"
	"github.com/fcfcqloow/first-todo-list/backend/usecase"
)

var (
	tasks    = usecase.Tasks{}
	todoRepo = dsdummy.NewTodoRepository(
		func() (*usecase.Tasks, error) { return &tasks, nil },
		func(t usecase.Tasks) error { return nil },
	)
	todoUseCase = usecase.NewTodoUseCase(todoRepo)
)

func main() {
	controller := wasmapp.NewController(todoUseCase)
	funcs := map[string]js.Func{
		"preload": js.FuncOf(func(_ js.Value, args []js.Value) any {
			return json.Unmarshal([]byte(args[0].String()), &tasks)
		}),
		"addTask":    js.FuncOf(controller.AddTask),
		"updateTask": js.FuncOf(controller.UpdateTask),
		"removeTask": js.FuncOf(controller.RemoveTask),
		"moveTask":   js.FuncOf(controller.MoveTask),
		"getTasks":   js.FuncOf(controller.GetTasks),
	}
	c := make(chan struct{}, 0)

	for name, fun := range funcs {
		js.Global().Set(name, fun)
	}
	<-c
}
