package chatauth

// Password has methods to generage hash and salt from passwords
type Password interface {
	Hash(password []byte) (hashedPass, salt string, err error)
}
