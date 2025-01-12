package web

import (
	"database/sql"
	"log/slog"
	"net/http"

	"gorr/internal/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// HTTPServer is a wrapper around echo.Echo with a logger and database queries
type HTTPServer struct {
	Server *echo.Echo
	Logger *slog.Logger

	Queries database.Queries
}

// NewHttpServer creates a new http server and sets up logging, routes and static files
func NewHttpServer(db *sql.DB) *HTTPServer {
	hs := &HTTPServer{
		Server:  echo.New(),
		Queries: *database.New(db),
		Logger:  slog.Default(),
	}

	hs.Server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano}: method=${method}, uri=${uri}, status=${status}\n",
	}))

	// setup routes
	hs.routes()

	// frontend app
	hs.Server.Static("/assets", "frontend/build/client/assets")

	return hs
}

// Error is a global error handler for the http server
func (hs *HTTPServer) Error(c echo.Context, err error) error {
	hs.Logger.Error("http handler failed", "error", err)
	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error":  "internal error",
		"status": "error",
	})
}
