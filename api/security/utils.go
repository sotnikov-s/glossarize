package security

import (
	"crypto/rand"
	"encoding/base64"
)

// randomStringLength represents the default length of random string.
const randomStringLength = 64

// generateRandomString generates a new random string.
func generateRandomString(length int) (string, error) {
	b := make([]byte, (length/4)*3)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
