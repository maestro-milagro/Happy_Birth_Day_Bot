package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"hb_bot/internal/domain"
	"log"
	"time"
)

var (
	ErrUserExist    = errors.New("user already exist")
	ErrUserNotFound = errors.New("user not found")
)

type HappyBDayDB interface {
	SaveUser(ctx context.Context, tgName string, userName string, birthDay time.Time) (err error)
	SubSome(ctx context.Context, tgName string, subTgName string) (err error)
	GetAll(ctx context.Context) ([]domain.User, error)
	GetByID(ctx context.Context, tgName string) (domain.User, error)
	GetWhoSub(ctx context.Context, tgName string) ([]domain.User, error)
	UnSubSome(ctx context.Context, tgName string, subTgName string) (err error)
}

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.New"
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(ctx context.Context, tgName string, userName string, birthDay time.Time) (err error) {
	stmt, err := s.db.Prepare("INSERT INTO users(user_name, tg_user_name, birth_day) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, tgName, userName, birthDay)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return ErrUserExist
		}
		return err
	}

	return nil
}

func (s *Storage) SubSome(ctx context.Context, tgName string, subTgName string) (err error) {
	stmt, err := s.db.Prepare("INSERT INTO subscriptions(tg_id, sub_tg_id) VALUES(?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, tgName, subTgName)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return ErrUserExist
		}
		return err
	}

	return nil
}

func (s *Storage) GetAll(ctx context.Context) ([]domain.User, error) {
	//stmt, err := s.db.Prepare("SELECT COUNT(*) FROM users ")
	//
	//if err != nil {
	//	return []domain.User{}, err
	//}
	//
	//row := stmt.QueryRowContext(ctx)
	//
	//var i int64
	//
	//err = row.Scan(&i)

	//stmt, err = s.db.Prepare("SELECT * FROM users ")
	//
	//if err != nil {
	//	return []domain.User{}, err
	//}
	//
	//row = stmt.QueryRowContext(ctx)

	//users := make([]domain.User, i)

	var users []domain.User

	rows, err := s.db.Query("SELECT * FROM users ")

	for rows.Next() { // Iterate and fetch the records from result cursor
		item := domain.User{}
		err := rows.Scan(&item.UserName, &item.TgUserName, &item.BirthDay)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, item)
	}

	//for a := range users {
	//	err = row.Scan(&users[a].UserName, &users[a].TgUserName, &users[a].BirthDay)
	//}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.User{}, ErrUserNotFound
		}

		return []domain.User{}, err
	}
	return users, nil
}

func (s *Storage) GetByID(ctx context.Context, tgName string) (domain.User, error) {
	stmt, err := s.db.Prepare("SELECT user_name, tg_user_name, birth_day FROM users WHERE tg_user_name = ? ")

	if err != nil {
		return domain.User{}, err
	}

	row := stmt.QueryRowContext(ctx, tgName)

	var user domain.User

	err = row.Scan(&user.UserName, &user.TgUserName, &user.BirthDay)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}

		return domain.User{}, err
	}
	return user, nil
}

func (s *Storage) GetWhoSub(ctx context.Context, tgName string) ([]domain.User, error) {
	//stmt, err := s.db.Prepare("SELECT COUNT(*) FROM subscriptions AS s LEFT JOIN users AS u ON s.sub_tg_id = u.tg_user_name WHERE s.tg_id = ?")
	//
	//if err != nil {
	//	return []domain.User{}, err
	//}
	//
	//row := stmt.QueryRowContext(ctx, tgName)
	//
	//var i int64
	//
	//err = row.Scan(&i)
	//
	//stmt, err = s.db.Prepare("SELECT u.user_name, s.sub_tg_id, u.birth_day FROM subscriptions AS s LEFT JOIN users AS u ON s.sub_tg_id = u.tg_user_name WHERE s.tg_id = ?")
	//
	//if err != nil {
	//	return []domain.User{}, err
	//}
	//
	//row = stmt.QueryRowContext(ctx, tgName)
	//
	//users := make([]domain.User, i)
	//
	//for a := range users {
	//	err = row.Scan(&users[a].UserName, &users[a].TgUserName, &users[a].BirthDay)
	//}
	var users []domain.User

	rows, err := s.db.Query("SELECT u.user_name, s.sub_tg_id, u.birth_day FROM subscriptions AS s LEFT JOIN users AS u ON s.sub_tg_id = u.tg_user_name WHERE s.tg_id = ?", tgName)

	for rows.Next() { // Iterate and fetch the records from result cursor
		item := domain.User{}
		err := rows.Scan(&item.UserName, &item.TgUserName, &item.BirthDay)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, item)
	}
	fmt.Println(users)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.User{}, ErrUserNotFound
		}

		return []domain.User{}, err
	}
	return users, nil
}

func (s *Storage) UnSubSome(ctx context.Context, tgName string, subTgName string) (err error) {
	stmt, err := s.db.Prepare("DELETE FROM subscriptions WHERE tg_id = ? AND sub_tg_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, tgName, subTgName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}
