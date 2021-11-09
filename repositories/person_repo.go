package repositories

import (
	"projectstructuring/configs"
	"projectstructuring/models"
)

func RepoGetById(id int) (result models.Person) {
	people := []models.Person{}

	db := configs.DB
	db.Select(&result, "SELECT * FROM person")
	result = people[id]
	return result
}
