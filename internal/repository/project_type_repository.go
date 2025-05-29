package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maevlava/ftf-clockify/internal/domain"
	"github.com/maevlava/ftf-clockify/internal/repository/postgres/db"
)

type PgProjectTypeRepository struct {
	q *db.Queries
}

func NewPgProjectTypeRepository(pool *pgxpool.Pool) domain.ProjectTypeRepository {
	return &PgProjectTypeRepository{
		q: db.New(pool),
	}
}

func (p PgProjectTypeRepository) GetProjectType(ctx context.Context, name string) (domain.ProjectType, error) {
	// TODO implement me
	panic("implement me")
}

func (p PgProjectTypeRepository) GetProjectTypes(ctx context.Context) ([]domain.ProjectType, error) {
	dbProjectTypes, err := p.q.GetProjectTypes(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error getting project types %w", err)
	}
	var projectTypes []domain.ProjectType
	for _, projectType := range dbProjectTypes {
		projectTypes = append(projectTypes, domain.ProjectType{
			Name: projectType.Name,
		})
	}
	return projectTypes, nil
}

func (p PgProjectTypeRepository) UpdateProjectType(ctx context.Context, name string) (domain.ProjectType, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgProjectTypeRepository) CreateProjectType(ctx context.Context, name string) (domain.ProjectType, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgProjectTypeRepository) DeleteProjectTypeByName(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}
