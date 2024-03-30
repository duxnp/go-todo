package server

import (
	"fmt"
	"net/http"

	"todo/cmd/web"
	"todo/cmd/web/todo"
	"todo/internal/model"
	"todo/internal/repository"
	"todo/internal/util"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DummyDB struct{}

func (db *DummyDB) GetUsernameByIDs(ctx echo.Context, ID int) (string, error) {
	return "adam", nil
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// fileServer := http.FileServer(http.FS(web.Files))
	// e.GET("/js/*", echo.WrapHandler(fileServer))

	// e.Static("/static/css", "cmd/web/static/css")
	// e.Static("/static/js", "cmd/web/static/js")

	e.Use(middleware.Static("cmd/web/static"))

	// e.RouteNotFound()

	db := s.db.GetConnection()
	// repos := repository.InitRepositories(db)
	dummyDb := &DummyDB{}
	web.RegisterRoutes(e, dummyDb)
	todo.RegisterRoutes(e, db)

	e.GET("/", echo.WrapHandler(templ.Handler(web.Home())))
	e.GET("/daisyui", echo.WrapHandler(templ.Handler(web.DaisyUI())))
	e.GET("/foo", echo.WrapHandler(templ.Handler(web.Home())))
	e.GET("/foo2", HomeHandler)
	e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	// e.GET("/", s.HelloWorldHandler)
	e.GET("/test", s.TestHandler)

	e.GET("/health", s.healthHandler)

	return e
}

func HomeHandler(c echo.Context) error {
	return util.Render(c, http.StatusOK, web.Home2())
}

// I made this as a way to see if the todo repo was working
// I might not use separate repos
func (s *Server) TestHandler(c echo.Context) error {
	db := s.db.GetConnection()
	repos := repository.InitRepositories(db)

	todo := model.Todo{
		Title:       "foo",
		Description: "foobar",
	}
	foo, _ := repos.TodoRepo.CreateTodo(todo)
	fmt.Println(foo)

	todo2 := repos.TodoRepo.GetTodo(1)
	fmt.Printf("%#v\n", todo2)

	resp := map[string]string{
		"message": "Test complete",
	}

	return c.JSON(http.StatusOK, resp)
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
