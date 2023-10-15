package web

import (
	"Telegram/pkg/errors"
	"Telegram/pkg/handlers/telegram/commands"
	"Telegram/pkg/repo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
)

type Handler interface {
	HandleUpdate(ctx *gin.Context)
}

type handler struct {
	bot   *tgbotapi.BotAPI
	repos repo.Repositories
}

func NewHandler(bot *tgbotapi.BotAPI, repos repo.Repositories) Handler {
	return &handler{bot: bot, repos: repos}
}

func (h *handler) HandleUpdate(ctx *gin.Context) {
	var update tgbotapi.Update

	err := json.NewDecoder(ctx.Request.Body).Decode(&update)
	if err != nil {
		errors.HandleError(ctx, http.StatusBadRequest, errors.InvalidJson, err)
		return
	}

	if update.Message.IsCommand() && update.Message.Command() == "start" {
		commands.HandleStartCom(h.bot, &update, h.repos)
	} else {
		commands.HandleUpdateCom(h.bot, &update, h.repos)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
