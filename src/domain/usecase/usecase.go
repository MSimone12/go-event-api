package usecase

import (
	"go-event-api/src/data/datasource"
	"go-event-api/src/domain/entity"
)

type Deposit func(destination, amount uint) entity.Deposit
type Withdraw func(origin, amount uint) (entity.Withdraw, error)
type GetBalance func(id uint) (int, error)
type Transfer func(origin, destination, amount uint) (entity.Transfer, error)

type UseCase struct {
	Deposit    Deposit
	Withdraw   Withdraw
	GetBalance GetBalance
	Transfer   Transfer
}

func NewUseCase(source datasource.DataSource) UseCase {
	return UseCase{
		GetBalance: getBalance(source),
		Deposit:    deposit(source),
		Withdraw:   withdraw(source),
		Transfer:   transfer(source),
	}
}
