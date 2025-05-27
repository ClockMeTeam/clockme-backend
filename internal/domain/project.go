package domain

import "github.com/google/uuid"

type Project struct {
	ID        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	TypeID    string      `json:"type_id"`
	Type      ProjectType `json:"type"`
	Users     []User      `json:"users"`
	CreatedAt string      `json:"created_at"`
	UpdateAt  string      `json:"update_at"`
}
