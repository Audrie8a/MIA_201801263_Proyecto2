package main

import (
	"./database"

	"./routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New())

	routes.Setup(app)
	app.Listen(":4000")
}
