package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/koybigino/learn-fiber/app/controllers"
	"github.com/koybigino/learn-fiber/app/middleware"
)

func HandlePostRoute(app *fiber.App) {

	router := app.Group("/posts")

	// Get all posts
	router.Get("/", controllers.GetAllPosts)

	// Get a specific post
	router.Get("/:id<int>", controllers.GetPostByID)

	// Create a Post
	router.Post("", middleware.AuthRequired(), controllers.CreatePost)

	// Upadate a post
	router.Put("/:id<int>", middleware.AuthRequired(), controllers.UpdatePost)

	// Delete a post
	router.Delete("/:id<int>", middleware.AuthRequired(), controllers.DeletePost)
}
