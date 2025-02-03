package entities

import (
	"database/sql"
)

type Note struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UserID      int          `json:"user_id"`
}
