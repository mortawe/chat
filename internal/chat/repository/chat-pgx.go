package repository

import (
	"context"
	"fmt"

	"github.com/mortawe/chat/internal/models"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

type ChatRepo struct {
	db *sqlx.DB
}

func NewChatRepo(db *sqlx.DB) *ChatRepo {
	return &ChatRepo{db: db}
}

const (
	createChatQ      = "INSERT INTO chats (name) VALUES (:name) RETURNING id"
	insertChatUsersQ = "INSERT INTO chat_users (chat_id, user_id) VALUES "
	getChatsByUserQ  = `SELECT chat_id, created_at, name FROM chat_users 
JOIN chats ON chat_id = chats.id
WHERE user_id = :user_id ORDER BY created_at`
)

func buildQueryInsertValues(chatID models.ID, ids []models.ID) string {
	res := ""
	for _, id := range ids {
		res = fmt.Sprint(res, "(", chatID, ",", id, "),")
	}
	if len(res) > 0 {
		return res[:len(res)-1]
	}
	return ""
}

func (r *ChatRepo) Create(ctx context.Context, chat *models.Chat) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query, args, err := tx.BindNamed(createChatQ, &chat)
	if err != nil {
		return err
	}

	if err := tx.GetContext(ctx, &chat.ID, query, args...); err != nil {
		return err
	}

	values := buildQueryInsertValues(chat.ID, chat.Users)
	if _, err := tx.ExecContext(ctx, insertChatUsersQ+values); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ChatRepo) GetChatsByUser(ctx context.Context, userID models.ID) ([]models.Chat, error) {
	chats := []models.Chat{}
	query, args, err := r.db.BindNamed(getChatsByUserQ, &userID)
	if err != nil {
		return nil, err
	}
	if err := r.db.GetContext(ctx, &chats, query, args...); err != nil {
		return nil, err
	}
	return chats, err
}
