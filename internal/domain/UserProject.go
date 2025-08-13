package domain

import "github.com/google/uuid"

type UserProject struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	ProjectID string    `json:"project_id"`
	User      *User     `json:"user"`
	Project   *Project  `json:"project"`
	CreatedAt string    `json:"created_at"`
	UpdateAt  string    `json:"update_at"`
}
