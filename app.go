package chatauth

import (
	"context"
	"time"

	"github.com/spf13/viper"
)

// ChatAuth communicates with storage to register
// players and authorize them in chat
// Implements Interface interface
type ChatAuth struct {
	storage           Storage
	password          Password
	userColl, aclColl string
	gameID            string
}

// NewChatAuth constructs chat auth
func NewChatAuth(
	storage Storage,
	passwordOrNil Password,
	configPrefix string,
	config *viper.Viper,
) *ChatAuth {
	var (
		userCollKey = configPrefix + ".userColl"
		aclCollKey  = configPrefix + ".aclColl"
		gameIDKey   = configPrefix + ".gameId"
	)

	config.SetDefault(userCollKey, "mqtt_user")
	config.SetDefault(aclCollKey, "mqtt_acl")

	if passwordOrNil == nil {
		passwordOrNil = &PasswordPBKDF2{}
	}

	return &ChatAuth{
		storage:  storage,
		password: passwordOrNil,
		userColl: config.GetString(userCollKey),
		aclColl:  config.GetString(aclCollKey),
		gameID:   config.GetString(gameIDKey),
	}
}

// RegisterPlayer registers player on storage
func (c *ChatAuth) RegisterPlayer(
	ctx context.Context,
	user string,
	password []byte,
) error {
	hashedPass, salt, err := c.password.Hash(password)
	if err != nil {
		return err
	}

	err = c.storage.Upsert(ctx, c.userColl, &Query{
		Selector: UserSelector{
			Username: getUser(c.gameID, user),
		},
		Update: UserAuthUpdater{
			Set: UserAuth{
				Username: getUser(c.gameID, user),
				Password: hashedPass,
				Salt:     salt,
			},
		},
	})

	return err
}

// Authorize authorizes player in room
func (c *ChatAuth) Authorize(ctx context.Context, user string, rooms []string) error {
	queries := make([]*Query, len(rooms))

	for idx, room := range rooms {
		queries[idx] = &Query{
			Selector: UserRoomSelector{
				Username: getUser(c.gameID, user),
				PubSub:   getRoom(c.gameID, room),
			},
			Update: UserTopicUpdater{
				Username: getUser(c.gameID, user),
				PubSub:   []string{getRoom(c.gameID, room)},
			},
		}
	}

	err := c.storage.BulkUpsert(ctx, c.aclColl, queries)

	return err
}

// AuthorizeWithExpire authorizes player in room and add an expires field
// In case of Mongo, this field must have a index with ttl with expireAfterSeconds=0
func (c *ChatAuth) AuthorizeWithExpire(
	ctx context.Context,
	user string,
	rooms []string,
	expiresAt time.Time,
) error {
	queries := make([]*Query, len(rooms))
	for idx, room := range rooms {
		queries[idx] = &Query{
			Selector: UserRoomSelector{
				Username: getUser(c.gameID, user),
				PubSub:   getRoom(c.gameID, room),
			},
			Update: UserTopicUpdater{
				Username: getUser(c.gameID, user),
				PubSub:   []string{getRoom(c.gameID, room)},
				Expires:  expiresAt,
			},
		}
	}

	err := c.storage.BulkUpsert(ctx, c.aclColl, queries)

	return err
}

// Unauthorize unauthorizes player in room
func (c *ChatAuth) Unauthorize(
	ctx context.Context,
	user string,
	rooms []string,
) error {
	topicRooms := make([]string, len(rooms))
	for idx, room := range rooms {
		topicRooms[idx] = getRoom(c.gameID, room)
	}

	err := c.storage.RemoveAll(ctx, c.aclColl, UserManyRoomsSelector{
		Username: getUser(c.gameID, user),
		PubSub: PubSubsSelector{
			In: topicRooms,
		},
	})

	return err
}
