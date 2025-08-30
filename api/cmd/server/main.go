package main

import (
	"github.com/joelewaldo/job-tracker/api/internal/config"
	"github.com/joelewaldo/job-tracker/api/internal/http"
	"github.com/joelewaldo/job-tracker/api/internal/repository"
	"github.com/joelewaldo/job-tracker/api/pkg/logger"
	"github.com/valyala/fasthttp"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg)
	logger.Log.Info("Logger initialized with level: ", cfg.LogLevel)

	db, err := repository.NewDB(cfg.DatabaseURL)

	r := http.NewRouter(db.Conn)

	if err != nil {
		logger.Log.Fatal("Cannot connect to DB")
	}
	defer db.Close()

	address := ":" + cfg.ServerPort
	logger.Log.WithField("address", address).Info("Server listening on address")

	if err := fasthttp.ListenAndServe(address, r); err != nil {
		logger.Log.WithError(err).Fatal("Failed to start server")
	}
}
