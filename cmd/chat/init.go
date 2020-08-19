package main

import (
	. "github.com/mortawe/chat/internal/chat/delivery"
	. "github.com/mortawe/chat/internal/chat/repository"
	. "github.com/mortawe/chat/internal/chat/usecase"
	. "github.com/mortawe/chat/internal/message/delivery"
	. "github.com/mortawe/chat/internal/message/repository"
	. "github.com/mortawe/chat/internal/message/usecase"
	. "github.com/mortawe/chat/internal/user/delivery"
	. "github.com/mortawe/chat/internal/user/repository"
	. "github.com/mortawe/chat/internal/user/usecase"

	"github.com/fasthttp/router"
	"github.com/jmoiron/sqlx"
)

func register(db *sqlx.DB, r *router.Router) (*UserHandler, *ChatHandler) {
	// user
	uR := NewUserRepo(db)
	uU := NewUserUC(uR)
	uH := NewUserHandler(uU)
	uH.Register(r)
	// chat
	cR := NewChatRepo(db)
	cU := NewChatUC(cR)
	cH := NewChatHandler(cU)
	cH.Register(r)
	// message
	mR := NewMsgRepo(db)
	mU := NewMsgUC(mR)
	mH := NewMsgHandler(mU)
	mH.Register(r)

	return uH, cH
}
