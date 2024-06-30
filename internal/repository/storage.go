package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"hb_bot/internal/domain"
	"time"
)

type HappyBDayDB interface {
	SaveUser(ctx context.Context, tgid string, userName string, birthDay time.Time) (uid int, err error)
	SubSome(ctx context.Context, tgid string)
	GetAll(ctx context.Context) (users []domain.User)
	GetByID(ctx context.Context, tgid string) (users []domain.User, err error)
	GetWhoSub(ctx context.Context, tgid string) (users []domain.User, err error)
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.New"
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

type Storage struct {
	db *sql.DB
}

func (s Storage) SaveUser(ctx context.Context, tgid string, userName string, birthDay time.Time) (uid int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) SubSome(ctx context.Context, tgid string) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetAll(ctx context.Context) (users []domain.User) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetByID(ctx context.Context, tgid string) (users []domain.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetWhoSub(ctx context.Context, tgid string) (users []domain.User, err error) {
	//TODO implement me
	panic("implement me")
}
