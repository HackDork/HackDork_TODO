package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Id       int
	TaskName string `json:"taskname" gorm:"text;not null;default:null"`
	Assignee string `json:"assignee" gorm:"text;not null;default:null"`
	IsDone   bool   `json:"isDone" gorm:"default:false"`
	UserID   uint
}
type TaskResponse struct {
	ID        uint
	TaskName  string `validate:"omitempty,ascii"`
	Assignee  string
	CreatedAt time.Time
	IsDone    bool `gorm:"default:false" json:"isDone"`
	UserID    uint
}
