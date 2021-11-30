package handler

import (
	"fmt"
	"projectstructuring/interfaces"
	"projectstructuring/models"
	"projectstructuring/services"

	"github.com/gofiber/fiber/v2"
)

func HandlerPostNewPerson(c *fiber.Ctx) error {

	var iPerson interfaces.IServicePerson = &services.Person{}

	arrayString := new(models.Person)
	if err := c.BodyParser(arrayString); err != nil {
		return c.SendString(fmt.Sprintf("There is an error occured. Error: %s", err))
	}

	personData := iPerson.ServiceCreateNew(arrayString)

	return c.JSON(personData)
}

func HandlerGetPersonById(c *fiber.Ctx) error {
	// variable declaration

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

	// defining interface and its type
	var iPerson interfaces.IServicePerson = &services.Person{}
	personData := iPerson.ServiceGetAll()

	// return JSON data
	return c.JSON(personData)
}

func HandlerDeletePersonById(c *fiber.Ctx) error {
	// variable declaration

	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println("Not an integer")
	}

	// defining interface and its type
	var iPerson interfaces.IServicePerson = &services.Person{}
	personData := iPerson.ServiceDeletePersonById(id)

	// return JSON data
	return c.JSON(personData)
}
