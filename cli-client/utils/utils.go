package utils

import (
	"math/rand"
	"time"
)

const DEFAULT_CHARSET = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset returns a random string of the specified length, using the given charset.
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomString generates a random string of the specified length using the default charset.
func RandomString(length int) string {
	return StringWithCharset(length, DEFAULT_CHARSET)
}
