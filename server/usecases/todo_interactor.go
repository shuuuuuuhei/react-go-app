package usecases

import (
	"github.com/react-go-app/backend/entities"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Add(t entities.Todo) (error) {
	_, err := interactor.TodoRepository.Store(t)
	return err
}

func (interactor *TodoInteractor) Users() (todoes entities.Todoes, err error) {
	todoes, err = interactor.TodoRepository.FindAll()
	return todoes, err
}

func (interactor *TodoInteractor) TodoById(identifier int) (entities.Todo, error) {
	todo, err := interactor.TodoRepository.FindById(identifier)
	return todo, err
}