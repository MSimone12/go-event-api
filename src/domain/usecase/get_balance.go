package usecase

import "go-event-api/src/data/datasource"

// getBalance is the function responsible for getting the balance from the account
// receives the datasource.DataSource object from the dependency injection
func getBalance(source datasource.DataSource) GetBalance {
	return func(id uint) (int, error) {
		account, getBalanceError := source.Account.GetAccountById(id)
		return account.Balance, getBalanceError
	}
}
