package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/joelewaldo/job-tracker/api/internal/models"
	"github.com/joelewaldo/job-tracker/api/internal/repository"
	"github.com/joelewaldo/job-tracker/api/pkg/logger"
	"github.com/valyala/fasthttp"
)

type JobHandler struct {
	repo *repository.JobRepository
}

func NewJobHandler(repo *repository.JobRepository) *JobHandler {
	return &JobHandler{repo: repo}
}

// CreateJob godoc
// @Summary Create a new job
// @Description Create a job with company, position, description, and optional status
// @Tags jobs
// @Accept json
// @Produce json
// @Param job body models.Job true "Job information"
// @Success 201 {object} models.Job
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Failed to create job"
// @Router /jobs [post]
func (h *JobHandler) CreateJob(ctx *fasthttp.RequestCtx) {
	var job models.Job

	if err := json.Unmarshal(ctx.PostBody(), &job); err != nil {
		logger.Log.WithError(err).Error("failed to parse job payload")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Invalid JSON payload")
		return
	}

	if job.Company == "" || job.Position == "" {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Company and Position are required")
		return
	}

	if job.Status == "" {
		job.Status = models.StatusApplied
	}

	if err := h.repo.Create(&job); err != nil {
		logger.Log.WithError(err).Error("failed to create job")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString("Failed to create job")
		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetContentType("application/json")
	_ = json.NewEncoder(ctx).Encode(job)
}

// GetJobs godoc
// @Summary Get all jobs
// @Description Retrieve all jobs, optionally filtered by status
// @Tags jobs
// @Accept json
// @Produce json
// @Param status query string false "Filter by job status"
// @Success 200 {array} models.Job
// @Failure 500 {string} string "Failed to fetch jobs"
// @Router /jobs [get]
func (h *JobHandler) GetJobs(ctx *fasthttp.RequestCtx) {
	statusQuery := string(ctx.QueryArgs().Peek("status"))

	var jobs []models.Job
	var err error

	if statusQuery != "" {
		jobs, err = h.repo.GetByStatus(models.JobStatus(statusQuery))
	} else {
		jobs, err = h.repo.GetAll()
	}

	if err != nil {
		logger.Log.WithError(err).Error("failed to get jobs")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString("Failed to fetch jobs")
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	_ = json.NewEncoder(ctx).Encode(jobs)
}

// GetJobByID godoc
// @Summary Get a job by ID
// @Description Retrieve a single job by its ID
// @Tags jobs
// @Accept json
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {object} models.Job
// @Failure 400 {string} string "Invalid job ID"
// @Failure 404 {string} string "Job not found"
// @Failure 500 {string} string "Failed to fetch job"
// @Router /jobs/{id} [get]
func (h *JobHandler) GetJobByID(ctx *fasthttp.RequestCtx) {
	idStr := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Invalid job ID")
		return
	}

	job, err := h.repo.GetByID(id)
	if err != nil {
		logger.Log.WithError(err).Error("failed to fetch job")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString("Failed to fetch job")
		return
	}

	if job == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString("Job not found")
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	_ = json.NewEncoder(ctx).Encode(job)
}

// UpdateJobStatus godoc
// @Summary Update job status
// @Description Update the status of a job by ID
// @Tags jobs
// @Accept json
// @Produce json
// @Param id path int true "Job ID"
// @Param status body object{status=string} true "New status"
// @Success 200 {string} string "Job status updated"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Failed to update job status"
// @Router /jobs/{id}/status [patch]
func (h *JobHandler) UpdateJobStatus(ctx *fasthttp.RequestCtx) {
	idStr := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Invalid job ID")
		return
	}

	var payload struct {
		Status string `json:"status"`
	}

	if err := json.Unmarshal(ctx.PostBody(), &payload); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Invalid JSON payload")
		return
	}

	if payload.Status == "" {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Status is required")
		return
	}

	if !models.JobStatus(payload.Status).IsValid() {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Invalid job status")
		return
	}

	if err := h.repo.UpdateStatus(id, models.JobStatus(payload.Status)); err != nil {
		logger.Log.WithError(err).Error("failed to update job status")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString("Failed to update job status")
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString(fmt.Sprintf("Job %d status updated to %s", id, payload.Status))
}

// DeleteJob godoc
// @Summary Delete a job
// @Description Delete a job by ID
// @Tags jobs
// @Accept json
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {string} string "Job deleted"
// @Failure 400 {string} string "Invalid job ID"
// @Failure 500 {string} string "Failed to delete job"
// @Router /jobs/{id} [delete]
func (h *JobHandler) DeleteJob(ctx *fasthttp.RequestCtx) {
	idStr := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Invalid job ID")
		return
	}

	if err := h.repo.Delete(id); err != nil {
		logger.Log.WithError(err).Error("failed to delete job")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString("Failed to delete job")
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString(fmt.Sprintf("Job %d deleted", id))
}
