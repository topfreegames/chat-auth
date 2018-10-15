package chatauth

import "context"

// Storage has methods to access db storage
// or any other system to store users
type Storage interface {
	Upsert(context.Context, *Query) error
	BulkUpsert(context.Context, []*Query) error
	Remove(ctx context.Context, selector interface{}) error
}

// Query contains a selector object and a
// update query
type Query struct {
	Selector interface{}
	Update   interface{}
}
