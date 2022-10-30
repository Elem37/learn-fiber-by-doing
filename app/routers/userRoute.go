package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/learn-fiber/app/controllers"
	"github.com/koybigino/learn-fiber/app/middleware"
)

func HandleUserRoute(app *fiber.App) {
	u := app.Group("/users")

	u.Post("/", controllers.CreateUser)

	u.Get("/:id<int>", middleware.AuthRequired(), controllers.GetUserByID)
}
