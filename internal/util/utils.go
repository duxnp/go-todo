package util

import (
	"context"

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
	e := ctx.Value("echo")
	return e.(*echo.Echo).Reverse(name, params...)
	// if e != nil {
	// 	return e.(*echo.Echo).Reverse(name, params...)
	// }
	// return "You forgot the SetEchoInstance middleware!"
}

// GetCurrentPath ensures that we're always
// returning a string from `ctx`
func GetCurrentPath(ctx context.Context) string {
	currentPath := ctx.Value("currentPath")
	if currentPath != nil {
		return currentPath.(string)
	}
	return ""
}
