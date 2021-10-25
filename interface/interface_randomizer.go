package interfaces

import "projectstructuring/models"

type Randomizer interface {
	ServiceGetRandomize() (result models.Response)
}
