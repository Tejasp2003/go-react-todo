package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tejasp2003/go-react-todo/config"
	"github.com/tejasp2003/go-react-todo/routes"
)

func main() {
	config.ConnectDB()

	app := fiber.New()

	//cors
	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin, Authorization ",
        AllowOrigins:     "http://localhost:5173",
        AllowCredentials: true,
		
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

		
	

	routes.AuthRoutes(app)
    routes.TodoRoutes(app)

	setUpRoutes(app)

	
	log.Fatal(app.Listen(":8000"))

}
