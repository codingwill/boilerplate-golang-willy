package handler

import (
	//"projectstructuring/services"

	"encoding/json"
	"fmt"
	"projectstructuring/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandlerGetPassword(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	response := services.ServiceGetPassword(id)

	data, err := json.MarshalIndent(response, "", "	")
	if err != nil {
		c.SendString(fmt.Sprintf("There is an error occured. Error: %s", err))
	}

	return c.SendString(string(data))

}
