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

	// simple password generator
	app.Get("/password/:id", handler.HandlerGetPassword)

	// simple randomizer
	app.Get("/randomizer/generate/:type/:option", handler.HandlerGetRandomizer)
	app.Get("/randomizer/test", handler.HandlerGetRandomizeTest)

	// insert data
	app.Post("/person/:first/:last/:email", handler.HandlerPostNewPerson)

	// get data
	app.Get("/person/:id", handler.HandlerGetPersonById)
	app.Get("/person", handler.HandlerGetPersonList)

	app.Listen(":3001")
}
