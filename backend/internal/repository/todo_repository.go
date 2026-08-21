package repository

import (
	"database/sql"
	"errors"
	"todo-backend/internal/model"
)

var ErrTodoNotFound = errors.New("todo not found")

// 定義存取資料須提供的功能
type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	Create(todo *model.Todo) error
	Update(id uint64, todo *model.Todo) error
	Finish(id uint64) error
	Delete(id uint64) error
	UpdateStatus(id uint64, status model.TodoStatus) error
}

// 實作TodoRepository
type todoRepository struct {
	db *sql.DB
}

// 建立新TodoRepository
func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) UpdateStatus(
	id uint64,
	status model.TodoStatus,
) error {
	result, err := r.db.Exec(`
		UPDATE todos
		SET status = ?
		WHERE id = ?
	`,
		status,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrTodoNotFound
	}

	return err
}

// 取得所有Todo資料
func (r *todoRepository) GetAll() ([]model.Todo, error) {
	rows, err := r.db.Query(`
		SELECT
			id,
			title,
			description,
			due_date,
			setup_date,
			status,
			updated_at
		FROM todos
		ORDER BY due_date ASC
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []model.Todo{}

	for rows.Next() {
		var todo model.Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.DueDate,
			&todo.SetupDate,
			&todo.Status,
			&todo.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	// 檢查迭代資料的過程中是否發生錯誤
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

// 新增Todo (Create)
func (r *todoRepository) Create(todo *model.Todo) error {
	result, err := r.db.Exec(`
		INSERT INTO todos (
			title,
			description,
			due_date,
			status
		)
		VALUES (?, ?, ?, ?)
	`,
		todo.Title,
		todo.Description,
		todo.DueDate,
		todo.Status,
	)

	if err != nil {
		return err
	}

	// 取得 MySQL AUTO_INCREMENT 產生的 ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	todo.ID = uint64(id)

	return nil
}

// 編輯Todo (Update)
func (r *todoRepository) Update(id uint64, todo *model.Todo) error {
	result, err := r.db.Exec(`
		UPDATE todos
		SET
			title = ?,
			description = ?,
			due_date = ?
		WHERE id = ?
	`,
		todo.Title,
		todo.Description,
		todo.DueDate,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected() //RowsAffected會回傳有多少row的資料會受result所帶的更動影響
	if err != nil {
		return err
	}

	//如果rowsAffected的數量是0，代表該row根本不存在
	if rowsAffected == 0 {
		return ErrTodoNotFound
	}

	return err
}

// 把Todo status改成finish (Finish)
func (r *todoRepository) Finish(id uint64) error {
	result, err := r.db.Exec(`
		UPDATE todos
		SET status = ?
		WHERE id = ?
	`,
		model.TodoStatusFinish,
		id,
	)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrTodoNotFound
	}

	return err
}

// 刪除Todo (Delete)
func (r *todoRepository) Delete(id uint64) error {
	result, err := r.db.Exec(`
		DELETE FROM todos
		WHERE id = ?
	`,
		id,
	)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrTodoNotFound
	}

	return err
}
