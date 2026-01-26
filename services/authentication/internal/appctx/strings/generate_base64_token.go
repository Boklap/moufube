package strings

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateBase64Token(byteSize int) (string, error) {
	b := make([]byte, byteSize)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
