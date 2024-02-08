package usecase

import "go-event-api/src/data/datasource"

func GetBalance(id uint) (uint, error) {
	return datasource.GetBalance(id)
}
