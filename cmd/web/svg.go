package web

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed "static/svg"
var SvgFiles embed.FS

type SvgRepo interface {
	ReadSVG(filename string, cssClasses ...string) string
	ReadSymbol(filename string, cssClasses ...string) string
}

type svgAssets struct {
	fs   embed.FS
	path string
}

// Read an ReadSVG file from the embedded assets.
// Return it as a string so it can be placed directly in the HTML.
// If the icon was not found it will return a circled exclamation point.
func (i svgAssets) ReadSVG(filename string, cssClasses ...string) string {
	classAttribute := ""
	if len(cssClasses) > 0 {
		classAttribute = fmt.Sprintf(`class="%s"`, strings.Join(cssClasses, " "))
	}

	filePath := fmt.Sprintf("%s/%s.svg", i.path, filename)
	file, error := i.fs.ReadFile(filePath)
	if error != nil {
		cssClasses = append(cssClasses, "animate-pulse")
		classAttribute = fmt.Sprintf(`class="%s"`, strings.Join(cssClasses, " "))

		// Red circled exclamation point
		return fmt.Sprintf(`<svg %s xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" aria-hidden="true" data-slot="icon" style="fill: red;">
		  <path fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12ZM12 8.25a.75.75 0 0 1 .75.75v3.75a.75.75 0 0 1-1.5 0V9a.75.75 0 0 1 .75-.75Zm0 8.25a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Z" clip-rule="evenodd"/>
	    </svg>`, classAttribute)
	}

	icon := string(file)
	icon = strings.Replace(icon, "<svg", fmt.Sprintf("<svg %s", classAttribute), 1)
	return icon
}

// Converts an SVG so it can be used as a symbol by changing the <svg> tag to a <symbol> tag and adding an id attribute.
// This is useful if the SVG will be used multiple times on the page. The data for the SVG only has to be transferred once.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/symbol
//
// # Example
//
//	<svg class="hidden">
//	  <symbol id="foo" ...>...</symbol>
//	</svg>
//	<svg>
//	  <use href="#foo"></use>
//	  <use href="#foo"></use>
//	</svg>
func (i svgAssets) ReadSymbol(filename string, cssClasses ...string) string {
	icon := i.ReadSVG(filename, cssClasses...)
	id := strings.ReplaceAll(filename, "/", "-")
	symbolTag := fmt.Sprintf(`<symbol id="%s"`, id)
	icon = strings.Replace(icon, "<svg", symbolTag, 1)
	icon = strings.Replace(icon, "</svg>", "</symbol>", 1)
	return icon
}

var Icons = SvgRepo(svgAssets{fs: SvgFiles, path: "static/svg"})
