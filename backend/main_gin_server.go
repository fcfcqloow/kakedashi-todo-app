//go:build ignore

package main

import (
	"os"

	"github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/ginapp"
	ldb "github.com/fcfcqloow/first-todo-list/backend/infrastructure/datasource/local"
	"github.com/fcfcqloow/first-todo-list/backend/usecase"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/fcfcqloow/go-advance/log"
)

func main() {
	ldb.Install()

	log.SetLevelOrDefault(
		os.Getenv("LOG_LEVEL"),
		log.LOG_LEVEL_INFO,
	)

	log.Info(log.GetLevel())

	todoRepo := ldb.NewTodoRepository(ldb.TodoDbFile)
	topicRepo := ldb.NewTopicRepository(ldb.TopicDbFile)
	settingsRepo := ldb.NewSettingsRepository(ldb.SettingDbFile)
	logRepo := ldb.NewLogRepository()

	todoUseCase := usecase.NewTodoUseCase(todoRepo)
	topicUseCase := usecase.NewTopicUseCase(topicRepo, todoRepo)
	settingsUseCase := usecase.NewSettingsUseCase(settingsRepo)
	logUseCase := usecase.NewLogUseCase(logRepo)

	ginapp.Run(
		todoUseCase,
		topicUseCase,
		settingsUseCase,
		logUseCase,
		ldb.StaticDir,
		cnv.SafeInt(os.Getenv("PORT"), 8080),
	)
}
