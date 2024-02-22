package datasource

import (
	"go-event-api/src/data/model"
	"go-event-api/src/db"
)

// getAccountById receives the db.Database object and
// redirects to a function that returns the account, if existent and the error, also if existent
func getAccountById(database db.Database) GetAccountById {
	return func(id uint) (*model.Account, error) {
		account := &model.Account{ID: id}
		res := database.DB.First(account)
		return account, res.Error
	}
}

// saveAccount receives a db.Database object and
// redirects to a function that will save the pointer from model.Account into the database
func saveAccount(database db.Database) SaveAccount {
	return func(account *model.Account) error {
		res := database.DB.Save(account)
		return res.Error

	}
}
