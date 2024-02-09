package datasource

import (
	"errors"
	"fmt"
	"go-event-api/src/data/model"
	"go-event-api/src/db"
	"go-event-api/src/domain/entity"

	"gorm.io/gorm"
)

func GetBalance(id uint) (uint, error) {
	db := db.Connect()
	account := model.Account{}
	result := db.First(&account, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("Account not found")
	}
	return account.Balance, nil
}

func Deposit(destination, amount uint) (entity.Deposit, error) {
	db := db.Connect()
	dest := model.Account{
		ID: destination,
	}

	res := db.First(&dest)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		res = db.Create(&dest)
		if res.Error != nil {
			return entity.Deposit{}, res.Error
		}
	}

	dest.Balance += amount
	db.Save(&dest)

	return entity.Deposit{
		Destination: dest.ToDestination(),
	}, nil
}

func Withdraw(origin, amount uint) (entity.Withdraw, error) {
	db := db.Connect()
	model := model.Account{
		ID: origin,
	}

	res := db.First(&model)

	if res.Error != nil {
		return entity.Withdraw{}, res.Error
	}
	if (model.Balance - amount) < 0 {
		return entity.Withdraw{}, fmt.Errorf("Not enough money in account")
	}
	model.Balance -= amount

	db.Save(&model)

	return entity.Withdraw{
		Origin: model.ToOrigin(),
	}, nil
}

func Transfer(origin, destination, amount uint) (entity.Transfer, error) {
	if origin == destination {
		return entity.Transfer{}, fmt.Errorf("You can't transfer for yourself")
	}
	db := db.Connect()

	orgn, dest := model.Account{
		ID: origin,
	}, model.Account{
		ID: destination,
	}

	originRes := db.Find(&orgn)
	if errors.Is(originRes.Error, gorm.ErrRecordNotFound) {
		return entity.Transfer{}, fmt.Errorf("Origin account not found")
	}
	if orgn.Balance < amount {
		return entity.Transfer{}, fmt.Errorf("Not enough balance in origin account")
	}

	destRes := db.Find(&dest)
	if errors.Is(destRes.Error, gorm.ErrRecordNotFound) {
		return entity.Transfer{}, fmt.Errorf("Destination account not found")
	}

	orgn.Balance -= amount
	dest.Balance += amount

	db.Save(&orgn)
	db.Save(&dest)

	return entity.Transfer{
		Origin:      orgn.ToOrigin(),
		Destination: dest.ToDestination(),
	}, nil
}
