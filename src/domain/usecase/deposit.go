package usecase

import (
	"go-event-api/src/data/datasource"
	"go-event-api/src/domain/entity"
)

// deposit is the function responsible for adding balance to the account
// receives the datasource.DataSource object from the dependency injection
func deposit(source datasource.DataSource) Deposit {
	return func(destination, amount uint) entity.Deposit {
		destinationAccount, _ := source.Account.GetAccountById(destination)
		destinationAccount.Balance += int(amount)
		source.Account.SaveAccount(destinationAccount)
		return entity.Deposit{
			Destination: destinationAccount.ToDestination(),
		}
	}
}
