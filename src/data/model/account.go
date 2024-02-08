package model

import (
	"fmt"
	"go-event-api/src/domain/entity"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID      uint `json:"id" gorm:"primaryKey"`
	Balance uint `json:"balance"`
}

func (a *Account) ToOrigin() entity.Origin {
	return entity.Origin{
		Id:      fmt.Sprintf("%d", a.ID),
		Balance: a.Balance,
	}
}

func (a *Account) ToDestination() entity.Destination {
	return entity.Destination{
		Id:      fmt.Sprintf("%d", a.ID),
		Balance: a.Balance,
	}
}
