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

// PasswordPBKDF2 generates a password hash using pbkdf2
type PasswordPBKDF2 struct{}

// Hash generates a hashed password and salt from password
func (p *PasswordPBKDF2) Hash(password []byte) (string, string, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", err
	}

	hashedPass := pbkdf2.Key(password, salt, iterations, keyLength, sha256.New)

	return fmt.Sprintf("%x", hashedPass), fmt.Sprintf("%x", salt), nil
}
