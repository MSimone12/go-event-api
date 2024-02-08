package usecase

import (
	"go-event-api/src/data/datasource"
	"go-event-api/src/domain/entity"
)

func Withdraw(origin uint, amount uint) (entity.Withdraw, error) {
	return datasource.Withdraw(origin, amount)
}
