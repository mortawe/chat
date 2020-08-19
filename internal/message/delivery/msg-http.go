package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/mortawe/chat/internal/errors/apierr"
	"github.com/mortawe/chat/internal/errors/dberr"
	"github.com/mortawe/chat/internal/message"
	"github.com/mortawe/chat/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"net/http"
)

type MsgHandler struct {
	msgUC message.UC
}

func NewMsgHandler(mU message.UC) *MsgHandler {
	return &MsgHandler{msgUC: mU}
}

func (h *MsgHandler) Register(r *router.Router) {
	r.POST("/messages/add", h.Create)
	r.POST("/messages/get", h.List)
}

func (h *MsgHandler) Create(ctx *fasthttp.RequestCtx) {
	msg := &models.Message{}
	if err := json.Unmarshal(ctx.PostBody(), &msg); err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}
	err := h.msgUC.Create(ctx, msg)
	if dberr.IsUniqueViolationErr(err) {
		ctx.Error(apierr.NoSuchUserOrChat, http.StatusConflict)
		return
	}
	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		logrus.Error(err)
		return
	}
	ctx.WriteString(fmt.Sprint(msg.ID))
	ctx.SetStatusCode(http.StatusOK)
}

type ListArgs struct {
	ChatID models.ID `json:"chat_id"`
}

func (h *MsgHandler) List(ctx *fasthttp.RequestCtx) {
	args := &ListArgs{}
	if err := json.Unmarshal(ctx.PostBody(), args); err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}
	msgList, err := h.msgUC.GetByChat(ctx, args.ChatID)
	if dberr.IsForeignKeyViolation(err) {
		ctx.Error(apierr.NoSuchUserOrChat, http.StatusConflict)
		return
	}
	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		logrus.Error(err)
		return
	}
	data, err := json.Marshal(msgList)
	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		logrus.Error(err)
		return
	}
	ctx.Write(data)
	ctx.SetStatusCode(http.StatusOK)
}
