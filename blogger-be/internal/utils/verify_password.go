package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func VerifyPassword(storedHash, password string) (bool, error) {
	decoded, err := base64.StdEncoding.DecodeString(storedHash)
	if err != nil {
		return false, err
	}

	salt := decoded[:16]
	hashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	return string(hashedPassword) == string(decoded[16:]), nil
}
