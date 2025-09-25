package crypto

import (
	"crypto/rand"
	"encoding/base64"
)

func GenSession() (string, error) {
	const lenSes = 32
	sesBytes := make([]byte, lenSes)

	_, err := rand.Read(sesBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sesBytes), nil
}
