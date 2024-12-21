package entities

import (
	"database/sql"
	"time"
)

type Task struct {
	ID           int          `json:"id,omitempty"`
	AuthorName   string       `json:"author_name" validate:"required"`
	AssigneeName string       `json:"assignee_name,omitempty"`
	CreatedDate  time.Time    `json:"created_date"`
	ResolvedDate sql.NullTime `json:"resolved_date,omitempty"`
	Status       string       `json:"status" validate:"required,oneof=new in_progress done"`
}
