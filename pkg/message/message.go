package message

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Producer  string    `json:"producer"`
	Payload   any       `json:"payload"`
}

func NewMessage[T any](msgType string, producer string, payload any) *Message {
	return &Message{
		ID:        uuid.NewString(),
		Type:      msgType,
		Timestamp: time.Now(),
		Producer:  producer,
		Payload:   payload,
	}
}

func (m *Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func UnmarshalMessage(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	return &msg, err
}
