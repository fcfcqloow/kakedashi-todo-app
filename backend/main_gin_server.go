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

	localRepository := ldb.NewLocalRepository()

	todoUseCase := usecase.NewTodoUseCase(localRepository, localRepository)
	topicUseCase := usecase.NewTopicUseCase(localRepository, localRepository)
	settingsUseCase := usecase.NewSettingsUseCase(localRepository)
	logUseCase := usecase.NewLogUseCase(localRepository)
	memoUseCase := usecase.NewMemoUseCase(localRepository)

	ginapp.Run(
		todoUseCase,
		topicUseCase,
		settingsUseCase,
		logUseCase,
		memoUseCase,
		ldb.StaticDir,
		cnv.SafeInt(os.Getenv("PORT"), 8080),
	)
}
