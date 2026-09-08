package models

import (
	"context"
	"time"
)

// Runner is responsible for executing remote task actions.
type Runner struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	LastHeardAt time.Time `json:"last_heard_at"`
}

func (r *Runner) Register(ctx context.Context) error {
	r.CreatedAt = time.Now()

	panic("implement me")
}

type Job struct {
	ID         uint64     `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	StartedAt  *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Status     string     `json:"status"`

	Action string `json:"action"`
	Notes  string `json:"notes,omitempty"`
	Logs   string `json:"logs,omitempty"`
}

func (j *Job) Register(ctx context.Context) error {
	panic("implement me")
}

func (j *Job) AppendLogs(ctx context.Context, data string) error {
	j.Logs += data

	// TODO save the updated logs to the database
	panic("implement me")
}
