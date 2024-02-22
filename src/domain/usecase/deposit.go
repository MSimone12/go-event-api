package usecase

import (
	"go-event-api/src/data/datasource"
	"go-event-api/src/domain/entity"
)

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
