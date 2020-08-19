package chat

import (
	"context"

	"github.com/mortawe/chat/internal/models"
)

type Repo interface {
	Create(ctx context.Context, chat *models.Chat) error
	GetChatsByUser(ctx context.Context, userID models.ID) ([]models.Chat, error)
}
