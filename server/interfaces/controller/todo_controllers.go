package controller

import (
	"github.com/react-go-app/backend/interfaces/database"
	"github.com/react-go-app/backend/usecases"
)

type TodoController struct {
	Interactor usecases.TodoInteractor
}

func CreateTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController {
		Interactor: usecases.TodoInteractor {
			TodoRepository: &database.TodoRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) Index(c Context) {
	
}