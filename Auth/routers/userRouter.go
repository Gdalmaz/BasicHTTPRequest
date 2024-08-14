package routers

import (
	"auth/controllers"
	"auth/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")

	user.Post("/sign-up", controllers.SignUp)
	user.Post("/sign-in", controllers.SignIn)
	user.Put("/update-account", controllers.UpdatePassword)
	user.Post("/token-control", middleware.TokenControl(), controllers.TokenControlHandler)
}
