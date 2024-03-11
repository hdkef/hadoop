package helper

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateRandomString() (string, error) {
	randomBytes := make([]byte, 18)

	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return "", err
	}

	return hex.EncodeToString(randomBytes), nil
}
