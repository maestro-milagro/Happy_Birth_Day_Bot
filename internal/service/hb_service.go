package service

import (
	"context"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hb_bot/internal/domain"
	"hb_bot/internal/logger/sl"
	"hb_bot/internal/repository"
	"log/slog"
	"time"
)

type HBService interface {
	AddUser(ctx context.Context, tgName string, userName string, birthDay time.Time, chatID int64) (msg tgbotapi.MessageConfig, err error)
	Sub(ctx context.Context, userName string, subName string, chatID int64) (msg tgbotapi.MessageConfig, err error)
	All(ctx context.Context, chatID int64) (msg tgbotapi.MessageConfig, err error)
	ByID(ctx context.Context, userName string, chatID int64) (msg tgbotapi.MessageConfig, err error)
	WhoSub(ctx context.Context, userName string, chatID int64) (msg tgbotapi.MessageConfig, err error)
	Unsub(ctx context.Context, userName string, subName string, chatID int64) (msg tgbotapi.MessageConfig, err error)
	IsSub(ctx context.Context, userName string, subName string, chatID int64) (msg tgbotapi.MessageConfig, err error)
}

type Service struct {
	logger  *slog.Logger
	storage repository.HappyBDayDB
}

func New(
	logger *slog.Logger,
	storage repository.HappyBDayDB,
) *Service {
	return &Service{logger: logger, storage: storage}
}

func (s *Service) AddUser(
	ctx context.Context,
	tgName string,
	userName string,
	birthDay time.Time,
	chatID int64,
) (msg tgbotapi.MessageConfig, err error) {
	err = s.storage.SaveUser(ctx, tgName, userName, birthDay)
	if err != nil {
		s.logger.Error("Error while saving user", sl.Err(err))

		return tgbotapi.MessageConfig{}, err
	}
	s.logger.Info("User add successfully")
	return tgbotapi.NewMessage(chatID, "User add successfully"), nil
}

func (s *Service) Sub(
	ctx context.Context,
	userName string,
	subName string,
	chatID int64,
) (msg tgbotapi.MessageConfig, err error) {
	if _, err = s.storage.GetByID(ctx, subName); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			s.logger.Error("User not found", sl.Err(err))
			return tgbotapi.NewMessage(chatID, "Cant found user with this @username"), err
		}
	}
	err = s.storage.SubSome(ctx, userName, subName)
	if err != nil {
		s.logger.Error("Error while subbing", sl.Err(err))

		return tgbotapi.MessageConfig{}, err
	}
	s.logger.Info("Subbed successfully")
	return tgbotapi.NewMessage(chatID, "User subbed successfully"), nil
}

func (s *Service) All(
	ctx context.Context,
	chatID int64,
) (msg tgbotapi.MessageConfig, err error) {
	users, err := s.storage.GetAll(ctx)
	if err != nil {
		s.logger.Error("Error while getting user list", sl.Err(err))

		return tgbotapi.MessageConfig{}, err
	}
	s.logger.Info("List get successfully")
	return tgbotapi.NewMessage(chatID, s.UsersToString(users)), nil
}

func (s *Service) ByID(
	ctx context.Context,
	userName string,
	chatID int64,
) (msg tgbotapi.MessageConfig, err error) {
	user, err := s.storage.GetByID(ctx, userName)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			s.logger.Error("User not found", sl.Err(err))
			return tgbotapi.NewMessage(chatID, "Cant found user with this @username"), err
		}
		s.logger.Error("Error while getting user by id", sl.Err(err))

		return tgbotapi.MessageConfig{}, err
	}
	s.logger.Info("User by id returned successfully")
	return tgbotapi.NewMessage(chatID, s.UserToString(user)), nil
}

func (s *Service) WhoSub(
	ctx context.Context,
	userName string,
	chatID int64,
) (msg tgbotapi.MessageConfig, err error) {
	users, err := s.storage.GetWhoSub(ctx, userName)
	if err != nil {
		s.logger.Error("Error while getting Subscriptions list", sl.Err(err))

		return tgbotapi.MessageConfig{}, err
	}
	s.logger.Info("Subscriptions list returned successfully")
	if len(users) == 0 {
		return tgbotapi.NewMessage(chatID, "You don't follow anyone"), nil
	}
	return tgbotapi.NewMessage(chatID, s.UsersToString(users)), nil
}

func (s *Service) Unsub(ctx context.Context, userName string, subName string, chatID int64) (msg tgbotapi.MessageConfig, err error) {
	if _, err = s.storage.GetByID(ctx, userName); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			s.logger.Error("User not found", sl.Err(err))
			return tgbotapi.NewMessage(chatID, "Cant found user with this @username"), err
		}
	}
	err = s.storage.UnSubSome(ctx, userName, subName)
	if err != nil {
		s.logger.Error("Error while unsubbing", sl.Err(err))

		return tgbotapi.MessageConfig{}, err
	}
	s.logger.Info("Unsubbed successfully")
	return tgbotapi.NewMessage(chatID, "You are no longer following this user"), nil
}

func (s *Service) IsSub(ctx context.Context, userName string, subName string, chatID int64) (msg tgbotapi.MessageConfig, err error) {
	user, err := s.storage.GetByID(ctx, subName)
	if err != nil {
		s.logger.Error("User not found", sl.Err(err))
		return tgbotapi.MessageConfig{}, err
	}
	years := int(time.Now().Sub(user.BirthDay).Hours()/24/365) + 1
	bd := user.BirthDay.AddDate(years, 0, 0)
	timer := time.NewTimer(bd.Sub(time.Now()))
	<-timer.C

	flag, err := s.storage.IsSub(ctx, userName, subName)
	if err != nil {
		s.logger.Error("Error while isSub check", sl.Err(err))

		return tgbotapi.MessageConfig{}, err
	}
	if !flag {
		s.logger.Warn("User no longer subbing")

		return tgbotapi.MessageConfig{}, nil
	}
	return tgbotapi.NewMessage(chatID, "Today is the user's: "+user.UserName+" birthday!!!"), nil
}

func (s *Service) UserToString(user domain.User) string {
	return fmt.Sprintf(" name: %s, tg name: %s, birth day: %s ", user.UserName, user.TgUserName, user.BirthDay)
}

func (s *Service) UsersToString(users []domain.User) string {
	ans := ""
	for i, v := range users {
		if i == len(users)-1 {
			ans += s.UserToString(v)
			return ans
		}
		ans += s.UserToString(v) + "\n"
	}
	return ans
}
