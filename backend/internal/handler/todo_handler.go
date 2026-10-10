package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"todo-backend/internal/apperror"
	"todo-backend/internal/model/api"
	"todo-backend/internal/response"
	"todo-backend/internal/service"
)

// 定義handler需要TodoService
type TodoHandler struct {
	service service.TodoService
}

// 建立新Handler
func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

// 統一解析URL中的Todo ID
func parseTodoID(c *gin.Context) (uuid.UUID, bool) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil || id == uuid.Nil {
		response.Error(c, apperror.NewWithCode(
			apperror.ErrInvalidID,
			err,
		))
		return uuid.Nil, false
	}
	return id, true
}

// 建立GetTodo作為取得所有Todo資料的handler
func (h *TodoHandler) GetTodo(c *gin.Context) {
	todos, err := h.service.GetAll()

	if err != nil {
		response.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, todos)
}

// 建立CreateTodo作為建立新todo的handler
// c *gin.Context代表HTTP Request的上下文，handler以c和HTTP溝通
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req api.CreateTodoRequest //用來承接Frontend的json資料

	// 讀取前端 JSON，並依照 binding tag 驗證
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.NewWithCode(
			apperror.ErrInvalidRequest,
			err,
		))
		return
	}

	// 呼叫 Service(Create)
	todo, err := h.service.Create(req)
	if err != nil {
		response.Error(c, err)
		return
	}

	// 新增成功
	c.JSON(http.StatusCreated, todo)
}

// 建立UpdateTodo作為更新特定id todo資料的handler
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, ok := parseTodoID(c)
	if !ok {
		return
	}

	var req api.UpdateTodoRequest //承接Frontend的json資料

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(
			c, apperror.NewWithCode(
				apperror.ErrInvalidRequest,
				err,
			))
		return
	}

	if err := h.service.Update(id, req); err != nil {
		response.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "todo update successfully",
	})
}

// 建立FinishTodo作為更改特定id todo資料的status為finish之handler
func (h *TodoHandler) FinishTodo(c *gin.Context) {
	id, ok := parseTodoID(c)
	if !ok {
		return
	}

	if err := h.service.Finish(id); err != nil {
		response.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "todo update successfully",
	})
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, ok := parseTodoID(c)
	if !ok {
		return
	}

	if err := h.service.Delete(id); err != nil {
		response.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "todo delete successfully",
	})
}
