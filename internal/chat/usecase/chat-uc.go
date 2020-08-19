package usecase

import (
	"context"

	"github.com/mortawe/chat/internal/chat"
	"github.com/mortawe/chat/internal/models"
)

type ChatUC struct {
	chatRepo chat.Repo
}

func NewChatUC(cR chat.Repo) *ChatUC {
	return &ChatUC{chatRepo: cR}
}

func (u *ChatUC) Create(ctx context.Context, chat *models.Chat) error {
	return u.chatRepo.Create(ctx, chat)
}

func (u *ChatUC) GetChatsByUser(ctx context.Context, userID models.ID) ([]models.Chat, error) {
	return u.chatRepo.GetChatsByUser(ctx, userID)
}
