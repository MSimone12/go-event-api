package usecase

import (
	"go-event-api/src/data/datasource"
	"go-event-api/src/domain/entity"
)

func Transfer(origin, destination, amount uint) (entity.Transfer, error) {
	return datasource.Transfer(origin, destination, amount)
}
