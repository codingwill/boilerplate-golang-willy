package routes

import (
	"fmt"
	"projectstructuring/handler"

	// "projectstructuring/repositories"
	// "projectstructuring/services"

	"github.com/gofiber/fiber/v2"
)

func Route() {
	fmt.Println("Accessing routes")

	app := fiber.New()

	app.Get("/welcome", handler.HandlerWelcome)
	app.Get("/welcome/:name", handler.HandlerWelcome)
	app.Get("/password/:id", handler.HandlerGetPassword)

	app.Listen(":3001")
}
