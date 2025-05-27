package domain

import (
	"context"
	"github.com/google/uuid"
)

type Project struct {
	ID        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	TypeID    string      `json:"type_id"`
	Type      ProjectType `json:"type"`
	Users     []User      `json:"users"`
	CreatedAt string      `json:"created_at"`
	UpdateAt  string      `json:"update_at"`
}

type ProjectRepository interface {
	GetProjects(ctx context.Context) ([]Project, error)
	UpdateProject(ctx context.Context, project Project) (Project, error)
	GetProjectByName(ctx context.Context, name string) (Project, error)
	CreateProject(ctx context.Context, project Project) (Project, error)
	DeleteProjectByName(ctx context.Context, name string) error

	GetProjectType(ctx context.Context, name string) (ProjectType, error)
	UpdateProjectType(ctx context.Context, typeID string) (ProjectType, error)

	GetProjectUsers(ctx context.Context, name string) ([]User, error)
}
