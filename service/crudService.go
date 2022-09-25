package service

import (
	"github.com/asherplotnik/golang-rest/model"
	"github.com/asherplotnik/golang-rest/repository"
)

func DeleteAllTodos() (int, error){
	deleteCount, err := repository.DeleteAllTodo()
	if err != nil {
		return 0, err
	}
	return deleteCount, nil
}


func AddOneTodo(newTodo model.Todo) (model.Todo, error) {

	_, err := repository.GetTodoById(newTodo.ID)
	if err == nil {
		return model.Todo{}, model.TodoError{Message: "this id already exist", Code: 409}
	}

	result, err := repository.InsertOneTodo(newTodo)
	if err != nil {
		return model.Todo{}, err
	}

	return result, nil
}

func GetAllTodos() ([]model.Todo, error) {
	result, err := repository.GetAllTodos()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteTodo(id string) (int, error) {
	result, err := repository.DeleteOneTodo(id)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetTodo(id string) (model.Todo, error) {
	result, err := repository.GetTodoById(id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func UpdateTodo(todo model.Todo) (model.Todo, error) {
	_, err := repository.GetTodoById(todo.ID)
	if err != nil {
		return model.Todo{}, model.TodoError{Message: "not found", Code: 404}
	}
	_ , err = repository.UpdateOneTodo(todo)
	
	if err != nil {
		return model.Todo{}, model.TodoError{Message: "error - failed to update"}
	}
	return todo, nil
}