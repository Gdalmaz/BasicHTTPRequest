package main

import (
	"post/database"
	"post/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routers.RecipeRouters(app)
	app.Listen(":9091")
}
