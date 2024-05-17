package db

import (
	"testing"
)

func TestPostgres(t *testing.T) {
	type model struct {
		Model
		Name string `gorm:"unique"`
		Code string
	}

	db, err := NewDB(
		&Postgres{
			Host:     "localhost:5432",
			User:     "user",
			Pass:     "pass",
			DBName:   "db",
			SSL:      "disable",
			TimeZone: "Asia/Tehran",
		},
		&model{},
	)
	if err != nil {
		t.Error(err)
		return
	}

	db.Create(&model{
		Name: "user_1",
		Code: "code_1",
	})

	var user model
	if db.First(&user, "name = ?", "user_1").Error != nil {
		t.Error("could not find model")
	}
	if user.Name != "user_1" || user.Code != "code_1" {
		t.Error("found a wrong model")
	}
}
