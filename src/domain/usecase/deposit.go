package usecase

import (
	"go-event-api/src/data/datasource"
	"go-event-api/src/domain/entity"
)

func Deposit(destination, amount uint) (entity.Deposit, error) {
	return datasource.Deposit(destination, amount)
}
