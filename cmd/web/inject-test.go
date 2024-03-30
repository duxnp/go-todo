package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// -----------------------------------------------------
// -----------------------------------------------------
// -----------------------------------------------------
// this is located in "users" package. So you would group code by domain
// package users

type db interface { // if you use interface in this package you can easily test handler with mocked db
	GetUsernameByIDs(ctx echo.Context, ID int) (string, error)
}

type users struct {
	db db
}

func RegisterRoutes(e *echo.Echo, db db) {
	u := users{
		db: db,
	}

	// g := parent.Group("/users")

	// register all routes on controller
	e.GET("/inject-test", u.username)
}

func (ctrl *users) username(c echo.Context) error {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "invalid ID")
	// }

	result, err := ctrl.db.GetUsernameByIDs(c, 1)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, result)
}
