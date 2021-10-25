package handler

import (
	//"projectstructuring/services"

	"encoding/json"
	"fmt"
	"projectstructuring/models"
	"projectstructuring/services"

	"github.com/gofiber/fiber/v2"
)

func HandlerGetRandomizer(c *fiber.Ctx) error {

	// variable declaration
	var data []byte
	var err error
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
		response = services.ServiceGetRandomizeString(arrayString, option)
	}

	// construct in json with some formatting
	data, err = json.MarshalIndent(response, "", "	")

	if err != nil {
		c.SendString(fmt.Sprintf("There is an error occured. Error: %s", err))
	}

	return c.SendString(string(data))

}
