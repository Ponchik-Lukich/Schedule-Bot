package main

import (
	"Telegram/pkg/bot"
	"Telegram/pkg/router"
	"Telegram/pkg/storage"
	"Telegram/pkg/storage/ydb"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

var handlerAdapter *httpadapter.HandlerAdapter

func LoadConfiguration() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func InitializeStorage() (storage.Storage, error) {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, fmt.Errorf("db url is required")
	}

	dbConfig := &ydb.Config{
		Database: dbUrl,
	}

	db, err := storage.NewStorage("ydb", dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	fmt.Println("Connecting to database")
	if err := db.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func Handler(ctx context.Context, event *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fullPath := event.Headers["X-Envoy-Original-Path"]
	fullPath = strings.Split(fullPath, "?")[0]
	event.Path = fullPath

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatalf("log token is rquired")
	}

	_, err := InitializeStorage()
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	if handlerAdapter == nil {
		handlerAdapter = router.SetupRouter(botAPI)
	}

	response, err := handlerAdapter.ProxyWithContext(ctx, *event)
	return &response, err
}

func main() {
	err := LoadConfiguration()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	log.Println("Configuration loaded")

	_, err = InitializeStorage()
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}
	log.Println("Storage initialized")

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatalf("log token is rquired")
	}
	log.Println("Bot token loaded")

	bot.RunBotLocal(botToken)
}
