package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
)

// extend echo.Context
type contextValue struct {
	echo.Context
}

// Get retrieves data from the context.
func (ctx contextValue) Get(key string) interface{} {
	// grab value from echo.Context
	val := ctx.Context.Get(key)

	// if it's not nil, return it
	if val != nil {
		return val
	}

	// otherwise, return Request.Context
	return ctx.Request().Context().Value(key)
}

// Save data to the context.
func (ctx contextValue) Set(key string, value interface{}) {
	// we're replacing the whole Request in echo.Context
	// with a copied request that has the updated context value
	ctx.SetRequest(
		ctx.Request().WithContext(
			context.WithValue(ctx.Request().Context(), key, value),
		),
	)
}

// Middleware which places the default echo context into a custom version of the context.
// Receiver functions can then be added to it.
//
// This should be the first function in the middleware pipeline.
func ContextValue(next echo.HandlerFunc) echo.HandlerFunc {
	// this is just an echo.HandlerFunc
	return func(c echo.Context) error {
		// instead of passing next(c) as you usually would,
		// you return it with the extended version
		return next(contextValue{c})
	}
}

// Middleware which adds the current URL path to the context with a key named "currentPath"
//
// It can be retrieved like so:
//
//	currentPath := ctx.Value("currentPath")
//	if currentPath != nil {
//		return currentPath.(string)
//	}
func SetCurrentPath(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("currentPath", c.Request().URL.Path)
		return next(c)
	}
}

// Middleware which adds the Echo instance to the context with a key named "echo"
//
// It can be retrieved like so:
//
//	e := ctx.Value("echo")
//	if e != nil {
//		return e.(*echo.Echo)
//	}
func SetEchoInstance(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("echo", c.Echo())
		return next(c)
	}
}
