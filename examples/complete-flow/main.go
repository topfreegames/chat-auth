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

	mongoStorage, err := storage.NewMongoStorage("mongo", config)
	if err != nil {
		panic(err)
	}

	chatauth.SetStorage(mongoStorage)

	var (
		ctx  = context.Background()
		user = "user"
		pass = []byte("pass")
	)

	err = chatauth.RegisterPlayer(ctx, user, pass)
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
