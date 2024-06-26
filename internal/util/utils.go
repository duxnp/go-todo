package util

import (
	"context"
	"errors"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

// Wrapper around the Echo Reverse function that will first retreive the Echo instance from the context
func Reverse(ctx context.Context, name string, params ...interface{}) string {
	e, error := GetEcho(ctx)
	if error != nil {
		return error.Error()
	}
	return e.Reverse(name, params...)
}

// GetCurrentPath ensures that we're always
// returning a string from `ctx`
// TODO: refactor something like this so it's co-located with the SetCurrentPath middleware function
// TODO: also use constants instead of duplicating the keys all over the place
func GetCurrentPath(ctx context.Context) string {
	currentPath := ctx.Value("currentPath")
	if currentPath != nil {
		return currentPath.(string)
	}
	return ""
}

func GetEcho(ctx context.Context) (*echo.Echo, error) {
	e := ctx.Value("echo")
	// Uncomment this to generate a 500 internal server error
	// return e.(*echo.Echo).Reverse(name, params...)
	if e != nil {
		return e.(*echo.Echo), nil
	}
	return nil, errors.New("you forgot the SetEchoInstance middleware")
}
