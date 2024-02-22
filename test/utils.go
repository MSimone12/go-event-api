package test

import (
	"go-event-api/src/data/datasource"
	"go-event-api/src/db"
	"go-event-api/src/domain/usecase"
	"os"
)

const dbName = "test.sqlite"

type TestUtils struct {
	TestDatabase   db.Database
	TestDataSource datasource.DataSource
	TestUseCase    usecase.UseCase
}

func SetupTestEnv() TestUtils {
	db.Reset(dbName)
	conn := db.Connect(dbName)
	TestDatabase := db.NewDatabase(conn)
	TestDataSource := datasource.NewDataSource(TestDatabase)
	TestUseCase := usecase.NewUseCase(TestDataSource)

	return TestUtils{
		TestDatabase:   TestDatabase,
		TestDataSource: TestDataSource,
		TestUseCase:    TestUseCase,
	}
}

func CleanUpTestEnv() {
	os.Remove(dbName)
}
