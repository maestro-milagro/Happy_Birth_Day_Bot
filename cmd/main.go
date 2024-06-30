package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"hb_bot/internal/app"
	"log"
	"log/slog"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	bot, err := tgbotapi.NewBotAPI(os.Getenv("SECRET_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	application := app.New(logger, "./storage/hb.db")

	application.Run(bot, context.Background())
}
