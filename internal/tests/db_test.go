package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"hb_bot/internal/repository"
	"log"
	"testing"
	"time"
)

const layout = "2006-01-02"

func TestGetByID(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		userName      string
		tgUserName    string
		BirthDay      time.Time
		expectedError string
	}{
		{
			name:          "Bad arg",
			userName:      "S K",
			tgUserName:    "SK",
			BirthDay:      timePars("2000-01-01"),
			expectedError: "user not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := storage.GetByID(context.Background(), tt.tgUserName)
			require.NoError(t, err)
			require.NotEmpty(t, user)
		})
	}
}

func TestGetByIDFail(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		userName      string
		tgUserName    string
		BirthDay      time.Time
		expectedError string
	}{
		{
			name:          "Unregistered user",
			userName:      "ds",
			tgUserName:    "@ds",
			BirthDay:      timePars("2002-01-01"),
			expectedError: "user not found",
		},
		{
			name:          "Empty tg name",
			userName:      "hg",
			tgUserName:    "",
			BirthDay:      timePars("2001-02-01"),
			expectedError: "user not found",
		},
		{
			name:          "Bad arg",
			userName:      "S K",
			tgUserName:    "?SK",
			BirthDay:      timePars("2000-01-01"),
			expectedError: "user not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := storage.GetByID(context.Background(), tt.tgUserName)
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedError)
		})
	}
}

func TestSaveUser(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		userName      string
		tgUserName    string
		BirthDay      time.Time
		expectedError string
	}{
		{
			name:          "user already exist",
			userName:      "Rt",
			tgUserName:    "Rt",
			BirthDay:      timePars("2002-01-01"),
			expectedError: "user already exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.SaveUser(context.Background(), tt.userName, tt.tgUserName, tt.BirthDay)
			require.NoError(t, err)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		userName      string
		tgUserName    string
		BirthDay      time.Time
		expectedError string
	}{
		{
			name:          "user already exist",
			userName:      "Rt",
			tgUserName:    "Rt",
			BirthDay:      timePars("2002-01-01"),
			expectedError: "user already exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.Delete(context.Background(), tt.userName)
			require.NoError(t, err)
		})
	}
}

func TestSaveUserFail(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		userName      string
		tgUserName    string
		BirthDay      time.Time
		expectedError string
	}{
		{
			name:          "User already exist",
			userName:      "RN",
			tgUserName:    "RN",
			BirthDay:      timePars("2002-01-01"),
			expectedError: "user already exist",
		},
		{
			name:          "User already exist",
			userName:      "F G",
			tgUserName:    "FG",
			BirthDay:      timePars("2002-01-02"),
			expectedError: "user already exist",
		},
		{
			name:          "User already exist",
			userName:      "S K",
			tgUserName:    "SK",
			BirthDay:      timePars("2000-01-01"),
			expectedError: "user already exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.SaveUser(context.Background(), tt.userName, tt.tgUserName, tt.BirthDay)
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedError)
		})
	}
}

func TestSubSome(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		tgName        string
		subTgName     string
		expectedError string
	}{
		{
			name:          "user already exist",
			tgName:        "SK",
			subTgName:     "FG",
			expectedError: "user already exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.SubSome(context.Background(), tt.tgName, tt.subTgName)
			require.NoError(t, err)
		})
	}
}

func TestUnSubSome(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		tgName        string
		subTgName     string
		expectedError string
	}{
		{
			name:          "user already exist",
			tgName:        "SK",
			subTgName:     "FG",
			expectedError: "user already exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.UnSubSome(context.Background(), tt.tgName, tt.subTgName)
			require.NoError(t, err)
		})
	}
}

func TestSubSomeFail(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		tgName        string
		subTgName     string
		expectedError string
	}{
		{
			name:          "user already exist",
			tgName:        "RN",
			subTgName:     "FG",
			expectedError: "user already exist",
		},
		{
			name:          "user already exist",
			tgName:        "FG",
			subTgName:     "RN",
			expectedError: "user already exist",
		},
		{
			name:          "user already exist",
			tgName:        "FG",
			subTgName:     "SK",
			expectedError: "user already exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.SubSome(context.Background(), tt.tgName, tt.subTgName)
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedError)
		})
	}
}

func TestGetAll(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	t.Run("Happy path", func(t *testing.T) {
		users, err := storage.GetAll(context.Background())
		require.NoError(t, err)
		require.NotEmpty(t, users)
	})
}

func TestGetWhoSub(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		userName      string
		tgUserName    string
		BirthDay      time.Time
		expectedError string
	}{
		{
			name:          "Happy path",
			userName:      "F G",
			tgUserName:    "FG",
			BirthDay:      timePars("2002-01-02"),
			expectedError: "no users found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users, err := storage.GetWhoSub(context.Background(), tt.tgUserName)
			require.NoError(t, err)
			require.NotEmpty(t, users)
		})
	}
}

func TestGetWhoSubFail(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		userName      string
		tgUserName    string
		BirthDay      time.Time
		expectedError string
	}{
		{
			name:          "Unregistered user",
			userName:      "ds",
			tgUserName:    "DS",
			BirthDay:      timePars("2002-01-01"),
			expectedError: "no users found",
		},
		{
			name:          "Empty tg name",
			userName:      "hg",
			tgUserName:    "",
			BirthDay:      timePars("2001-02-01"),
			expectedError: "no users found",
		},
		{
			name:          "Unsubed user",
			userName:      "S K",
			tgUserName:    "SK",
			BirthDay:      timePars("2000-01-01"),
			expectedError: "no users found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users, err := storage.GetWhoSub(context.Background(), tt.tgUserName)
			require.NoError(t, err)
			require.Empty(t, users)
		})
	}
}

func TestIsSub(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		tgName        string
		subTgName     string
		expectedError string
	}{
		{
			name:          "user already exist",
			tgName:        "RN",
			subTgName:     "FG",
			expectedError: "user not found",
		},
		{
			name:          "user already exist",
			tgName:        "FG",
			subTgName:     "RN",
			expectedError: "user not found",
		},
		{
			name:          "user already exist",
			tgName:        "FG",
			subTgName:     "SK",
			expectedError: "user not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, err := storage.IsSub(context.Background(), tt.tgName, tt.subTgName)
			require.NoError(t, err)
			require.True(t, flag)
		})
	}
}

func TestIsSubFail(t *testing.T) {
	storage, err := repository.New("../../storage/hb.db")
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		name          string
		tgName        string
		subTgName     string
		expectedError string
	}{
		{
			name:          "Don't follow",
			tgName:        "SK",
			subTgName:     "FG",
			expectedError: "user not found",
		},
		{
			name:          "Bad arg",
			tgName:        "",
			subTgName:     "RN",
			expectedError: "user not found",
		},
		{
			name:          "Same user",
			tgName:        "SK",
			subTgName:     "SK",
			expectedError: "user not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, err := storage.IsSub(context.Background(), tt.tgName, tt.subTgName)
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedError)
			require.False(t, flag)
		})
	}
}

func timePars(date string) time.Time {
	ans, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
	}
	return ans
}
