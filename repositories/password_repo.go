package repositories

import (
	"math/rand"
	"projectstructuring/models"
	"strings"
	"time"
)

// func RepoCar(c *fiber.Ctx) error {
// 	return c.SendString("Masuk repo")
// }

var isGenerated = false
var list []models.Password

func RepoGetPassword(id int) (result models.Password, err error) {
	if !isGenerated {
		list = GeneratePasswordList()
	}
	result = list[id]

	return result, nil
}

func GeneratePasswordList() (result []models.Password) {

	isGenerated = true

	var passwordItem models.Password
	for i := 0; i < 10; i++ {
		passwordItem.FirstCode = "This"
		passwordItem.SecondCode = "Is"
		passwordItem.SpecialCode = "A"
		passwordItem.RandomCode = GenerateString()
		passwordItem.CreatedAt = time.Now()
		passwordItem.UpdatedAt = time.Now()

		result = append(result, passwordItem)
	}
	return result
}

func GenerateString() (random string) {

	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length := 10

	var result strings.Builder

	for i := 0; i < length; i++ {
		randomNumber := rand.Intn(len(characters))
		character := characters[randomNumber]
		result.WriteString(string(character))
	}

	return result.String()

}
