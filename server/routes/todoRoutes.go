package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/tejasp2003/go-react-todo/controllers"
	"github.com/tejasp2003/go-react-todo/middleware"
)

func TodoRoutes(app *fiber.App) {
    api := app.Group("/api")
    api.Use(middleware.JWTMiddleware)
    api.Get("/todos", controllers.GetTodos)
    api.Get("/todos/:todoId", controllers.GetTodoById)
    api.Post("/todos", controllers.CreateTodo)
    api.Put("/todos/:todoId", controllers.UpdateTodo)
    api.Put("/todos/:todoId/complete", controllers.MarkTodoAsCompleted)
    api.Delete("/todos/:todoId", controllers.DeleteTodo)
}
