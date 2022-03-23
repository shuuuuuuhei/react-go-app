package usecases

import (
	"github.com/react-go-app/backend/entities"
)

type TodoRepository interface {
	Store(entities.Todo) (int, error)
	FindById(int) (entities.Todo, error)
	FindAll() (entities.Todoes, error)
}