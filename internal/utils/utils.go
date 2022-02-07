package utils

import (
	"crypto/rand"
	"io"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateRandomExpiryCode(key string) string {

	max := 6
	value := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, value, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(value); i++ {
		value[i] = table[int(value[i])%len(table)]
	}

	return string(value)
}
