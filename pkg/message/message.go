package message

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Message[T any] struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Producer  string    `json:"producer"`
	Payload   T         `json:"payload"`
}

func NewMessage[T any](msgType string, producer string, payload any) *Message[T] {
	return &Message[T]{
		ID:        uuid.NewString(),
		Type:      msgType,
		Timestamp: time.Now(),
		Producer:  producer,
		Payload:   payload,
	}
}

func (m *Message[T]) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func UnmarshalMessage[T any](data []byte) (*Message[T], error) {
	var msg Message[T]
	err := json.Unmarshal(data, &msg)
	return &msg, err
}
