package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maevlava/ftf-clockify/internal/domain"
	"github.com/maevlava/ftf-clockify/internal/repository/postgres/db"
)

type PgProjectRepository struct {
	q *db.Queries
}

func NewPgProjectRepository(pool *pgxpool.Pool) domain.ProjectRepository {
	return &PgProjectRepository{
		q: db.New(pool),
	}
}
func (p PgProjectRepository) GetProjects(ctx context.Context) ([]domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgProjectRepository) UpdateProject(ctx context.Context, project domain.Project) (domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgProjectRepository) GetProjectByName(ctx context.Context, name string) (domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgProjectRepository) CreateProject(ctx context.Context, project domain.Project) (domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgProjectRepository) DeleteProjectByName(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}

func (p PgProjectRepository) GetProjectType(ctx context.Context, clockifyId string) (domain.ProjectType, error) {

	panic("implement me")
}

func (p PgProjectRepository) GetProjectTypeByClockifyID(ctx context.Context, clockifyID string) (domain.ProjectType, error) {
	dbProjectType, err := p.q.GetProjectTypeByClockifyId(ctx, clockifyID)
	if err != nil {
		return domain.ProjectType{}, fmt.Errorf("could not get project type: %w", err)
	}

	return domain.ProjectType{
		ID:   *dbProjectType.ID,
		Name: dbProjectType.Name,
	}, nil
}

func (p PgProjectRepository) UpdateProjectType(ctx context.Context, typeID string) (domain.ProjectType, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgProjectRepository) GetProjectUsers(ctx context.Context, name string) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}
