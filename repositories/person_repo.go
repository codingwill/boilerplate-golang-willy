package repositories

import (
	"fmt"
	"projectstructuring/configs"
	"projectstructuring/models"
)

func RepoGetAll() (result []models.Person) {

	//vars initialization
	db := configs.DB

	err := db.Select(&result, "SELECT * FROM person")
	if err != nil {
		fmt.Println("Error:", err)
	}

	return result
}

func RepoGetById(id int) (isDataExist bool, result models.Person) {

	//vars initialization
	isDataExist = true
	db := configs.DB

	err := db.Get(&result, "SELECT * FROM person WHERE id = $1", id)
	if err != nil {
		fmt.Println("Error:", err)
		isDataExist = false
	}

	return isDataExist, result
}

func RepoCreatePerson(input *models.Person) (result int64) {

	//vars initialization
	db := configs.DB
	var id int64

	query := "INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3) returning id;"
	sqlResult := db.QueryRow(query, input.FirstName, input.LastName, input.Email)

	err := sqlResult.Scan(&id)
	if err != nil {
		fmt.Println("Error id:", err)
	}

	return id
}

func RepoDeleteById(id int) (isSuccess bool, err error) {
	//vars initialization
	isSuccess = false
	db := configs.DB

	query := "DELETE FROM person WHERE id = $1 RETURNING id;"
	result, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Error id:", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error id:", err)
	}

	if count > 0 {
		isSuccess = true
	}

	return isSuccess, err
}
