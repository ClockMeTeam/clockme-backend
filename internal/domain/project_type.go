package domain

import (
	"context"
	"github.com/google/uuid"
)

type ProjectType struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	BaseHour  int32     `json:"base_hour" db:"base_hour"`
	CreatedAt string    `json:"created_at"`
	UpdateAt  string    `json:"update_at"`
}

type ProjectTypeRepository interface {
	GetProjectType(ctx context.Context, name string) (ProjectType, error)
	GetProjectTypes(ctx context.Context) ([]ProjectType, error)
	UpdateProjectType(ctx context.Context, name string) (ProjectType, error)
	CreateProjectType(ctx context.Context, name string) (ProjectType, error)
	DeleteProjectTypeByName(ctx context.Context, name string) error
}
