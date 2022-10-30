package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/learn-fiber/app/controllers"
	"github.com/koybigino/learn-fiber/app/middleware"
)

func HandleVoteRoute(app *fiber.App) {
	app.Post("/votes", middleware.AuthRequired(), controllers.CreateVote)
}
