package model

import "time"

type TodoStatus string

const (
	TodoStatusUndo   TodoStatus = "undo"
	TodoStatusFinish TodoStatus = "finish"
	TodoStatusDelay  TodoStatus = "delay"
)

type Todo struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	DueDate     time.Time  `json:"due_date"`
	SetupDate   time.Time  `json:"setup_date"`
	Status      TodoStatus `json:"status"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// POST /api/todos 使用
type CreateTodoRequest struct {
	Title       string  `json:"title" binding:"required,max=50"`
	Description *string `json:"description" binding:"omitempty,max=1000"`
	DueDate     string  `json:"due_date" binding:"required"`
}

// PUT /api/todos/:id 使用
type UpdateTodoRequest struct {
	Title       string  `json:"title" binding:"required,max=50"`
	Description *string `json:"description" binding:"omitempty,max=1000"`
	DueDate     string  `json:"due_date" binding:"required"`
}
