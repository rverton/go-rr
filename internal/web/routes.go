package web

import (
	"gorr/frontend"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (hs *HTTPServer) routes() {
	hs.Server.GET("/", hs.index)

	// API routes
	hs.Server.GET("/api/authors", hs.authorsList)
	hs.Server.POST("/api/authors", hs.authorsInsert)
}

func (hs *HTTPServer) index(c echo.Context) error {
	return c.HTML(http.StatusOK, frontend.IndexHTML)
}

func (hs *HTTPServer) authorsList(c echo.Context) error {
	authors, err := hs.Queries.AuthorsList(c.Request().Context())
	if err != nil {
		return hs.Error(c, err)
	}

	return c.JSON(http.StatusOK, authors)
}

func (hs *HTTPServer) authorsInsert(c echo.Context) error {
	name := c.FormValue("name")

	author, err := hs.Queries.AuthorsInsert(c.Request().Context(), name)
	if err != nil {
		return hs.Error(c, err)
	}

	return c.JSON(http.StatusOK, author)
}
