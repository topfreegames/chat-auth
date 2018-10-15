package chatauth

type (
	userAuthUpdater struct {
		Set userAuth `bson:"$set"`
	}

	userAuth struct {
		Username string `bson:"username"`
		Password string `bson:"password"`
		Salt     string `bson:"salt"`
	}

	userTopicUpdater struct {
		Username string   `bson:"username"`
		PubSub   []string `bson:"pubsub"`
	}
)

type (
	userSelector struct {
		Username string `bson:"username"`
	}

	userRoomSelector struct {
		Username string `bson:"username"`
		PubSub   string `bson:"pubsub"`
	}
)
