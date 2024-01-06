package handlers

import (
	"devlocator/database"
	"devlocator/interfaces"
	"devlocator/models"
	"devlocator/openapi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	EventRepository interfaces.EventRepositoryInterface
}

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

	events, count, err := s.EventRepository.GetEvents(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

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
	event, err := s.EventRepository.GetDetailEvent(eventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	type EventDetailResponse struct {
		Event models.Event `json:"event"`
	}
	var responseEvent = EventDetailResponse{Event: event}
	return ctx.JSON(http.StatusOK, responseEvent)
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
