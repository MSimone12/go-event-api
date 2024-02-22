package usecase

import (
	"go-event-api/src/data/datasource"
	"go-event-api/src/domain/entity"
)

type Deposit func(destination, amount uint) entity.Deposit
type Withdraw func(origin, amount uint) (entity.Withdraw, error)
type GetBalance func(id uint) (int, error)
type Transfer func(origin, destination, amount uint) (entity.Transfer, error)

// UseCase is the struct for the dependency injection,
// redirecting to the right function, passing the needed dependency
type UseCase struct {
	Deposit    Deposit
	Withdraw   Withdraw
	GetBalance GetBalance
	Transfer   Transfer
}

// NewUseCase is the factory for the dependency injection,
// returning the UseCase struct containing the redirect functions
func NewUseCase(source datasource.DataSource) UseCase {
	return UseCase{
		GetBalance: getBalance(source),
		Deposit:    deposit(source),
		Withdraw:   withdraw(source),
		Transfer:   transfer(source),
	}
}
