package usecase

import (
	"fmt"
	"go-event-api/src/data/datasource"
	"go-event-api/src/domain/entity"
)

// transfer is the function responsible for taking balance from the origin and putting into the destination
// receives the datasource.DataSource object from the dependency injection
func transfer(source datasource.DataSource) Transfer {
	return func(origin, destination, amount uint) (entity.Transfer, error) {
		if origin == destination {
			return entity.Transfer{}, fmt.Errorf("CANNOT TRANSFER TO THE SAME PERSON")
		}
		withdrawFunction := withdraw(source)
		withdraw, withdrawError := withdrawFunction(origin, amount)
		if withdrawError != nil {
			return entity.Transfer{}, withdrawError
		}
		depositFunction := deposit(source)
		deposit := depositFunction(destination, amount)

		return entity.Transfer{
			Origin:      withdraw.Origin,
			Destination: deposit.Destination,
		}, nil
	}
}
