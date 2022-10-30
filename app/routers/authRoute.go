package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/learn-fiber/app/controllers"
)

func HandleAuthRoute(app *fiber.App) {
	app.Post("/login", controllers.Login)
}
