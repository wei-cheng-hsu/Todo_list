package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"todo-backend/internal/handler"
	"todo-backend/internal/repository"
	"todo-backend/internal/service"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//mysql連線
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	//建立db物件
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	//確認是否順利連上db
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

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
			"OPTION",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
		},
	}))

	router.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":   "error",
				"database": "disconnected",
			})
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

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
