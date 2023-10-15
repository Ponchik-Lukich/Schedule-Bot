package router

import (
	"Telegram/pkg/handlers/web/message"
	"Telegram/pkg/handlers/web/test"
	"Telegram/pkg/repo"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetupRouter(bot *tgbotapi.BotAPI, repos repo.Repositories) *httpadapter.HandlerAdapter {
	router := gin.Default()

	messageH := message.NewHandler(bot, repos.GetUserRepo(), repos.GetRoomRepo())

	router.POST("/", messageH.HandleUpdate)
	router.GET("/", test.Test)

	return httpadapter.New(router)
}
