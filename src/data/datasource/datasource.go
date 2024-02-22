package datasource

import (
	"go-event-api/src/data/model"
	"go-event-api/src/db"
)

type GetAccountById func(id uint) (*model.Account, error)
type SaveAccount func(account *model.Account) error

// AccountDataSource is the struct responsible for managing the AccountDataSource
type AccountDataSource struct {
	GetAccountById GetAccountById
	SaveAccount    SaveAccount
}

// DataSource is the struct responsible for managing the dependency injection
type DataSource struct {
	Account AccountDataSource
}

// NewDataSource is the factory that handles the dependency injection
func NewDataSource(database db.Database) DataSource {
	return DataSource{
		Account: AccountDataSource{
			GetAccountById: getAccountById(database),
			SaveAccount:    saveAccount(database),
		},
	}
}
