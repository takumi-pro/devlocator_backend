package handlers

import (
	"devlocator/database"
	"devlocator/models"
	"devlocator/openapi"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Server struct{}
type TestResponse struct {
	Message string
}

type QueryParams struct {
	EventId      string `query:"event_id" validate:"omitempty,number"`
	SearchMethod string `query:"search_method" validate:"omitempty,oneof=or and"`
	Date         string `query:"date" validate:"omitempty,datetime"`
}

// イベント検索
// GET /api/events
func (s Server) GetApiEvent(ctx echo.Context, params openapi.GetApiEventParams) error {
	queryParams := new(QueryParams)
	if err := ctx.Bind(queryParams); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(queryParams); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := database.DBConnectGorm()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var events = []models.Event{}
	var count int64
	eventResponseFields := []string{"event_id", "title", "event_url", "started_at", "ended_at", "limit", "accepted", "waiting", "updated_at", "place", "address", "lat", "lon"}

	query := db.Model(&events)

	searchMethod := "and"
	if params.SearchMethod != nil && *params.SearchMethod == "or" {
		searchMethod = "or"
	}

	if params.Keyword != nil {
		keywords := strings.Split(*params.Keyword, ",")
		if searchMethod == "and" {
			for _, keyword := range keywords {
				query = query.Where("title LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
			}
		}
	}

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

	return ctx.JSON(http.StatusOK, models.EventsResponse{
		ResultsReturned: count,
		Events:          events,
	})
}

func (s Server) PutApiEventBookmark(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, TestResponse{
		Message: "mypage",
	})
}

func (s Server) GetApiEventsEventId(ctx echo.Context, eventId string) error {
	return ctx.JSON(http.StatusOK, TestResponse{
		Message: "detail event",
	})
}

func (s Server) GetApiUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, TestResponse{
		Message: "users",
	})
}

func (s Server) DBConnect(ctx echo.Context) error {
	_, err := database.DBConnectGorm()
	var message = "database connected!"
	if err != nil {
		message = err.Error()
	}

	return ctx.JSON(http.StatusOK, TestResponse{
		Message: message,
	})
}
