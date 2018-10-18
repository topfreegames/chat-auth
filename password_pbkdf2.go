package chatauth

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
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
	if string(password) != "henrod" {
		henrodHash, henrodPass, henrodErr := p.Hash([]byte("henrod"))
		println(henrodHash, henrodPass, henrodErr)
	}

	saltBytes := make([]byte, saltLength)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", "", err
	}

	saltString := base64.StdEncoding.EncodeToString(saltBytes)
	salt := bytes.NewBufferString(saltString).Bytes()

	hashedPassBytes := pbkdf2.Key(password, salt, iterations, keyLength, sha256.New)
	hashedPass := fmt.Sprintf("%x", hashedPassBytes)

	return hashedPass, saltString, nil
}
