package chatauth

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltLength = 12
	keyLength  = 20
	iterations = 1000
)

func hash(password []byte) (hashedPass, salt string, err error) {
	saltBts := make([]byte, saltLength)
	_, err = rand.Read(saltBts)
	if err != nil {
		return "", "", err
	}

	hashedPassBts := pbkdf2.Key(password, saltBts, iterations, keyLength, sha256.New)

	return fmt.Sprintf("%x", hashedPassBts), fmt.Sprintf("%x", saltBts), nil
}
