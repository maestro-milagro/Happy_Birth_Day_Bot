package app

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hb_bot/internal/repository"
	"hb_bot/internal/service"
	"log"
	"log/slog"
	"strings"
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
		case "start":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Birth Day in format YEAR-MONTH-DAY")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			some := <-updates
			birthDayS := some.Message.Text
			birthDay, err := time.Parse(layout, birthDayS)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(birthDay)
			msg, err = a.hb_service.AddUser(ctx,
				update.Message.Chat.FirstName+" "+update.Message.Chat.LastName,
				update.Message.Chat.UserName,
				birthDay,
				update.Message.Chat.ID)
			if err != nil {
				log.Println(err)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		case "all":
			msg, err := a.hb_service.All(ctx, update.Message.Chat.ID)
			if err != nil {
				log.Println(err)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		case "add":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter First Name")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			some := <-updates
			tgName := some.Message.Text
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Last Name")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			some = <-updates
			tgName += " " + some.Message.Text
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Telegram UserName (starts with @)")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			some = <-updates
			userName := strings.Trim(some.Message.Text, "@")
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Birth Day in format YEAR-MONTH-DAY")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			some = <-updates
			birthDayS := some.Message.Text
			birthDay, err := time.Parse(layout, birthDayS)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(birthDay)
			msg, err = a.hb_service.AddUser(ctx, tgName, userName, birthDay, update.Message.Chat.ID)
			if err != nil {
				log.Println(err)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		case "sub":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Telegram UserName (starts with @)")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			some := <-updates
			userName := strings.Trim(some.Message.Text, "@")
			msg, err := a.hb_service.Sub(ctx, update.Message.Chat.UserName, userName, update.Message.Chat.ID)
			if err != nil {
				log.Println(err)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			go func() {
				msg, err = a.hb_service.IsSub(ctx, update.Message.Chat.UserName, userName, update.Message.Chat.ID)
				if err != nil {
					log.Println(err)
				}
				if _, err := bot.Send(msg); err != nil {
					log.Println(err)
				}
			}()
		case "tg_id":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Telegram UserName (starts with @)")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			some := <-updates
			userName := strings.Trim(some.Message.Text, "@")
			msg, err := a.hb_service.ByID(ctx, userName, update.Message.Chat.ID)
			if err != nil {
				log.Println(err)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		case "subs":
			msg, err := a.hb_service.WhoSub(ctx, update.Message.Chat.UserName, update.Message.Chat.ID)
			if err != nil {
				log.Println(err)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		case "unsub":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Telegram UserName (starts with @)")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			some := <-updates
			userName := strings.Trim(some.Message.Text, "@")
			msg, err := a.hb_service.Unsub(ctx, update.Message.Chat.UserName, userName, update.Message.Chat.ID)
			if err != nil {
				log.Println(err)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		}
		//if _, err := bot.Send(msg); err != nil {
		//	log.Panic(err)
		//}
	}
}
