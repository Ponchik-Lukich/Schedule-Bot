package main

import (
	"Telegram/pkg/bot"
	cst "Telegram/pkg/constants"
	"Telegram/pkg/repo"
	"Telegram/pkg/router"
	"Telegram/pkg/storage"
	"Telegram/pkg/storage/postges"
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

// init for serverless function
func init() {
	err := LoadConfiguration()
	if err != nil {
		log.Printf("failed to load configuration: %v", err)
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatalf("log token is rquired")
	}

	db, err := InitializeStorage()
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	repos := repo.NewRepositories(db)

	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	handlerAdapter = router.SetupRouter(botAPI, repos)
}

func LoadConfiguration() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func InitializeStorage() (storage.Storage, error) {
	dbUrl := os.Getenv("DSN")
	if dbUrl == "" {
		return nil, fmt.Errorf("db url is required")
	}

	dbConfig := &postges.Config{
		DSN: dbUrl,
	}

	db, err := storage.NewStorage(cst.Postgres, dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	if err := db.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

var handlerAdapter *httpadapter.HandlerAdapter

// Handler handler for serverless function
func Handler(ctx context.Context, event *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fullPath := event.Headers["X-Envoy-Original-Path"]
	fullPath = strings.Split(fullPath, "?")[0]
	event.Path = fullPath

	response, err := handlerAdapter.ProxyWithContext(ctx, *event)
	return &response, err
}

// local run
func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatalf("log token is rquired")
	}
	log.Println("Bot token loaded")
	db, err := InitializeStorage()
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	repos := repo.NewRepositories(db)

	bot.RunBotLocal(botToken, repos)
}
