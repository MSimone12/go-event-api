package db

import (
	"go-event-api/src/data/model"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Account{})

	return db
}

func Reset() {
	os.Remove("db.sqlite")
	os.Create("db.sqlite")
}
