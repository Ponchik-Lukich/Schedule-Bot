package router

import (
	"Telegram/pkg/handlers/web"
	"Telegram/pkg/repo"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetupRouter(bot *tgbotapi.BotAPI, repos repo.Repositories) *httpadapter.HandlerAdapter {
	router := gin.Default()

	messageH := web.NewHandler(bot, repos)

	router.POST("/", messageH.HandleUpdate)

	return httpadapter.New(router)
}
