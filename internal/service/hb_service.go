package service

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hb_bot/internal/repository"
	"log/slog"
	"time"
)

type HBService interface {
	AddUser(ctx context.Context, tgName string, userName string, birthDay time.Time, chatID int64) (msg tgbotapi.MessageConfig, err error)
	Sub(ctx context.Context, tgid string, chatID int64) (msg tgbotapi.MessageConfig, err error)
	All(ctx context.Context, chatID int64) (msg tgbotapi.MessageConfig, err error)
	ByID(ctx context.Context, tgid string, chatID int64) (msg tgbotapi.MessageConfig, err error)
	WhoSub(ctx context.Context, tgid string, chatID int64) (msg tgbotapi.MessageConfig, err error)
}

type Service struct {
	logger  *slog.Logger
	storage repository.HappyBDayDB
}

func (s Service) AddUser(
	ctx context.Context,
	tgName string,
	userName string,
	birthDay time.Time,
	chatID int64,
) (msg tgbotapi.MessageConfig, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Sub(
	ctx context.Context,
	userName string,
	chatID int64,
) (msg tgbotapi.MessageConfig, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) All(
	ctx context.Context,
	chatID int64,
) (msg tgbotapi.MessageConfig, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ByID(
	ctx context.Context,
	userName string,
	chatID int64,
) (msg tgbotapi.MessageConfig, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) WhoSub(
	ctx context.Context,
	userName string,
	chatID int64,
) (msg tgbotapi.MessageConfig, err error) {
	//TODO implement me
	panic("implement me")
}

func New(
	logger *slog.Logger,
	storage repository.HappyBDayDB,
) *Service {
	return &Service{logger: logger, storage: storage}
}
