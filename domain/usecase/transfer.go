package usecase

import (
	"go-event-api/data/datasource"
	"go-event-api/domain/entity"
)

func Transfer(origin, destination, amount uint) (entity.Transfer, error) {
	return datasource.Transfer(origin, destination, amount)
}
