package app

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hb_bot/internal/repository"
	"hb_bot/internal/service"
	"log"
	"log/slog"
	"time"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

const layout = "2006-01-02"

type App struct {
	hb_service service.HBService
}

func New(logger *slog.Logger,
	storagePath string,
) *App {
	storage, err := repository.New(storagePath)
	if err != nil {
		panic(err)
	}
	hbService := service.New(logger, storage)
	return &App{
		hb_service: hbService,
	}
}

func (a *App) Run(bot *tgbotapi.BotAPI, ctx context.Context) {
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Command() {
		case "open":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "send_2":
			go func() {
				tm := time.NewTimer(time.Minute)
				<-tm.C
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "2")
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}()
		case "all":
			msg, err := a.hb_service.All(ctx, update.Message.Chat.ID)
			if err != nil {
				log.Panic(err)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		case "add":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter First Name")
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			some := <-updates
			tgName := some.Message.Text
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Last Name")
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			some = <-updates
			tgName += " " + some.Message.Text
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Telegram UserName (starts with @)")
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			some = <-updates
			userName := some.Message.Text
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Birth Day in format YEAR-MONTH-DAY")
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			some = <-updates
			birthDayS := some.Message.Text
			birthDay, err := time.Parse(layout, birthDayS)
			fmt.Println(birthDay)
			msg, err = a.hb_service.AddUser(ctx, tgName, userName, birthDay, update.Message.Chat.ID)
			if err != nil {
				log.Panic(err)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		case "sub":
		case "tg_id":
		case "subs":

		}
		//if _, err := bot.Send(msg); err != nil {
		//	log.Panic(err)
		//}
	}
}
