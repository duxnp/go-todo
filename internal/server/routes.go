package server

import (
	"fmt"
	"net/http"
	"strings"

	"todo/cmd/web"
	"todo/cmd/web/todo"
	mw "todo/internal/middleware"
	"todo/internal/util"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	precompiler "github.com/parnic/go-assetprecompiler"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(mw.ContextValue)
	e.Use(mw.SetCurrentPath)
	e.Use(mw.SetEchoInstance)
	e.Use(mw.SetAssets(s.assets))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/js/*", echo.WrapHandler(fileServer))

	// e.Static("/static/css", "cmd/web/static/css")
	// e.Static("/static/js", "cmd/web/static/js")
	e.Use(middleware.Static("cmd/web/static"))

	e.RouteNotFound("/*", echo.WrapHandler(templ.Handler(web.NotFound())))
	e.HTTPErrorHandler = customHTTPErrorHandler

	db := s.db.GetConnection()

	g := e.Group("")

	todo.RegisterRoutes(g, db)

	g.GET("/", echo.WrapHandler(templ.Handler(web.Home())))
	g.GET("/daisyui", echo.WrapHandler(templ.Handler(web.DaisyUI()))).Name = "daisyui"
	g.GET("/foo", echo.WrapHandler(templ.Handler(web.Home())))
	g.GET("/foo2", HomeHandler)
	g.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	g.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	// g.GET("/", s.HelloWorldHandler)

	g.GET("/assets/css/:file", s.CssAssetHandler).Name = "assets-css"

	g.GET("/health", s.healthHandler)

	return e
}

func (s *Server) CssAssetHandler(c echo.Context) error {
	file := c.Param("file")

	if strings.HasSuffix(file, ".css") {
		return c.Blob(http.StatusOK, "text/css", s.assets[precompiler.CSS].Bytes)
	} else {
		return c.Blob(http.StatusOK, "text/css", s.assets[precompiler.CSS].Bytes)
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	errorPage := fmt.Sprintf("./cmd/web/%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func HomeHandler(c echo.Context) error {
	return util.Render(c, http.StatusOK, web.Home2())
}
