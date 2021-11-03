package handler

import (
	"fmt"
	"projectstructuring/interfaces"
	"projectstructuring/models"
	"projectstructuring/services"

	"github.com/gofiber/fiber/v2"
)

func HandlerGetRandomizer(c *fiber.Ctx) error {

	// variable declaration
	var response models.Response

	option := c.Params("option")
	dataType := c.Params("type")

	if dataType == "number" {
		arrayInt := models.RandomArrayInteger{[]int32{1, 2, 3, 4, 5}}
		response = services.ServiceGetRandomizeNumber(arrayInt, option)
	} else if dataType == "string" {
		arrayString := new(models.RandomArrayString)
		if err := c.BodyParser(arrayString); err != nil {
			c.SendString(fmt.Sprintf("There is an error occured. Error: %s", err))
		}

		// defining interface and its type
		var interfaceRandomizer interfaces.IServiceRandomizer = &models.RandomArrayString{arrayString.Data}

		randomizedList := interfaceRandomizer.Randomize()

		response = services.ServiceGetResponse(randomizedList)
	}

	// return JSON data
	return c.JSON(response)

}

func HandlerGetRandomizeTest(c *fiber.Ctx) error {

	// defining interface and its type
	var interfaceRandomizer interfaces.IServiceRandomizer = &models.RandomArrayString{[]string{"Mike", "Ujang", "Kyle", "Ecin", "Ilham", "Penny", "Sheldon"}}
	randomizedList := interfaceRandomizer.Randomize()

	fmt.Printf("in handler: %s\n", randomizedList)

	output := services.ServiceGetResponse(randomizedList)

	return c.JSON(output)
}
