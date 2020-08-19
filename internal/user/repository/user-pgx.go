package repository

import (
	"context"

	"github.com/mortawe/chat/internal/models"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

const (
	createUserQ = "INSERT INTO users (username) VALUES (:username) RETURNING id"
)

func (r *UserRepo) Create(ctx context.Context, user *models.User) error {
	query, args, err := r.db.BindNamed(createUserQ, &user)
	if err != nil {
		return err
	}
	return r.db.GetContext(ctx, &user.ID, query, args...)
}
