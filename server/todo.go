package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID   int    `json:"_id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var todos = []*Todo{
	{ID: 1, Task: "Learn Go", Done: false},
	{ID: 2, Task: "Learn Fiber", Done: false},
	{ID: 3, Task: "Build a Fiber app", Done: false},
	{ID: 4, Task: "Build a RESTful API", Done: false},
	{ID: 5, Task: "Build a RESTful API with Fiber", Done: false},
}

func setUpRoutes(app *fiber.App) {
	app.Get("/todos", getTodos)
	app.Post("/todos", addTodo)
	app.Get("/todos/:id", getTodo)
	app.Put("/todos/:id", updateTodo)
	app.Delete("/todos/:id", deleteTodo) 
}

func getTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func getTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id")) // atoi = ascii to integer
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	for _, todo := range todos {
		if id == todo.ID {
			return c.JSON(todo)
		}

	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Todo not found",
	})

}

func addTodo(c *fiber.Ctx) error {

	type request struct {
		Task string `json:"task"`
	}

	var body request

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	todo := &Todo{
		ID:   len(todos) + 1,
		Task: body.Task,
		Done: false,
	}

	todos = append(todos, todo)

	return c.JSON(todo)

}

func updateTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params(("id")))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}


	// dont make compulsion to update all fields
	type request struct {
		Task string `json:"task"`
		Done bool   `json:"done"`
	}

	var body request

	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	for _, todo := range todos {
		if id == todo.ID {
			if body.Task != "" {
				todo.Task = body.Task
			}
			todo.Done = body.Done

			return c.JSON(todo)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Todo not found",
	})

}


func deleteTodo (c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params(("id")))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	for i, todo := range todos {
		if id == todo.ID {
			var deletedTodo Todo = *todo
			todos = append(todos[:i], todos[i+1:]...)
			return c.JSON(deletedTodo)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Todo not found",
	})




		

}