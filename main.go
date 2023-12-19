package main

import (
	"devlocator/openapi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct{}
type TestResponse struct {
	Message string
}

func (s Server) GetApiEvent(ctx echo.Context, params openapi.GetApiEventParams) error {
	return ctx.JSON(http.StatusOK, TestResponse{
		Message: "search events",
	})
}

func (s Server) GetApiMypage(ctx echo.Context, params openapi.GetApiMypageParams) error {
	return ctx.JSON(http.StatusOK, TestResponse{
		Message: "mypage",
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

func main() {
	e := echo.New()
	s := Server{}
	openapi.RegisterHandlers(e, s)
	e.Logger.Fatal(e.Start(":8000"))
}
