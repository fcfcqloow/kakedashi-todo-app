package ginapp

import (
	"strconv"

	"github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/ginapp/controller"
	"github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/ginapp/middleware"
	"github.com/fcfcqloow/first-todo-list/backend/usecase"

	"github.com/fcfcqloow/go-advance/log"
	"github.com/gin-gonic/gin"
)

func Run(
	todoUseCase usecase.TodoUseCase,
	topicUseCase usecase.TopicUseCase,
	settingsUseCase usecase.SettingsUseCase,
	logUseCase usecase.LogUseCase,
	htmlIndexDir string,
	port int,
) {

	todoController := controller.NewController(
		todoUseCase,
		topicUseCase,
		settingsUseCase,
		logUseCase,
	)

	{
		gin.SetMode(gin.ReleaseMode)

		engine := gin.Default()
		engine.
			Use(middleware.Recovery()).
			Use(middleware.ErrorHandling()).
			Use(middleware.RequestLog()).
			Use(middleware.CORS()).
			// TODO Controller
			GET("/todo", todoController.ListTodo).
			PATCH("/todo", todoController.UpdateTodo).
			PATCH("/todo/sort", todoController.SortTodo).
			DELETE("/todo", todoController.RemoveTodo).
			PUT("/todo", todoController.AddTodo).
			POST("/todo", todoController.MoveTodo).
			// Topic Controller
			GET("/topic", todoController.ListTopics).
			PUT("/topic", todoController.AddTopic).
			DELETE("/topic", todoController.RemoveTopic).
			PATCH("/topic", todoController.UpdateTopic).
			// Settings Controller
			GET("/settings", todoController.GetSettings).
			PATCH("/settings", todoController.UpdateSettings).
			GET("/summary", todoController.GetLogs).
			GET("/logs", todoController.GetDates).
			GET("/logs/:date", todoController.GetLogs).
			// Websocket Controller
			GET("/ws", controller.WebsocketHandle(map[string]gin.HandlerFunc{
				"log": todoController.Log,
			})).
			// Static html
			Static("/index", htmlIndexDir)

		log.Info("listen http://localhost:" + strconv.Itoa(port))
		if err := engine.Run(":" + strconv.Itoa(port)); err != nil {
			log.Fatal(err)
		}
	}
}
