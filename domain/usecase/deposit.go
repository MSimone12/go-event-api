package usecase

import (
	"go-event-api/data/datasource"
	"go-event-api/domain/entity"
)

func Deposit(destination, amount uint) (entity.Deposit, error) {
	return datasource.Deposit(destination, amount)
}
