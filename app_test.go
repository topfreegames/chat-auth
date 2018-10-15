// +build unit

package chatauth_test

import (
	"context"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/topfreegames/chat-auth/mocks"

	gomock "github.com/golang/mock/gomock"
	chatauth "github.com/topfreegames/chat-auth"
)

func TestRegisterPlayer(t *testing.T) {
	t.Parallel()

	var (
		ctrl         = gomock.NewController(t)
		mockStorage  = mocks.NewMockStorage(ctrl)
		mockPassword = mocks.NewMockPassword(ctrl)
		ctx          = context.Background()
		user         = "user"
		password     = []byte("password")
		hash         = "hash"
		salt         = "salt"
		coll         = "mqtt_user"
		game         = "game"
		prefix       = "chat"
	)

	mockPassword.EXPECT().Hash(password).Return(hash, salt, nil)

	mockStorage.EXPECT().Upsert(ctx, coll, &chatauth.Query{
		Selector: chatauth.UserSelector{
			Username: "game:user",
		},
		Update: chatauth.UserAuthUpdater{
			Set: chatauth.UserAuth{
				Username: "game:user",
				Password: hash,
				Salt:     salt,
			},
		},
	})

	config := viper.New()
	config.SetDefault("chat.gameId", game)

	auth := chatauth.NewChatAuth(mockStorage, mockPassword, prefix, config)
	err := auth.RegisterPlayer(ctx, user, password)
	assert.NoError(t, err)
}

func TestAuthorize(t *testing.T) {
	t.Parallel()

	var (
		mockStorage = mocks.NewMockStorage(gomock.NewController(t))
		ctx         = context.Background()
		user        = "user"
		room        = "room"
		coll        = "mqtt_acl"
		game        = "game"
		prefix      = "chat"
	)

	mockStorage.EXPECT().Upsert(ctx, coll, &chatauth.Query{
		Selector: chatauth.UserRoomSelector{
			Username: "game:user",
			PubSub:   "chat/game/room/room",
		},
		Update: chatauth.UserTopicUpdater{
			Username: "game:user",
			PubSub:   []string{"chat/game/room/room"},
		},
	})

	config := viper.New()
	config.SetDefault("chat.gameId", game)

	auth := chatauth.NewChatAuth(mockStorage, nil, prefix, config)
	err := auth.Authorize(ctx, user, room)
	assert.NoError(t, err)
}

func TestUnauthorize(t *testing.T) {
	t.Parallel()

	var (
		mockStorage = mocks.NewMockStorage(gomock.NewController(t))
		ctx         = context.Background()
		user        = "user"
		room        = "room"
		coll        = "mqtt_acl"
		game        = "game"
		prefix      = "chat"
	)

	mockStorage.EXPECT().Remove(ctx, coll, chatauth.UserRoomSelector{
		Username: "game:user",
		PubSub:   "chat/game/room/room",
	})

	config := viper.New()
	config.SetDefault("chat.gameId", game)

	auth := chatauth.NewChatAuth(mockStorage, nil, prefix, config)
	err := auth.Unauthorize(ctx, user, room)
	assert.NoError(t, err)
}
