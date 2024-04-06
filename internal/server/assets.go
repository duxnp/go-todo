package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	precompiler "github.com/parnic/go-assetprecompiler"
)

const COMPILED_ASSETS_FILE_PREFIX = "app-"
const COMPILED_ASSETS_ROUTE_NAME = "compiled-assets"

type compileResults = map[precompiler.FileType]*precompiler.CompileResult

func (s *Server) compiledAssetsHandler(c echo.Context) error {
	file := c.Param("file")

	if strings.HasSuffix(file, ".css") {
		return c.Blob(http.StatusOK, "text/css", s.assets[precompiler.CSS].Bytes)
	} else {
		return c.Blob(http.StatusOK, "text/css", s.assets[precompiler.CSS].Bytes)
	}
}

// This will compile the assets but instead of writing to disk it will store it in memory
// For cache busting
func precompileAssets() compileResults {
	assets, error := precompiler.Compile(precompiler.Config{
		Files: []string{
			// "cmd/web/static/css/foo.css",
			"cmd/web/static/css/output.css",
		},
		Minify:     false,
		FilePrefix: COMPILED_ASSETS_FILE_PREFIX,
	})
	if error != nil {
		log.Fatal(error)
	}
	return assets
}

func setAssetsMiddleware(assets compileResults) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("assets", assets)
			return next(c)
		}
	}
}
