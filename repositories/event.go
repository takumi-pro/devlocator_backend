package repositories

import (
	"devlocator/models"
	"devlocator/openapi"
	"strings"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (repo *EventRepository) GetDetailEvent(eventId string) (models.Event, error) {
	var event models.Event
	result := repo.db.Model(&models.Event{}).Where("event_id = ?", eventId).First(&event)
	return event, result.Error
}

func (repo *EventRepository) GetEvents(params openapi.GetApiEventParams) ([]models.Event, int64, error) {
	var events []models.Event
	var count int64
	eventResponseFields := []string{"event_id", "title", "description", "event_url", "started_at", "ended_at", "limit", "accepted", "waiting", "updated_at", "place", "address", "lat", "lon"}

	query := repo.db.Model(&events)

	searchMethod := "or"
	if params.SearchMethod != nil && *params.SearchMethod == "and" {
		searchMethod = "and"
	}

	keywordCondition := repo.db
	if params.Keyword != nil {
		keywords := strings.Split(*params.Keyword, ",")
		if searchMethod == "or" {
			for _, keyword := range keywords {
				keywordCondition = keywordCondition.Or("title LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
			}
		}

		if searchMethod == "and" {
			for _, keyword := range keywords {
				keywordCondition = keywordCondition.Where("title LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
			}
		}
	}

	query = query.Where(keywordCondition)

	if params.Date != nil {
		dates := strings.Split(*params.Date, ",")
		query = query.Where("DATE(started_at) IN ?", dates)
	}

	if params.Prefecture != nil {
		query = query.Where("address LIKE ?", "%"+*params.Prefecture+"%")
	}

	if params.EventId != nil {
		query = query.Where("event_id = ?", *params.EventId)
	}

	query = query.
		Select(eventResponseFields).
		Order("started_at ASC").
		Find(&events).
		Count(&count)

	return events, count, nil
}
