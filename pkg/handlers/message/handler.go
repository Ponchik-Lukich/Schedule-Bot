package message

import (
	"Telegram/pkg/errors"
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
)

func HandleMessage(ctx *gin.Context) {
	var update tgbotapi.Update

	err := json.NewDecoder(ctx.Request.Body).Decode(&update)
	if err != nil {
		errors.HandleError(ctx, http.StatusBadRequest, errors.InvalidJson, err)
		return
	}

	//bot.HandleUpdate(bot, &update)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
