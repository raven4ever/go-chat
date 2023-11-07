package utils

import (
	"fmt"
	"strings"
)

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
