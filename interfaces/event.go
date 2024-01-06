package interfaces

import (
	"devlocator/models"
	"devlocator/openapi"
)

type EventRepositoryInterface interface {
	GetEvents(params openapi.GetApiEventParams) ([]models.Event, int64, error)
	GetDetailEvent(eventId string) (models.Event, error)
}
