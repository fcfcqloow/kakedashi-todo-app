package ldb

import (
	"github.com/fcfcqloow/first-todo-list/backend/domain"
)

var topicCacheKey = "topic-cache-key"

func (t *localRepository) ListTopics() ([]domain.Topic, error) {
	topics, err := load[[]domain.Topic](TopicDbFile, topicCacheKey)
	if err != nil {
		return nil, err
	}

	return *topics, nil
}

func (t *localRepository) SyncTopics(topics []domain.Topic) error {
	if err := save(TopicDbFile, topics, topicCacheKey); err != nil {
		return err
	}

	return nil
}
