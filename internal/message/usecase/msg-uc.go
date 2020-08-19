package usecase

import (
	"context"

	"github.com/mortawe/chat/internal/message"
	"github.com/mortawe/chat/internal/models"
)

type MsgUC struct {
	msgRepo message.Repo
}

func NewMsgUC(mR message.Repo) *MsgUC {
	return &MsgUC{msgRepo: mR}
}

func (u *MsgUC) Create(ctx context.Context, msg *models.Message) error {
	return u.msgRepo.Create(ctx, msg)
}

func (u *MsgUC) GetByChat(ctx context.Context, chatID models.ID) ([]models.Message, error) {
	return u.msgRepo.GetByChat(ctx, chatID)
}
