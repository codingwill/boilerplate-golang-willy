package services

import (
	"projectstructuring/models"
	"projectstructuring/repositories"
)

type Person models.Person

func (person Person) ServiceGetAll() (result models.Response) {

	people := repositories.RepoGetAll()

	result = models.Response{
		Message: "Successfully retrieved people.",
		From:    "system",
		Status:  200,
		Data:    people,
	}

	return result
}

func (person Person) ServiceGetById(id int) (result models.Response) {

	isPersonExist, personData := repositories.RepoGetById(id)

	if isPersonExist {
		result = models.Response{
			Message: "Successfully retrieved person.",
			From:    "system",
			Status:  200,
			Data:    personData,
		}

	} else {
		result = models.Response{
			Message: "Failed to retrieved person. Reason: No data with such ID.",
			From:    "system",
			Status:  400,
			Data:    struct{}{},
		}
	}

	return result
}

func (person Person) ServiceCreateNew(input *models.Person) (result models.Response) {

	id := repositories.RepoCreatePerson(input)

	result = models.Response{
		Message: "Successfully created a person.",
		From:    "system",
		Status:  200,
		Data: struct {
			Id int64 `json:"id"`
		}{
			id,
		},
	}

	return result
}

func (person Person) ServiceDeletePersonById(id int) (result models.Response) {

	isSuccess, err := repositories.RepoDeleteById(id)

	if isSuccess {
		result = models.Response{
			Message: "Successfully deleted a person.",
			From:    "system",
			Status:  200,
			Data:    struct{}{},
		}

	} else {
		errorMessage := "data does not exist"
		if err != nil {
			errorMessage = err.Error()
		}
		result = models.Response{
			Message: "Failed to delete a person. Reason: " + errorMessage + ".",
			From:    "system",
			Status:  400,
			Data:    struct{}{},
		}

	}

	return result
}
