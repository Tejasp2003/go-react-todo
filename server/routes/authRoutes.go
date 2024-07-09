package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/tejasp2003/go-react-todo/controllers"
)

func AuthRoutes(app *fiber.App) {
    app.Post("/register", controllers.Register)
    app.Post("/login", controllers.Login)
}
