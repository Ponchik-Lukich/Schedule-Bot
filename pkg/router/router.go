package router

import (
	"Telegram/pkg/handlers/message"
	"Telegram/pkg/handlers/test"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetupRouter(bot *tgbotapi.BotAPI) *httpadapter.HandlerAdapter {
	router := gin.Default()

	messageH := message.NewHandler(bot)

	router.POST("/", messageH.HandleUpdate)
	router.GET("/", test.Test)

	return httpadapter.New(router)
}