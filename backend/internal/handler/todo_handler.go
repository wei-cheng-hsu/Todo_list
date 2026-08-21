package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"todo-backend/internal/model"
	"todo-backend/internal/repository"
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

// 建立GetTodo作為取得所有Todo資料的handler
func (h *TodoHandler) GetTodo(c *gin.Context) {
	todos, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// 建立CreateTodo作為建立新todo的handler
// c *gin.Context代表HTTP Request的上下文，handler以c和HTTP溝通
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req model.CreateTodoRequest //用來承接Frontend的json資料

	// 讀取前端 JSON，並依照 binding tag 驗證
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	// 呼叫 Service(Create)
	todo, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 新增成功
	c.JSON(http.StatusCreated, todo)
}

// 建立UpdateTodo作為更新特定id todo資料的handler
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	//取得api中id所帶的數字(string)，並以strconv.ParseUint轉成uint64(10進位)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid todo id",
		})
		return
	}

	var req model.UpdateTodoRequest //承接Frontend的json資料

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if err := h.service.Update(id, req); err != nil {
		if errors.Is(err, repository.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "todo not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "todo update successfully",
	})
}

// 建立FinishTodo作為更改特定id todo資料的status為finish之handler
func (h *TodoHandler) FinishTodo(c *gin.Context) {
	//取得api中id所帶的數字(string)，並以strconv.ParseUint轉成uint64(10進位)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid todo id",
		})
		return
	}

	if err := h.service.Finish(id); err != nil {
		if errors.Is(err, repository.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "todo not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "todo update successfully",
	})
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	//取得api中id所帶的數字(string)，並以strconv.ParseUint轉成uint64(10進位)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid todo id",
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, repository.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "todo not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "todo delete successfully",
	})
}
