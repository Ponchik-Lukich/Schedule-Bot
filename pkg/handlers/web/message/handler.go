package message

import (
	"Telegram/pkg/bot"
	"Telegram/pkg/errors"
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
)

type Handler interface {
	HandleUpdate(ctx *gin.Context)
}

type handler struct {
	bot *tgbotapi.BotAPI
}

func NewHandler(bot *tgbotapi.BotAPI) Handler {
	return &handler{bot: bot}
}

func (h *handler) HandleUpdate(ctx *gin.Context) {
	var update tgbotapi.Update

	err := json.NewDecoder(ctx.Request.Body).Decode(&update)
	if err != nil {
		errors.HandleError(ctx, http.StatusBadRequest, errors.InvalidJson, err)
		return
	}

	bot.HandleUpdate(h.bot, &update)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
