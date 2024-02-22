package datasource

import (
	"go-event-api/src/data/model"
	"go-event-api/src/db"
)

func getAccountById(database db.Database) GetAccountById {
	return func(id uint) (*model.Account, error) {
		account := &model.Account{ID: id}
		res := database.DB.First(account)
		return account, res.Error
	}
}

func saveAccount(database db.Database) SaveAccount {
	return func(account *model.Account) error {
		res := database.DB.Save(account)
		return res.Error

	}
}
