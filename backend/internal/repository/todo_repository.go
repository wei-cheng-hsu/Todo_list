package repository

import (
	"errors"
	"todo-backend/internal/model"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

var ErrTodoNotFound = errors.New("todo not found")

// 定義存取資料須提供的功能
type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	Create(todo *model.Todo) error
	Update(id uuid.UUID, todo *model.Todo) error
	Finish(id uuid.UUID) error
	Delete(id uuid.UUID) error
	UpdateStatus(id uuid.UUID, status model.TodoStatus) error
}

// 實作TodoRepository
type todoRepository struct {
	db *gorm.DB
}

// 建立新TodoRepository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) UpdateStatus(
	id uuid.UUID,
	status model.TodoStatus,
) error {
	result := r.db.
		Model(&model.Todo{}).
		Where("id = ?", id.String()).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrTodoNotFound
	}

	return nil
}

// 取得所有Todo資料
func (r *todoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo

	result := r.db.
		Order("due_date ASC").
		Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

// 新增Todo (Create)
func (r *todoRepository) Create(todo *model.Todo) error {
	if err := r.db.Create(todo).Error; err != nil {
		return err
	}

	return r.db.First(todo, "id = ?", todo.ID.String()).Error
}

// 編輯Todo (Update)
func (r *todoRepository) Update(id uuid.UUID, todo *model.Todo) error {
	result := r.db.
		Model(&model.Todo{}).
		Where("id = ?", id.String()).
		Updates(map[string]interface{}{
			"title":       todo.Title,
			"description": todo.Description,
			"due_date":    todo.DueDate,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrTodoNotFound
	}

	return nil
}

// 把Todo status改成finish (Finish)
func (r *todoRepository) Finish(id uuid.UUID) error {
	result := r.db.
		Model(&model.Todo{}).
		Where("id = ?", id.String()).
		Update("status", model.TodoStatusFinish)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrTodoNotFound
	}

	return nil
}

// 刪除Todo (Delete)
func (r *todoRepository) Delete(id uuid.UUID) error {
	result := r.db.
		Where("id = ?", id.String()).
		Delete(&model.Todo{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrTodoNotFound
	}

	return nil
}
