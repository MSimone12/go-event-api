package datasource

import (
	"go-event-api/src/data/model"
	"go-event-api/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AccountDataSourceSuite struct {
	suite.Suite
	utils      test.TestUtils
	loadedDeps bool
}

func (suite *AccountDataSourceSuite) SetupTest() {
	suite.utils = test.SetupTestEnv()
	suite.loadedDeps = true
}

func (suite *AccountDataSourceSuite) TearDownTest() {
	test.CleanUpTestEnv()
	suite.utils = test.TestUtils{}
	suite.loadedDeps = false
}

func (suite *AccountDataSourceSuite) TestShouldNotGetAccount() {
	_, error := suite.utils.TestDataSource.Account.GetAccountById(1)
	assert.NotNil(suite.T(), error)
}
func (suite *AccountDataSourceSuite) TestShouldSaveNonExistentAccount() {
	account := model.Account{
		ID:      1,
		Balance: 10,
	}
	error := suite.utils.TestDataSource.Account.SaveAccount(&account)
	assert.Nil(suite.T(), error)
}
func (suite *AccountDataSourceSuite) TestShouldSaveExistentAccount() {
	account := model.Account{
		ID:      1,
		Balance: 10,
	}
	suite.utils.TestDataSource.Account.SaveAccount(&account)
	savedAccount, _ := suite.utils.TestDataSource.Account.GetAccountById(account.ID)
	assert.Equal(suite.T(), savedAccount.Balance, 10)
	account.Balance += 10
	error := suite.utils.TestDataSource.Account.SaveAccount(&account)
	assert.Nil(suite.T(), error)
	savedAccount, _ = suite.utils.TestDataSource.Account.GetAccountById(account.ID)
	assert.Equal(suite.T(), savedAccount.Balance, 20)
}

func (suite *AccountDataSourceSuite) TestShouldGetAccount() {
	account := model.Account{
		ID:      1,
		Balance: 10,
	}
	saveError := suite.utils.TestDataSource.Account.SaveAccount(&account)
	assert.Nil(suite.T(), saveError)
	savedAccount, error := suite.utils.TestDataSource.Account.GetAccountById(1)
	assert.Nil(suite.T(), error)
	assert.NotNil(suite.T(), savedAccount)
	assert.Equal(suite.T(), account.ID, savedAccount.ID)
	assert.Equal(suite.T(), account.Balance, savedAccount.Balance)
}

func TestAccountDataSourceSuite(t *testing.T) {
	suite.Run(t, new(AccountDataSourceSuite))
}
