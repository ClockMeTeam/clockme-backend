package users

import (
	"context"
	"github.com/clockme/clockme-backend/internal/common"
	"github.com/clockme/clockme-backend/internal/db"
	"github.com/rs/zerolog/log"
	"net/http"
)

type UserHandler struct {
	db *db.Queries
}

func NewUserHandler(db *db.Queries) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (u *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	users, err := u.db.GetAllUsers(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get users from database")
	}
	common.RespondWithJSON(w, http.StatusOK, users)
}
