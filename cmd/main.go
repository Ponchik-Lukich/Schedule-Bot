package main

import (
	"Telegram/pkg/bot"
	"Telegram/pkg/router"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var handlerAdapter *httpadapter.HandlerAdapter

func Handler(ctx context.Context, event *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatalf("log token is rquired")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	if handlerAdapter == nil {
		handlerAdapter = router.SetupRouter(bot)
	}

	response, err := handlerAdapter.ProxyWithContext(ctx, *event)
	return &response, err
}

func main() {
	err := godotenv.Load("cmd/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		panic("Bot token is required")
	}

	bot.RunBotLocal(botToken)
}
