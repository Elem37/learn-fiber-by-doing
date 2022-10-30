package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	"github.com/koybigino/learn-fiber/app/routers"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	routers.HandleHomeRoute(app)
	routers.HandlePostRoute(app)
	routers.HandleUserRoute(app)
	routers.HandleAuthRoute(app)
	routers.HandleVoteRoute(app)

	log.Fatal(app.Listen(":3000"))
}
