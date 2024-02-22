package usecase

import (
	"fmt"
	"go-event-api/src/data/datasource"
	"go-event-api/src/domain/entity"
)

func withdraw(source datasource.DataSource) Withdraw {
	return func(origin, amount uint) (entity.Withdraw, error) {
		originAccount, errorOrigin := source.Account.GetAccountById(origin)
		if errorOrigin != nil {
			return entity.Withdraw{}, errorOrigin
		}
		if (originAccount.Balance - int(amount)) < 0 {
			return entity.Withdraw{}, fmt.Errorf("NOT ENOUGH BALANCE")
		}
		originAccount.Balance -= int(amount)

		errorSave := source.Account.SaveAccount(originAccount)
		if errorSave != nil {
			return entity.Withdraw{}, errorSave
		}
		return entity.Withdraw{
			Origin: originAccount.ToOrigin(),
		}, nil
	}
}
