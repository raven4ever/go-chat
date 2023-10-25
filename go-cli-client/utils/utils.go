package utils

import (
	"fmt"
	"math/rand"
	"strings"
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

// message struct
type Message struct {
	Username string
	Content  string
}

// returns a string representation of the Message struct.
func (m *Message) String() string {
	return fmt.Sprintf("%s: %s\n", m.Username, m.Content)
}

// returns a Message struct from a string
func NewMessage(s string) *Message {
	content := strings.Split(s, ":")
	if len(content) == 2 {
		return &Message{Username: content[0], Content: content[1]}
	}
	return nil
}

// returns the []byte representation of the Message struct
func (m *Message) Bytes() []byte {
	return []byte(m.String())
}
