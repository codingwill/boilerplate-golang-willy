package interfaces

import "projectstructuring/models"

type IServicePerson interface {
	ServiceGetAll() (result models.Response)
	ServiceGetById(id int) (result models.Response)
	ServiceCreateNew(input *models.Person) (result models.Response)
	ServiceDeletePersonById(id int) (result models.Response)
}
