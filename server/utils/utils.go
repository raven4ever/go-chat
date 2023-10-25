package utils

import (
	"fmt"
	"strings"
)

// IsValidPortNumber checks if a given integer value is a valid port number.
// A valid port number is between 1024 and 65535 (inclusive).
func IsValidPortNumber(value int) bool {
	return value > 1023 && value < 65536
}

// message struct
type Message struct {
	Username string
	Content  string
}

// returns a string representation of the Message struct.
func (m *Message) String() string {
	return fmt.Sprintf("%s: %s\n", m.Username, strings.TrimSpace(m.Content))
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
