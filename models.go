package chatauth

type (
	// UserAuthUpdater is the user updater query
	UserAuthUpdater struct {
		Set UserAuth `bson:"$set"`
	}

	// UserAuth holds info to save on storage
	UserAuth struct {
		Username string `bson:"username"`
		Password string `bson:"password"`
		Salt     string `bson:"salt"`
	}

	// UserTopicUpdater is the user+room updater query
	UserTopicUpdater struct {
		Username string   `bson:"username"`
		PubSub   []string `bson:"pubsub"`
	}
)

type (
	// UserSelector is the select query by username
	UserSelector struct {
		Username string `bson:"username"`
	}

	// UserRoomSelector is the select query by username and room
	UserRoomSelector struct {
		Username string `bson:"username"`
		PubSub   string `bson:"pubsub"`
	}
)
