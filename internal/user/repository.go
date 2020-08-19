package user

import (
	"context"

	"github.com/mortawe/chat/internal/models"
)

type Repo interface {
	Create(ctx context.Context, user *models.User) error
}
