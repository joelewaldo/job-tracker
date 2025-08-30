package http

import (
	"database/sql"

	"github.com/fasthttp/router"
	"github.com/joelewaldo/job-tracker/api/internal/http/handlers"
	"github.com/joelewaldo/job-tracker/api/internal/repository"
	"github.com/valyala/fasthttp"
)

func NewRouter(conn *sql.DB) fasthttp.RequestHandler {
	r := router.New()

	r.Handle("GET", "/health", handlers.HealthHandler)

	jobRepo := repository.NewJobRepository(conn)
	jobHandler := handlers.NewJobHandler(jobRepo)

	r.Handle("GET", "/jobs", jobHandler.GetJobs)
	r.Handle("POST", "/jobs", jobHandler.CreateJob)
	r.Handle("GET", "/jobs/{id}", jobHandler.GetJobByID)
	r.PATCH("/jobs/{id}/status", jobHandler.UpdateJobStatus)
	r.DELETE("/jobs/{id}", jobHandler.DeleteJob)

	return LoggingMiddleware(r.Handler)
}
