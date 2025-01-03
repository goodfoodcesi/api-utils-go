package message

import (
	"encoding/json"
	"fmt"
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

func NewMessage(msgType string, producer string, payload any) *Message {
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

func (m *Message) ExtractPayloadEntity(target any) error {
	payloadBytes, err := json.Marshal(m.Payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	err = json.Unmarshal(payloadBytes, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal payload into target: %w", err)
	}
	return nil
}
