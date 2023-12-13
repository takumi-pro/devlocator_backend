// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// bookmark
	// (PUT /api/event/bookmark)
	PutApiEventBookmark(ctx echo.Context) error
	// search events
	// (GET /api/events)
	GetApiEvent(ctx echo.Context, params GetApiEventParams) error
	// get mypage
	// (GET /api/mypage)
	GetApiMypage(ctx echo.Context, params GetApiMypageParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PutApiEventBookmark converts echo context to params.
func (w *ServerInterfaceWrapper) PutApiEventBookmark(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutApiEventBookmark(ctx)
	return err
}

// GetApiEvent converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiEvent(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiEventParams
	// ------------- Optional query parameter "event_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "event_id", ctx.QueryParams(), &params.EventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter event_id: %s", err))
	}

	// ------------- Optional query parameter "keyword" -------------

	err = runtime.BindQueryParameter("form", true, false, "keyword", ctx.QueryParams(), &params.Keyword)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter keyword: %s", err))
	}

	// ------------- Optional query parameter "search_method" -------------

	err = runtime.BindQueryParameter("form", true, false, "search_method", ctx.QueryParams(), &params.SearchMethod)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search_method: %s", err))
	}

	// ------------- Optional query parameter "date" -------------

	err = runtime.BindQueryParameter("form", true, false, "date", ctx.QueryParams(), &params.Date)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter date: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApiEvent(ctx, params)
	return err
}

// GetApiMypage converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiMypage(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiMypageParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "Authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Authorization")]; found {
		var Authorization string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for Authorization, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, valueList[0], &Authorization)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Authorization: %s", err))
		}

		params.Authorization = &Authorization
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApiMypage(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.PUT(baseURL+"/api/event/bookmark", wrapper.PutApiEventBookmark)
	router.GET(baseURL+"/api/events", wrapper.GetApiEvent)
	router.GET(baseURL+"/api/mypage", wrapper.GetApiMypage)

}