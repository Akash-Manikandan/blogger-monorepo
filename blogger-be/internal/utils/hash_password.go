package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hashed := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	hashWithSalt := append(salt, hashed...)

	return base64.StdEncoding.EncodeToString(hashWithSalt), nil
}
