package controller

import (
	"go-event-api/src/data/datasource"
	"go-event-api/src/db"
	"go-event-api/src/domain/usecase"
)

type Controller struct {
	UseCase usecase.UseCase
}

func Setup() Controller {
	database := db.NewDatabase(db.Connect(db.DefaultDatabaseName))
	source := datasource.NewDataSource(database)
	useCase := usecase.NewUseCase(source)
	return Controller{
		UseCase: useCase,
	}
}

var controller Controller = Setup()
