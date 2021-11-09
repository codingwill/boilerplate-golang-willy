package handler

import (
	"fmt"
	"projectstructuring/interfaces"
	"projectstructuring/services"

	"github.com/gofiber/fiber/v2"
)

func HandlerPostNewPerson(c *fiber.Ctx) error {
	//first := c.Params("first")
	//last := c.Params("last")
	//email := c.Params("email")
	return c.SendString("not yet implemented")
}

func HandlerGetPersonById(c *fiber.Ctx) error {
	// variable declaration
	//var response models.Response

	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println("Not an integer")
	}
	// defining interface and its type
	var iPerson interfaces.IServicePerson = &services.Person{}

	personData := iPerson.ServiceGetById(id)

	// return JSON data
	return c.JSON(personData)
}

func HandlerGetPersonList(c *fiber.Ctx) error {
	// variable declaration
	//var response models.Response

	// defining interface and its type
	var iPerson interfaces.IServicePerson = &services.Person{}

	personData := iPerson.ServiceGetAll()

	// return JSON data
	return c.JSON(personData)
}
