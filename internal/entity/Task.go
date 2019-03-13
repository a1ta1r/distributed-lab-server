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
	DbAwareEntity
	Objective    string     `json:"objective"`
	Description  *string    `json:"description"`
	ParentTask   *Task      `gorm:"foreignkey:ParentTaskId;"`
	ParentTaskId uint       `json:"parent_task_id"`
	Initiator    User       `gorm:"foreignkey:InitiatorId;"`
	InitiatorId  uint       `json:"initiator_id"`
	Assignees    []User     `gorm:"many2many:user_tasks;"`
	Deadline     *time.Time `json:"deadline"`
	Severity     Severity   `json:"difficulty"`
	Status       Status     `gorm:"foreignkey:StatusId;"`
	StatusId     uint       `json:"status_id"`
	Project      Project    `gorm:"foreignkey:ProjectId;"`
	ProjectId    uint       `json:"project_id"`
}
