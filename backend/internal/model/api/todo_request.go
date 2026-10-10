package api

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
