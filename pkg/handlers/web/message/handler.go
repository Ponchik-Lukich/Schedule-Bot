package message

import (
	"Telegram/pkg/bot"
	"Telegram/pkg/errors"
	"Telegram/pkg/repo/room"
	"Telegram/pkg/repo/user"
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
)

type Handler interface {
	HandleUpdate(ctx *gin.Context)
}

type handler struct {
	bot      *tgbotapi.BotAPI
	roomRepo room.Repository
	userRepo user.Repository
}

func NewHandler(bot *tgbotapi.BotAPI, userRepo user.Repository, roomRepo room.Repository) Handler {
	return &handler{bot: bot, userRepo: userRepo, roomRepo: roomRepo}
}

func (h *handler) HandleUpdate(ctx *gin.Context) {
	var update tgbotapi.Update

	err := json.NewDecoder(ctx.Request.Body).Decode(&update)
	if err != nil {
		errors.HandleError(ctx, http.StatusBadRequest, errors.InvalidJson, err)
		return
	}

	if update.Message.IsCommand() && update.Message.Command() == "start" {
		// add user
		return
	}

	bot.HandleUpdate(h.bot, &update)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
