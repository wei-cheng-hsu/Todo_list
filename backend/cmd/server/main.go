package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"todo-backend/config"
	"todo-backend/internal/handler"
	"todo-backend/internal/repository"
	"todo-backend/internal/service"
)

func main() {

	//載入設定
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	//建立gorm db
	db, err := config.NewDatabase(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 取得 GORM 底層 sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection:", err)
	}
	defer sqlDB.Close()

	log.Println("Database connected successfully")

	//建立repository
	todoRepository := repository.NewTodoRepository(db)

	//建立service
	todoService := service.NewTodoService(todoRepository)

	//建立handler
	todoHandler := handler.NewTodoHandler(todoService)

	//建立gin router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
		},
	}))

	router.GET("/health", func(c *gin.Context) {
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":   "error",
				"database": "disconnected",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"database": "connected",
		})
	})

	//各個api route
	api := router.Group("/api")
	{
		api.POST("/todos", todoHandler.CreateTodo)
		api.GET("/todos", todoHandler.GetTodo)
		api.PUT("/todos/:id", todoHandler.UpdateTodo)
		api.PATCH("/todos/:id/finish", todoHandler.FinishTodo)
		api.DELETE("/todos/:id", todoHandler.DeleteTodo)
	}

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
