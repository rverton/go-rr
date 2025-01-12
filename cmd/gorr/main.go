package main

import (
	"context"
	"database/sql"
	"flag"
	"gorr/internal/database"
	"gorr/internal/web"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var dbName = flag.String("dbname", "db.sqlite", "database name")

func main() {
	flag.Parse()

	ctx := context.Background()

	db, err := sql.Open("sqlite3", *dbName)
	if err != nil {
		slog.Error("unable to open sqlite db", "error", err)
		os.Exit(-1)
	}

	if err := database.ConfigureSqlite(db); err != nil {
		slog.Error("unable to open sqlite db", "error", err)
		os.Exit(-1)
	}

	if _, err := db.ExecContext(ctx, database.Schema); err != nil {
		slog.Error("unable to exec schema", "error", err)
		os.Exit(-1)
	}

	hs := web.NewHttpServer(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("starting server", "url", "http://localhost:"+port)
	http.ListenAndServe(":"+port, hs.Server)
}
