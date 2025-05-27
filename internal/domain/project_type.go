package domain

import "github.com/google/uuid"

type ProjectType struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt string    `json:"created_at"`
	UpdateAt  string    `json:"update_at"`
}
