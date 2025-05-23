package domain

import (
	"context"
	"time"
)

// todo change from ID to clockify ID and create UUID for current ID
// todo add start date
type User struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	HoursOwed time.Duration `json:"hours_owed"`
	//Projects  []Project        `json:"projects"`
}
type UserRepository interface {
	GetUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, name string) (User, error)
	CreateUser(ctx context.Context, user User) (User, error)
	DeleteAllUsers(ctx context.Context) error
}
