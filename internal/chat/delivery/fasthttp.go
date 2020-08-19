package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mortawe/chat/internal/chat"
	"github.com/mortawe/chat/internal/errors/apierr"
	"github.com/mortawe/chat/internal/errors/dberr"
	"github.com/mortawe/chat/internal/models"

	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type ChatHandler struct {
	chatUC chat.UC
}

func NewChatHandler(cUC chat.UC) *ChatHandler {
	return &ChatHandler{chatUC: cUC}
}

func (h *ChatHandler) Register(r *router.Router) {
	r.POST("/chats/add", h.Create)
	r.POST("/chats/get", h.GetChatsByUserID)
}

func (h *ChatHandler) Create(ctx *fasthttp.RequestCtx) {
	chat := &models.Chat{}
	if err := json.Unmarshal(ctx.PostBody(), &chat); err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}
	err := h.chatUC.Create(ctx, chat)
	if dberr.IsUniqueViolationErr(err) {
		ctx.Error(apierr.NameAlreadyInUse, http.StatusConflict)
		return
	}
	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		logrus.Error(err)
		return
	}
	ctx.WriteString(fmt.Sprint(chat.ID))
	ctx.SetStatusCode(http.StatusOK)
}

type GetByUserArgs struct {
	User models.ID
}

func (h *ChatHandler) GetChatsByUserID(ctx *fasthttp.RequestCtx) {
	args := &GetByUserArgs{}
	if err := json.Unmarshal(ctx.PostBody(), &args); err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}
	chat, err := h.chatUC.GetChatsByUser(ctx, args.User)
	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		logrus.Error(err)
		return
	}
	data, err := json.Marshal(chat)
	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		logrus.Error(err)
		return
	}
	ctx.Write(data)
	ctx.SetStatusCode(http.StatusOK)
}
