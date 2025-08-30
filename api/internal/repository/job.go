package repository

import (
	"database/sql"
	"time"

	"github.com/joelewaldo/job-tracker/api/internal/models"
	"github.com/joelewaldo/job-tracker/api/pkg/logger"
)

type JobRepository struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) *JobRepository {
	return &JobRepository{db: db}
}

func (r *JobRepository) Create(job *models.Job) error {
	query := `
		INSERT INTO jobs (company, position, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at;
	`

	err := r.db.QueryRow(
		query,
		job.Company,
		job.Position,
		job.Description,
		job.Status,
		time.Now(),
		time.Now(),
	).Scan(&job.ID, &job.CreatedAt, &job.UpdatedAt)

	if err != nil {
		logger.Log.WithError(err).Error("failed to insert job")
		return err
	}

	return nil
}

func (r *JobRepository) GetByID(id int64) (*models.Job, error) {
	query := `
		SELECT id, company, position, description, status, created_at, updated_at
		FROM jobs
		WHERE id = $1
		LIMIT 1;
	`

	var job models.Job
	err := r.db.QueryRow(query, id).Scan(
		&job.ID,
		&job.Company,
		&job.Position,
		&job.Description,
		&job.Status,
		&job.CreatedAt,
		&job.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Log.WithError(err).Error("failed to get job by id")
		return nil, err
	}

	return &job, nil
}

func (r *JobRepository) GetAll() ([]models.Job, error) {
	query := `SELECT id, company, position, description, status, created_at, updated_at FROM jobs ORDER BY created_at DESC;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	jobs := []models.Job{}
	for rows.Next() {
		var job models.Job
		err := rows.Scan(&job.ID, &job.Company, &job.Position, &job.Description, &job.Status, &job.CreatedAt, &job.UpdatedAt)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (r *JobRepository) GetByStatus(status models.JobStatus) ([]models.Job, error) {
	rows, err := r.db.Query(
		`SELECT id, company, position, description, status, created_at, updated_at 
         FROM jobs 
         WHERE status = $1 
         ORDER BY created_at DESC`, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	jobs := []models.Job{}
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.Company, &job.Position, &job.Description, &job.Status, &job.CreatedAt, &job.UpdatedAt); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (r *JobRepository) UpdateStatus(id int64, status models.JobStatus) error {
	query := `
		UPDATE jobs
		SET status = $1,
		    updated_at = $2
		WHERE id = $3;
	`

	_, err := r.db.Exec(query, status, time.Now(), id)
	return err
}
func (r *JobRepository) Delete(id int64) error {
	query := `DELETE FROM jobs WHERE id = $1;`
	_, err := r.db.Exec(query, id)
	return err
}
