package entity

import (
	"time"
)

const (
	STATUS_CREATED = iota + 1
	STATUS_OPEN
	STATUS_IN_PROGRESS
	STATUS_TESTING
	STATUS_DONE
)

type Task struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
	Name        string     `json:"name"`
	Objective   string     `json:"objective"`
	Description *string    `json:"description"`
	ParentTask  *Task      `json:"parent_task"`
	Initiator   User       `json:"initiator"`
	Assignees   []User     `json:"assignees"`
	Deadline    *time.Time `json:"deadline"`
	Severity    Severity   `json:"difficulty"`
	Status      Status     `json:"status"`
}
