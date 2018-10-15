package storage

import (
	"context"

	"github.com/spf13/viper"
	"github.com/topfreegames/extensions/mongo"
	"github.com/topfreegames/extensions/mongo/interfaces"

	chatauth "github.com/topfreegames/chat-auth"
)

// MongoStorage implements Storage interface
type MongoStorage struct {
	client interfaces.MongoDB
}

// NewMongoStorage connects to mongo and returns client
func NewMongoStorage(
	configPrefix string,
	config *viper.Viper,
) (*MongoStorage, error) {
	client, err := mongo.NewClient(configPrefix, config)
	if err != nil {
		return nil, err
	}

	return &MongoStorage{
		client: client.MongoDB,
	}, nil
}

// Upsert upserts on mongo
func (m *MongoStorage) Upsert(ctx context.Context, q *chatauth.Query) error {
	return nil
}

// BulkUpsert bulk upserts on mongo
func (m *MongoStorage) BulkUpsert(ctx context.Context, qs []*chatauth.Query) error {
	return nil
}

// Remove removes document from mongo
func (m *MongoStorage) Remove(ctx context.Context, selector interface{}) error {
	return nil
}
