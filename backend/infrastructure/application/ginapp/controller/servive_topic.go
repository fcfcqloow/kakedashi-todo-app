package controller

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
	"github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/ginapp/model"
)

func toFeTopicResponse(topics []domain.Topic) []model.Topic {
	result := []model.Topic{}
	for _, topic := range topics {
		result = append(result, model.Topic{
			Id:          topic.Id,
			Value:       topic.Value,
			Color:       topic.Color,
			CreatedDate: topic.CreatedDate,
			DoneDate:    topic.DoneDate,
		})
	}

	return result
}
