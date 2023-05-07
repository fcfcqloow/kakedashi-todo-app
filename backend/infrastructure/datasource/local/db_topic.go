package ldb

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
	"github.com/fcfcqloow/first-todo-list/backend/usecase"
)

var topicCacheKey = "topic-cache-key"

type topicDB struct {
	localTopicPath string
}

func NewTopicRepository(localTopicPath string) usecase.TopicRepository {
	return &topicDB{
		localTopicPath: localTopicPath,
	}
}

func (t *topicDB) List() ([]domain.Topic, error) {
	topics, err := load[[]domain.Topic](t.localTopicPath, topicCacheKey)
	if err != nil {
		return nil, err
	}

	return *topics, nil
}

func (t *topicDB) Sync(topics []domain.Topic) error {
	if err := save(t.localTopicPath, topics, topicCacheKey); err != nil {
		return err
	}

	return nil
}
