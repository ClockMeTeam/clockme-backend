package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maevlava/ftf-clockify/internal/domain"
	"github.com/maevlava/ftf-clockify/internal/repository/postgres/db"
)

//Interfaces

type PgUserRepository struct {
	q *db.Queries
}

func (r *PgUserRepository) GetUserProjects(ctx context.Context, name string) ([]domain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func NewPgUserRepository(pool *pgxpool.Pool) domain.UserRepository {
	return &PgUserRepository{
		q: db.New(pool),
	}
}
func (r *PgUserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	dbUsers, err := r.q.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting users from db %v", err)
	}
	var users []domain.User

	for _, user := range dbUsers {
		users = append(users, domain.User{
			ID:         user.ID,
			ClockifyID: user.ClockifyID,
			Name:       user.Name,
			Email:      user.Email,
		})
	}

	return users, nil
}
func (r *PgUserRepository) GetUser(ctx context.Context, name string) (domain.User, error) {
	dbUser, err := r.q.GetUser(ctx, name)
	if err != nil {
		return domain.User{}, errors.New("error getting user from db")
	}
	return domain.User{
		ID:         dbUser.ID,
		ClockifyID: dbUser.ClockifyID,
		Name:       dbUser.Name,
		Email:      dbUser.Email,
	}, nil
}
func (r *PgUserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	dbUserParam := db.CreateUserParams{
		ID:         user.ID,
		ClockifyID: user.ClockifyID,
		Name:       user.Name,
		Email:      user.Email,
	}

	dbUser, err := r.q.CreateUser(ctx, dbUserParam)

	if err != nil {
		return domain.User{}, errors.New("error creating user in db")
	}

	return domain.User{
		ID:         dbUser.ID,
		ClockifyID: dbUser.ClockifyID,
		Name:       dbUser.Name,
		Email:      dbUser.Email,
	}, nil
}
func (r *PgUserRepository) DeleteAllUsers(ctx context.Context) error {
	err := r.q.DeleteAllUsers(ctx)
	if err != nil {
		return errors.New("error deleting all users from db")
	}
	return nil
}
