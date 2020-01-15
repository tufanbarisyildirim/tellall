package utils

import (
	"crypto/rand"
	"fmt"
)

func RandHexId(byteSize int32) (string, error) {
	b := make([]byte, byteSize)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
