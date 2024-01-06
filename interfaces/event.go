package interfaces

import (
	"devlocator/models"
	"devlocator/openapi"
)

type EventRepository interface {
	GetEvents(params openapi.GetApiEventParams) ([]models.Event, int64, error)
}
