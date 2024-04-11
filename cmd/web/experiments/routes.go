package experiments

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type handlers struct{}

func RegisterRoutes(parent *echo.Group) {
	// h := handlers{}

	g := parent.Group("/experiments")

	g.GET("/autoanimate", echo.WrapHandler(templ.Handler(AutoAnimate()))).Name = "autoanimate"
	g.GET("/daisyui", echo.WrapHandler(templ.Handler(DaisyUI()))).Name = "daisyui"
	g.GET("/hyperscript", echo.WrapHandler(templ.Handler(Hyperscript()))).Name = "hyperscript"
	g.GET("/tailwind", echo.WrapHandler(templ.Handler(Tailwind()))).Name = "tailwind"
}
