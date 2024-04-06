package assets

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	precompiler "github.com/parnic/go-assetprecompiler"
)

const COMPILED_ASSETS_FILE_PREFIX = "app-"
const COMPILED_ASSETS_CSS_ROUTE_NAME = "compiled-assets-css"
const COMPILED_ASSETS_JS_ROUTE_NAME = "compiled-assets-js"

type compileResults = map[precompiler.FileType]*precompiler.CompileResult

// This will compile the assets for cache busting
// Could also be configured to write the files to disk
func compile() compileResults {
	assets, error := precompiler.Compile(precompiler.Config{
		Files: []string{
			// "cmd/web/static/css/foo.css",
			"cmd/web/static/css/output.css",
			"cmd/web/static/js/application.js",
		},
		Minify:     false,
		FilePrefix: COMPILED_ASSETS_FILE_PREFIX,
	})
	if error != nil {
		log.Fatal(error)
	}
	return assets
}

// Compile CSS and JS assets then register routes to serve the assets from memory
func RegisterFeature(e *echo.Echo) {
	assets := compile()

	if assets[precompiler.CSS] != nil {
		e.GET(
			fmt.Sprintf("/assets/%s.css", assets[precompiler.CSS].Hash),
			func(c echo.Context) error {
				return c.Blob(http.StatusOK, "text/css", assets[precompiler.CSS].Bytes)
			},
		).Name = COMPILED_ASSETS_CSS_ROUTE_NAME
	}

	if assets[precompiler.JS] != nil {
		e.GET(
			fmt.Sprintf("/assets/%s.js", assets[precompiler.JS].Hash),
			func(c echo.Context) error {
				return c.Blob(http.StatusOK, "text/javascript", assets[precompiler.JS].Bytes)
			},
		).Name = COMPILED_ASSETS_JS_ROUTE_NAME
	}
}
