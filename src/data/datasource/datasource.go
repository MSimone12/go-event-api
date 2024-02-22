package datasource

import (
	"go-event-api/src/data/model"
	"go-event-api/src/db"
)

type GetAccountById func(id uint) (*model.Account, error)
type SaveAccount func(account *model.Account) error

type AccountDataSource struct {
	GetAccountById GetAccountById
	SaveAccount    SaveAccount
}

type DataSource struct {
	Account AccountDataSource
}

func NewDataSource(database db.Database) DataSource {
	return DataSource{
		Account: AccountDataSource{
			GetAccountById: getAccountById(database),
			SaveAccount:    saveAccount(database),
		},
	}
}
