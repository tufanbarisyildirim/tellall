package utils

import (
	"crypto/rand"
	"fmt"
)

func RandHexId(byteSize int32) string {
	b := make([]byte, byteSize)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
