package entity

import (
	"github.com/jinzhu/gorm"
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
	gorm.Model
	Name        string
	Objective   string
	Description *string
	ParentTask  *Task
	Initiator   User
	Assignees   []User
	Deadline    *time.Time
	Difficulty  uint
	Status      uint
}
