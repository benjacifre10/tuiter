package utils

import "golang.org/x/crypto/bcrypt"

/* Encrypt the uncoming password for save later */
func EncryptPassword(pass string) (string, error) {
	salt := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), salt)
	return string(bytes), err
}
