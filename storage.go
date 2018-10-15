package chatauth

import "context"

// Storage has methods to access db storage
// or any other system to store users
type Storage interface {
	Upsert(ctx context.Context, collection string, q *Query) error
	Remove(ctx context.Context, collection string, selector interface{}) error
}

// Query contains a selector object and a
// update query
type Query struct {
	Selector interface{}
	Update   interface{}
}
