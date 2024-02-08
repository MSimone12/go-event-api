package usecase

import "go-event-api/data/datasource"

func GetBalance(id uint) (uint, error) {
	return datasource.GetBalance(id)
}
