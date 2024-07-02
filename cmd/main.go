package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"hb_bot/internal/app"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
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

	go application.Run(bot, context.Background())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	logger.Info("stopping app", slog.String("signal", sign.String()))

	application.Stop(bot)

	logger.Info("app stopped")
}
