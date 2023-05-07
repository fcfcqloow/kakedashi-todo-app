package usecase

import (
	"fmt"

	"github.com/fcfcqloow/first-todo-list/backend/domain"
	domainerr "github.com/fcfcqloow/first-todo-list/backend/domain/errors"
	"github.com/fcfcqloow/go-advance/log"
)

type (
	TopicUseCase interface {
		ListTopics() ([]domain.Topic, error)
		AddTopic(domain.Topic) ([]domain.Topic, error)
		RemoveTopic(domain.Topic) ([]domain.Topic, error)
		UpdateTopic(domain.Topic) ([]domain.Topic, error)
	}
	topicUseCase struct {
		todoRepository  TodoRepository
		topicRepository TopicRepository
	}
)

func NewTopicUseCase(topicRepository TopicRepository, todoRepository TodoRepository) TopicUseCase {
	return &topicUseCase{
		todoRepository:  todoRepository,
		topicRepository: topicRepository,
	}
}

func (t *topicUseCase) ListTopics() ([]domain.Topic, error) {
	result, err := t.topicRepository.List()
	if err != nil {
		return nil, fmt.Errorf("[topic usecase] failed to list topics: %w", err)
	}

	return result, err
}

func (t *topicUseCase) GetTopic(id string) (*domain.Topic, error) {
	topics, err := t.topicRepository.List()
	if err != nil {
		return nil, fmt.Errorf("[topic usecase] failed to get topic: %w", err)
	}

	for _, v := range topics {
		if v.Id == id {
			return &v, nil
		}
	}

	return nil, &domainerr.NotFoundError{}
}

func (t *topicUseCase) AddTopic(topic domain.Topic) ([]domain.Topic, error) {
	topics, err := t.ListTopics()
	if err != nil {
		return nil, fmt.Errorf("[topic usecase] failed to add topics: %w", err)
	}

	topics = append(topics, topic)

	if err := t.topicRepository.Sync(topics); err != nil {
		return nil, fmt.Errorf("[topic usecase] failed to add topics: %w", err)
	}

	return topics, nil

}

func (t *topicUseCase) RemoveTopic(topic domain.Topic) ([]domain.Topic, error) {
	topics, err := t.ListTopics()
	if err != nil {
		return nil, fmt.Errorf("[topic usecase] failed to remove topics: %w", err)
	}

	if tasks, err := t.todoRepository.List(); err != nil {
		log.Warn(err)
	} else if err := tasks.ClearTopic(topic.Id); err != nil {
		log.Warn(err)
	} else if err := t.todoRepository.Sync(*tasks); err != nil {
		log.Warn(err)
	}

	for i, v := range topics {
		if v.Id == topic.Id {
			topics = append(topics[:i], topics[i+1:]...)
			if err := t.topicRepository.Sync(topics); err != nil {
				return nil, fmt.Errorf("[topic usecase] failed to remove topics: %w", err)
			}

			return topics, nil
		}
	}

	return nil, &domainerr.NotFoundError{}
}

func (t *topicUseCase) UpdateTopic(topic domain.Topic) ([]domain.Topic, error) {
	topics, err := t.ListTopics()
	if err != nil {
		return nil, fmt.Errorf("[topic usecase] failed to update topics: %w", err)
	}

	for i, v := range topics {
		if v.Id == topic.Id {
			topics[i] = topic
			if err := t.topicRepository.Sync(topics); err != nil {
				return nil, fmt.Errorf("[topic usecase] failed to update topics: %w", err)
			}

			return topics, nil
		}
	}

	return nil, &domainerr.NotFoundError{}

}
