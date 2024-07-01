package domain

import "time"

type User struct {
	UserName   string
	TgUserName string
	BirthDay   time.Time
}
