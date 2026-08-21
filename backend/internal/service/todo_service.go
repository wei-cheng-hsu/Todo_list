package service

import (
	"time"

	"todo-backend/internal/model"
	"todo-backend/internal/repository"
)

type TodoService interface {
	Create(req model.CreateTodoRequest) (*model.Todo, error)
	GetAll() ([]model.Todo, error)
	Update(id uint64, req model.UpdateTodoRequest) error
	Finish(id uint64) error
	Delete(id uint64) error
}

// Service中持有一個Repository
type todoService struct {
	repository repository.TodoRepository
}

// 建立新TodoService
func NewTodoService(repository repository.TodoRepository) TodoService {
	return &todoService{
		repository: repository,
	}
}

// 判斷 todo status(如果due晚於today，就改成delay)
func determineTodoStatus(
	dueDate time.Time,
	currentStatus model.TodoStatus,
) model.TodoStatus {
	if currentStatus == model.TodoStatusFinish {
		return model.TodoStatusFinish
	}

	now := time.Now()
	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)
	due := time.Date(
		dueDate.Year(),
		dueDate.Month(),
		dueDate.Day(),
		0, 0, 0, 0,
		dueDate.Location(),
	)
	if due.Before(today) {
		return model.TodoStatusDelay
	}
	return model.TodoStatusUndo
}

// GetAll 取得所有table的資料
func (s *todoService) GetAll() ([]model.Todo, error) {
	todos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	//把todos中每筆todo交由determineTodoStatus判斷status狀態
	for i := range todos {
		newStatus := determineTodoStatus(
			todos[i].DueDate,
			todos[i].Status,
		)

		//透過repository的UpdateStatus更新db status欄位
		if newStatus != todos[i].Status {
			if err := s.repository.UpdateStatus(
				todos[i].ID,
				newStatus,
			); err != nil {
				return nil, err
			}
			todos[i].Status = newStatus
		}
	}

	return todos, nil
}

// Create 把新todo加入到資料庫中
// 把Create()此method掛在todoService上
// 回傳值: (*model.Todo, error)
func (s *todoService) Create(req model.CreateTodoRequest) (*model.Todo, error) {
	dueDate, err := time.Parse("2006-01-02", req.DueDate) //把前端傳來的日期字串轉換成time.Time格式(給Todo使用並寫入資料庫)
	if err != nil {
		return nil, err
	}

	status := determineTodoStatus(
		dueDate,
		model.TodoStatusUndo,
	)

	// 把傳來的request正式變成Todo model
	todo := &model.Todo{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     dueDate,
		Status:      status,
	}

	// 傳給Repository寫入MySQL
	if err := s.repository.Create(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// Update 更新資料庫內該id的資料
func (s *todoService) Update(id uint64, req model.UpdateTodoRequest) error {
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		return err
	}

	todo := &model.Todo{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     dueDate,
	}

	return s.repository.Update(id, todo)
}

// Finish 更新指定id的todo status為finish
func (s *todoService) Finish(id uint64) error {
	return s.repository.Finish(id)
}

// Delete 刪除指定id的todo row
func (s *todoService) Delete(id uint64) error {
	return s.repository.Delete(id)
}
