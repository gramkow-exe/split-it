package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func GetTodos(context *fiber.Ctx) error {
	return context.Status(200).JSON(todos)
}

func AddTodo(context *fiber.Ctx) error {
	var newTodo todo

	if err := context.BodyParser(&newTodo); err != nil {
		println("Error binding JSON:", err)
		return err
	}

	todos = append(todos, newTodo)

	return context.Status(200).JSON(newTodo)
}

func GetTodo(context *fiber.Ctx) error {
	id := context.Params("id")
	todo, err := GetTodoById(id)

	if err != nil {
		context.Status(404).JSON(fiber.Map{"message": "Todo not found"})
		return err
	}

	return context.Status(200).JSON(todo)
}

func GetTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func ToggleTodoStatus(context *fiber.Ctx) error {
	id := context.Params("id")
	todo, err := GetTodoById(id)

	if err != nil {
		context.Status(404).JSON(fiber.Map{"message": "Todo not found"})
		return err
	}

	todo.Completed = !todo.Completed

	return context.Status(200).JSON(todo)
}
