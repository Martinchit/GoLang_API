package model

import "github.com/first-api/view"

// CreateTodo query
func CreateTodo(name string, todo string) error {
	insertQ, err := con.Query("INSERT INTO todo (name, todo) Values($1, $2)", name, todo)
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}

// GetTodos Query
func GetTodos() ([]view.Todo, error) {
	rows, err := con.Query("SELECT * From todo")
	if err != nil {
		return nil, err
	}
	todos := []view.Todo{}
	for rows.Next() {
		data := view.Todo{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

// GetTodo Query
func GetTodo(name string) ([]view.Todo, error) {
	rows, err := con.Query("SELECT * From todo where name = $1", name)
	if err != nil {
		return nil, err
	}
	todos := []view.Todo{}
	for rows.Next() {
		data := view.Todo{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

// DeleteTodo Query
func DeleteTodo(name string) error {
	deleteQ, err := con.Query("DELETE FROM todo WHERE name = $1", name)
	defer deleteQ.Close()
	if err != nil {
		return err
	}
	return nil
}
