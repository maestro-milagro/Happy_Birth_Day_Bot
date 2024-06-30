package domain

import "time"

type User struct {
	Id       int
	UserName string
	TgID     string
	BirthDay time.Time
}
