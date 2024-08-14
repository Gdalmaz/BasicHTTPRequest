package routers

import (
	"post/controllers"

	"github.com/gofiber/fiber/v2"
)

func RecipeRouters(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	recipe := v1.Group("/recipe")
	recipe.Post("/add-recipe", controllers.AddTaste)
	recipe.Put("/update-recipe", controllers.UpdateTaste)
	recipe.Delete("/delete-recipe", controllers.DeleteTaste)
}
