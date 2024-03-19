package server

import (
	"net/http"

	"todo/cmd/web"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// fileServer := http.FileServer(http.FS(web.Files))
	// e.GET("/js/*", echo.WrapHandler(fileServer))

	// e.Static("/static/css", "cmd/web/static/css")
	// e.Static("/static/js", "cmd/web/static/js")

	e.Use(middleware.Static("cmd/web/static"))

	e.GET("/", echo.WrapHandler(templ.Handler(web.Home())))
	e.GET("/foo", echo.WrapHandler(templ.Handler(web.Home())))
	e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	// e.GET("/", s.HelloWorldHandler)

	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
