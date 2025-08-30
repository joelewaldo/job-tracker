package models

import "time"

type JobStatus string

const (
	StatusApplied      JobStatus = "applied"
	StatusInterviewing JobStatus = "interview"
	StatusOffer        JobStatus = "offer"
	StatusRejected     JobStatus = "rejected"
	StatusAccepted     JobStatus = "accepted"
	StatusArchived     JobStatus = "archived"
)

type Job struct {
	ID          int64     `json:"id" db:"id"`
	Company     string    `json:"company" db:"company"`
	Position    string    `json:"position" db:"position"`
	Description string    `json:"description" db:"description"`
	Status      JobStatus `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (s JobStatus) IsValid() bool {
	switch s {
	case StatusApplied, StatusInterviewing, StatusOffer, StatusRejected, StatusAccepted, StatusArchived:
		return true
	default:
		return false
	}
}
