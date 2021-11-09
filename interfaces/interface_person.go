package interfaces

import "projectstructuring/models"

type IServicePerson interface {
	ServiceGetAll() (result []models.Person)
	ServiceGetById(id int) (result models.Person)
}
