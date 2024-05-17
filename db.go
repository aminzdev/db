package db

import (
	"gorm.io/gorm"
)

type Model gorm.Model

type DB struct {
	*gorm.DB
}

func NewDB(driver Driver, models ...interface{}) (*DB, error) {
	db, err := driver.Connect()
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(models...); err != nil {
		return nil, err
	}
	return db, nil
}
