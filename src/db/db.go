package db

import (
	"fmt"
	"go-event-api/src/data/model"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DefaultDatabaseName = "db.sqlite"

type Database struct {
	DB *gorm.DB
}

func NewDatabase(db *gorm.DB) Database {
	migrate(db)
	fmt.Println("Migrated database")

	return Database{
		DB: db,
	}
}

func Connect(name string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{
		AllowGlobalUpdate: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func Reset(name string) {
	createIfNotExists(name)
	db := Connect(name)
	migrate(db)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Account{})
}

func createIfNotExists(name string) {
	os.Create(name)
}
