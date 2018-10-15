package chatauth

import "fmt"

func getUser(gameID, user string) string {
	return fmt.Sprintf("%s:%s", gameID, user)
}

func getRoom(gameID, room string) string {
	return fmt.Sprintf("chat/%s/room/%s", gameID, room)
}

func getUserRoom(gameID, user, room string) string {
	return fmt.Sprintf("%s-%s", getUser(gameID, user), getRoom(gameID, room))
}
