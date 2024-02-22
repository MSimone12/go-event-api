package usecase

import (
	"go-event-api/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite
	utils      test.TestUtils
	loadedDeps bool
}

func (suite *UseCaseSuite) SetupTest() {
	suite.utils = test.SetupTestEnv()
	suite.loadedDeps = true
}

func (suite *UseCaseSuite) TearDownTest() {
	test.CleanUpTestEnv()
	suite.utils = test.TestUtils{}
	suite.loadedDeps = false
}

func (suite *UseCaseSuite) TestShouldGetBalanceForNonExistentAccount() {
	_, error := suite.utils.TestUseCase.GetBalance(1)
	assert.NotNil(suite.T(), error)
}

func (suite *UseCaseSuite) TestShouldCreateAccountWithInitialBalance() {
	deposit := suite.utils.TestUseCase.Deposit(100, 10)
	assert.NotNil(suite.T(), deposit)
	assert.Equal(suite.T(), deposit.Destination.Id, "100")
	assert.Equal(suite.T(), deposit.Destination.Balance, 10)
}

func (suite *UseCaseSuite) TestShouldDepositIntoExistingAccount() {
	suite.utils.TestUseCase.Deposit(100, 10)
	deposit := suite.utils.TestUseCase.Deposit(100, 10)
	assert.Equal(suite.T(), deposit.Destination.Id, "100")
	assert.Equal(suite.T(), deposit.Destination.Balance, 20)
}

func (suite *UseCaseSuite) TestShouldNotWithdrawFromNonExistentAccount() {
	_, err := suite.utils.TestUseCase.Withdraw(200, 10)
	assert.NotNil(suite.T(), err)
}
func (suite *UseCaseSuite) TestShouldWithdrawFromExistingAccount() {
	suite.utils.TestUseCase.Deposit(100, 20)
	withdraw, _ := suite.utils.TestUseCase.Withdraw(100, 5)
	assert.Equal(suite.T(), withdraw.Origin.Id, "100")
	assert.Equal(suite.T(), withdraw.Origin.Balance, 15)
}

func (suite *UseCaseSuite) TestShouldNotWithdrawFromExistingAccount() {
	suite.utils.TestUseCase.Deposit(100, 20)
	withdraw, _ := suite.utils.TestUseCase.Withdraw(100, 5)
	assert.Equal(suite.T(), withdraw.Origin.Id, "100")
	assert.Equal(suite.T(), withdraw.Origin.Balance, 15)
	_, balanceError := suite.utils.TestUseCase.Withdraw(100, 20)
	assert.NotNil(suite.T(), balanceError)
}

func (suite *UseCaseSuite) TestShouldNotTransferFromNonExistentOrigin() {
	_, err := suite.utils.TestUseCase.Transfer(100, 200, 10)
	assert.NotNilf(suite.T(), err, "%s", err)
}
func (suite *UseCaseSuite) TestShouldNotTransferToOrigin() {
	_, err := suite.utils.TestUseCase.Transfer(100, 100, 30)
	assert.NotNilf(suite.T(), err, "error message: %s", err)
}

func (suite *UseCaseSuite) TestShouldNotTransferFromExistingOrigin() {
	suite.utils.TestUseCase.Deposit(100, 20)
	_, err := suite.utils.TestUseCase.Transfer(100, 200, 30)
	assert.NotNilf(suite.T(), err, "error message: %s", err)
}

func (suite *UseCaseSuite) TestShouldTransferFromExistingOrigin() {
	suite.utils.TestUseCase.Deposit(100, 20)
	transfer, _ := suite.utils.TestUseCase.Transfer(100, 200, 10)
	assert.Equal(suite.T(), transfer.Origin.Id, "100")
	assert.Equal(suite.T(), transfer.Origin.Balance, 10)
	assert.Equal(suite.T(), transfer.Destination.Id, "200")
	assert.Equal(suite.T(), transfer.Destination.Balance, 10)
}

func TestUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
