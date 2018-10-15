package main

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	chatauth "github.com/topfreegames/chat-auth"
	"github.com/topfreegames/chat-auth/storage"
)

func main() {
	config, err := getConfig()
	if err != nil {
		panic(err)
	}

	configPrefix := "mongo"
	mongoStorage, err := storage.NewMongoStorage(configPrefix, config)
	if err != nil {
		panic(err)
	}

	auth := chatauth.NewChatAuth(mongoStorage, configPrefix, config)

	var (
		ctx  = context.Background()
		user = "user"
		pass = []byte("pass")
		room = "room"
	)

	err = auth.RegisterPlayer(ctx, user, pass)
	if err != nil {
		panic(err)
	}

	err = auth.Authorize(ctx, user, room)
	if err != nil {
		panic(err)
	}

	err = auth.Unauthorize(ctx, user, room)
	if err != nil {
		panic(err)
	}

	fmt.Println("success")
}

func getConfig() (*viper.Viper, error) {
	config := viper.New()
	config.AddConfigPath(".")
	config.SetConfigName("config")
	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}
