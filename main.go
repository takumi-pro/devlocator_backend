package main

import (
	"devlocator/handlers"
	"devlocator/openapi"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func validateDateTime(fl validator.FieldLevel) bool {
	// 日付の形式を正規表現でチェック
	re := regexp.MustCompile(`^\d{4}(\d{2})(\d{2})$`)
	// カンマで分割
	dates := strings.Split(fl.Field().String(), ",")

	// 分割された各日付に対してチェック
	for _, date := range dates {
		if !re.MatchString(date) {
			return false
		}
	}

	return true
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	if err := e.Validator.(*CustomValidator).validator.RegisterValidation("datetime", validateDateTime); err != nil {
		e.Logger.Fatal(err)
	}
	s := handlers.Server{}
	openapi.RegisterHandlers(e, s)
	e.Logger.Fatal(e.Start(":" + port))
}
