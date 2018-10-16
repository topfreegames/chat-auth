package chatauth

import (
	"context"
	"time"
)

// Interface contains the methods of this lib
type Interface interface {
	RegisterPlayer(ctx context.Context, user string, password []byte) error
	Authorize(ctx context.Context, user string, rooms []string) error
	AuthorizeWithExpire(ctx context.Context, user string, room []string, expiresAt time.Time) error
	Unauthorize(ctx context.Context, user string, room []string) error
}
