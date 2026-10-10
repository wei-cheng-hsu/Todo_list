package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoStatus string

const (
	TodoStatusUndo   TodoStatus = "undo"
	TodoStatusFinish TodoStatus = "finish"
	TodoStatusDelay  TodoStatus = "delay"
)

type Todo struct {
	ID          uuid.UUID  `json:"id" gorm:"type:char(36);primaryKey"`
	Title       string     `json:"title" gorm:"size:50;not null"`
	Description *string    `json:"description" gorm:"size:1000"`
	DueDate     time.Time  `json:"due_date" gorm:"type:date;not null"`
	SetupDate   time.Time  `json:"setup_date" gorm:"<-:false"`
	Status      TodoStatus `json:"status" gorm:"type:enum('undo','finish','delay');default:undo"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"<-:false"`
}

// 在新增 Todo 前自動產生 UUID
func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return nil
}
