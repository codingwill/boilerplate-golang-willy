package services

import (
	"projectstructuring/models"
	repo "projectstructuring/repositories"
)

func ServiceGetPassword(id int) (result models.Response) {

	// Getting data from repo
	data, err := repo.RepoGetPassword(id)

	// Error handling
	if err != nil {
		var response models.Response
		response.Status = 500
		response.Message = "failed to get password"
		response.Data = ""
		response.From = "system"
		return response
	}

	// storing data to a returnable var
	result.Status = 200
	result.From = "system"
	result.Message = "Fetching data success"
	result.Data = data

	return result
}
