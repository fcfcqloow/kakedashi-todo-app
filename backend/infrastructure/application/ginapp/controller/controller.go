package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
	"github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/ginapp/model"
	"github.com/fcfcqloow/first-todo-list/backend/usecase"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/gin-gonic/gin"
)

type (
	controller struct {
		taskUseCase     usecase.TodoUseCase
		topicUseCase    usecase.TopicUseCase
		settingsUseCase usecase.SettingsUseCase
		logUseCase      usecase.LogUseCase
	}
)

func NewController(
	taskUseCase usecase.TodoUseCase,
	topicUseCase usecase.TopicUseCase,
	settingsUseCase usecase.SettingsUseCase,
	logUseCase usecase.LogUseCase,
) *controller {
	return &controller{
		taskUseCase:     taskUseCase,
		topicUseCase:    topicUseCase,
		settingsUseCase: settingsUseCase,
		logUseCase:      logUseCase,
	}
}

func (c *controller) ListTodo(ctx *gin.Context) {
	var (
		target = ctx.Query("target")
		offset = cnv.MustInt(ctx.Query("offset"))
		limit  = cnv.MustInt(ctx.Query("limit"))
	)

	tasks, err := c.taskUseCase.ListTodo(
		usecase.TodoLimit(limit),
		usecase.TodoOffset(offset),
		usecase.TodoTarget(target),
	)

	if err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, feResponse(*tasks))
}

func (c *controller) SortTodo(ctx *gin.Context) {
	if _, err := c.taskUseCase.Sort(); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, model.GeneralMapResponse{"message": "OK"})
}

func (c *controller) RemoveTodo(ctx *gin.Context) {
	request := model.TasksRemoveRequest{}
	if err := ctx.Bind(&request); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	if _, err := c.taskUseCase.DeleteTodo(domainTask(model.Task(request))); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, model.GeneralMapResponse{"message": "OK"})
}

func (c *controller) AddTodo(ctx *gin.Context) {
	var request model.TasksAddRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	if _, err := c.taskUseCase.AddTodo(domainTask(model.Task(request))); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, model.GeneralMapResponse{"message": "OK"})
}

func (c *controller) UpdateTodo(ctx *gin.Context) {
	request := model.TasksUpdateRequest{}
	if err := ctx.Bind(&request); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	if _, err := c.taskUseCase.UpdateTodo(domainTask(model.Task(request))); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, model.GeneralMapResponse{"message": "OK"})
}

func (c *controller) MoveTodo(ctx *gin.Context) {
	request := model.TasksMoveRequest{}
	if err := ctx.Bind(&request); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	if _, err := c.taskUseCase.MoveTodo(
		request.From,
		request.To,
		request.Idx,
		domainTask(model.Task(request.Task)),
	); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, model.GeneralMapResponse{"message": "OK"})
}

func (c *controller) ListTopics(ctx *gin.Context) {
	topics, err := c.topicUseCase.ListTopics()
	if err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, toFeTopicResponse(topics))
}

func (c *controller) RemoveTopic(ctx *gin.Context) {
	request := model.TopicRemoveRequest{}
	if err := ctx.Bind(&request); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	if _, err := c.topicUseCase.RemoveTopic(
		domain.Topic{
			Id:          request.Id,
			Value:       request.Value,
			Color:       request.Color,
			CreatedDate: request.CreatedDate,
		},
	); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, model.GeneralMapResponse{"message": "OK"})
}

func (c *controller) UpdateTopic(ctx *gin.Context) {
	request := model.TopicUpdateRequest{}
	if err := ctx.Bind(&request); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	if _, err := c.topicUseCase.UpdateTopic(
		domain.Topic{
			Id:          request.Id,
			Value:       request.Value,
			Color:       request.Color,
			CreatedDate: request.CreatedDate,
		},
	); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, model.GeneralMapResponse{"message": "OK"})
}

func (c *controller) AddTopic(ctx *gin.Context) {
	request := model.TopicAddRequest{}
	if err := ctx.Bind(&request); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	if _, err := c.topicUseCase.AddTopic(
		domain.Topic{
			Id:          request.Id,
			Value:       request.Value,
			Color:       request.Color,
			CreatedDate: request.CreatedDate,
			DoneDate:    request.DoneDate,
		},
	); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, model.GeneralMapResponse{"message": "OK"})
}

func (c *controller) GetSettings(ctx *gin.Context) {
	settings, err := c.settingsUseCase.GetSettings()
	if err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, feSettingsResponse(settings))
}

func (c *controller) UpdateSettings(ctx *gin.Context) {
	request := model.SettingsUpdateRequest{}
	if err := ctx.Bind(&request); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	if _, err := c.settingsUseCase.UpdateSettings(domainSettings(request)); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, model.GeneralMapResponse{"message": "OK"})
}

func (c *controller) Log(ctx *gin.Context) {
	var request model.LogRequest
	byts, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	if err := json.Unmarshal(byts, &request); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	defer ctx.Request.Body.Close()

	if _, err := c.logUseCase.Append(cnv.MustCompactJson(request.Value)); err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}
}

func (c *controller) GetLogs(ctx *gin.Context) {
	date := ctx.Param("date")
	logs, err := c.logUseCase.List(usecase.Date(date))
	if err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, feLogs(logs))
}

func (c *controller) GetDates(ctx *gin.Context) {
	dates, err := c.logUseCase.ListDate()
	if err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, feDateList(dates))
}

func (c *controller) GetSummary(ctx *gin.Context) {
	var (
		start = cnv.MustInt(ctx.Query("start"))
		end   = cnv.MustInt(ctx.Query("end"))
	)

	summary, err := c.logUseCase.CreateSummary(start, end)
	if err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, feSummary(*summary))
}
