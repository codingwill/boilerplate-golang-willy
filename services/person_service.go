package services

import (
	"fmt"
	"projectstructuring/configs"
	"projectstructuring/models"
)

type Person models.Person

func (person Person) ServiceGetAll() (result []models.Person) {
	// Fisher-Yates algo

	db := configs.DB

	fmt.Println("DB in Service: ", db)

	db.Select(&result, "SELECT * FROM person")

	return result
}

func (person Person) ServiceGetById(id int) (result models.Person) {
	// Fisher-Yates algo
	people := []models.Person{}

	db := configs.DB
	fmt.Println("DB in Service: ", db)

	db.Select(&people, "SELECT * FROM person")
	result = people[id]
	return result
}
