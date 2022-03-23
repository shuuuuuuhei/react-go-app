package entities

type Todo struct {
	Id          int
	Title       string
	Description string
	Author      string
}

type Todoes []Todo