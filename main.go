package main

import (
	"devlocator/openapi"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	e := echo.New()
	s := Server{}
	openapi.RegisterHandlers(e, s)
	e.Logger.Fatal(e.Start(":" + port))
}
