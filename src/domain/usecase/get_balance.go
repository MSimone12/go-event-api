package usecase

import "go-event-api/src/data/datasource"

func getBalance(source datasource.DataSource) GetBalance {
	return func(id uint) (int, error) {
		account, getBalanceError := source.Account.GetAccountById(id)
		return account.Balance, getBalanceError
	}
}
