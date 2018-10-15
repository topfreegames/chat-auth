package chatauth

import (
	"context"

	"github.com/spf13/viper"
)

// ChatAuth communicates with storage to register
// players and authorize them in chat
type ChatAuth struct {
	storage           Storage
	userColl, aclColl string
	gameID            string
}

// NewChatAuth constructs chat auth
func NewChatAuth(
	storage Storage,
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

	return &ChatAuth{
		storage:  storage,
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
	hashedPass, salt, err := hash(password)
	if err != nil {
		return err
	}

	err = c.storage.Upsert(ctx, c.userColl, &Query{
		Selector: userSelector{
			Username: getUser(c.gameID, user),
		},
		Update: userAuthUpdater{
			Set: userAuth{
				Username: getUser(c.gameID, user),
				Password: hashedPass,
				Salt:     salt,
			},
		},
	})

	return err
}

// Authorize authorizes player in room
func (c *ChatAuth) Authorize(ctx context.Context, user, room string) error {
	err := c.storage.Upsert(ctx, c.aclColl, &Query{
		Selector: userRoomSelector{
			Username: getUser(c.gameID, user),
			PubSub:   getRoom(c.gameID, room),
		},
		Update: userTopicUpdater{
			Username: getUser(c.gameID, user),
			PubSub:   []string{getRoom(c.gameID, room)},
		},
	})

	return err
}

// Unauthorize unauthorizes player in room
func (c *ChatAuth) Unauthorize(ctx context.Context, user, room string) error {
	err := c.storage.Remove(ctx, c.aclColl, userRoomSelector{
		Username: getUser(c.gameID, user),
		PubSub:   getRoom(c.gameID, room),
	})

	return err
}
