package repository

import (
	"context"
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
	//TODO implement me
	panic("implement me")
}

func (p PgProjectTypeRepository) GetProjectTypes(ctx context.Context) ([]domain.ProjectType, error) {
	//TODO implement me
	panic("implement me")
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
