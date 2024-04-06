package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	precompiler "github.com/parnic/go-assetprecompiler"

	"todo/internal/database"
)

type Server struct {
	port int

	assets map[precompiler.FileType]*precompiler.CompileResult
	db     database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		assets: precompileAssets(),
		db:     database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

// This will compile the assets but instead of writing to disk it will store it in memory
// For cache busting
func precompileAssets() map[precompiler.FileType]*precompiler.CompileResult {
	assets, error := precompiler.Compile(precompiler.Config{
		Files: []string{
			// "cmd/web/static/css/foo.css",
			"cmd/web/static/css/output.css",
		},
		Minify:     false,
		FilePrefix: "app-",
	})
	if error != nil {
		log.Fatal(error)
	}
	return assets
}
