package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func GeneratePassword(passwordSize int) (string, error) {
	// validate password size requested
	if passwordSize <= 5 {
		return "", errors.New("The password size needs to be bigger than 5")
	}
	if passwordSize >= 100 {
		return "", errors.New("The password size needs to be smaller than 100")
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

	// Generate the password
	password := make([]byte, passwordSize)
	for i := range password {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password), nil
}
