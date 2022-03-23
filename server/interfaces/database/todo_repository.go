package database

import (
	"github.com/react-go-app/backend/entities"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) Store(t entities.Todo) (int, error) {
	result, err := repo.Execute(
		"INSERT INTO todo (id, title) VALUES(?,?)", t.Id, t.Title,
	)

	if err != nil {
		return 0, err
	}
	id64, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	id := int(id64)

	return id, nil
}

func (repo *TodoRepository) FindById(identigier int) (entities.Todo, error) {
	row, err := repo.Query(
		"SELECT * FROM todo WHERE id = ?", identigier,
	)
	defer row.Close()

	t := entities.Todo{}
	if err != nil {
		return t, err
	}

	var id int
	var title string
	var description string
	var author string

	row.Next()
	if err = row.Scan(&id, &title, &description, &author); err != nil {
		return t, err
	}

	t.Id = id
	t.Title = title
	t.Description = description
	t.Author = author

	return t, nil
}

func (repo *TodoRepository) FindAll() (entities.Todoes, error) {
	rows, err := repo.Query("SELECT * FROM todo")
	defer rows.Close()

	todoes := entities.Todoes{}
	if err != nil {
		return todoes, err
	}
	for rows.Next() {
		var id int
		var title string
		var description string
		var author string

		if err := rows.Scan(&id, &title, &description, &author); err != nil {
			continue
		}
		todo := entities.Todo{
			Id:          id,
			Title:       title,
			Description: description,
			Author:      author,
		}
		todoes = append(todoes, todo)
	}
	return todoes, nil
}
