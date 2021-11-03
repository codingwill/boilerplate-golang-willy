package interfaces

import "projectstructuring/models"

type ServicePassword interface {
	ServiceGetPassword() (result models.Response)
}
