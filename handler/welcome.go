package handler

import (
	//"projectstructuring/services"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandlerWelcome(c *fiber.Ctx) error {

	welcomeText := fmt.Sprintf("Entering [HandlerWelcome]. Hello, %s!", c.Params("name"))
	return c.SendString(welcomeText)
}
