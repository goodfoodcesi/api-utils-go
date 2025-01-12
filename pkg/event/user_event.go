package event

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type UserCreatedEvent struct {
	ID        pgtype.UUID `json:"id"`
	Email     string      `json:"email"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Role      string      `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
}

type UserUpdatedEvent struct {
	ID        pgtype.UUID `json:"id"`
	Email     string      `json:"email"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Role      string      `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
}

type UserDeletedEvent struct {
	ID pgtype.UUID `json:"id"`
}
