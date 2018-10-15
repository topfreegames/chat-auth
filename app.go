package chatauth

import "context"

var (
	storage  Storage
	userColl = "mqtt_user"
	aclColl  = "mqtt_acl"
)

// SetStorage set storage to use to save users and auths
func SetStorage(str Storage) {
	storage = str
}

// SetUserCollection updates user collection name
func SetUserCollection(coll string) {
	userColl = coll
}

// SetACLCollection updates ACL collection name
func SetACLCollection(coll string) {
	aclColl = coll
}

// RegisterUser registers user on storage
func RegisterUser(ctx context.Context, user string, password []byte) error {
	hashedPass, salt, err := hash(password)
	if err != nil {
		return err
	}

	err = storage.Upsert(ctx, userColl, &Query{
		Selector: map[string]string{
			"username": user,
		},
		Update: map[string]interface{}{
			"$set": userAuth{
				Username: user,
				Password: hashedPass,
				Salt:     salt,
			},
		},
	})

	return err
}
