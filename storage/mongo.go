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
func (m *MongoStorage) Upsert(
	ctx context.Context,
	collection string,
	q *chatauth.Query,
) error {
	coll, session := m.client.WithContext(ctx).C(collection)
	defer session.Close()

	_, err := coll.Upsert(q.Selector, q.Update)

	return err
}

// RemoveAll removes document from mongo
func (m *MongoStorage) RemoveAll(
	ctx context.Context,
	collection string,
	selector interface{},
) error {
	coll, session := m.client.WithContext(ctx).C(collection)
	defer session.Close()

	_, err := coll.RemoveAll(selector)
	return err
}

// BulkUpsert bulk upserts documents on mongo
func (m *MongoStorage) BulkUpsert(
	ctx context.Context,
	collection string,
	qs []*chatauth.Query,
) error {
	coll, session := m.client.WithContext(ctx).C(collection)
	defer session.Close()

	bulk := coll.Bulk()
	pairs := make([]interface{}, 0, 2*len(qs))
	for _, q := range qs {
		pairs = append(pairs, q.Selector, q.Update)
	}

	bulk.Upsert(pairs...)

	_, err := bulk.Run()
	return err
}
