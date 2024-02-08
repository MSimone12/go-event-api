package usecase

import (
	"go-event-api/data/datasource"
	"go-event-api/domain/entity"
)

func Withdraw(origin uint, amount uint) (entity.Withdraw, error) {
	return datasource.Withdraw(origin, amount)
}
