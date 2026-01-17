package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"TodoApp-Gin/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env 파일을 찾을 수 없습니다.")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	DB.AutoMigrate(&models.Todo{})
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Todo 리스트 조회
	r.GET("/todos", func(c *gin.Context) {
		var todos []models.Todo
		if err := DB.Find(&todos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)
	})

	// Todo 생성
	r.POST("/todos", func(c *gin.Context) {
		var newTodo models.Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.Create(&newTodo)
		c.JSON(http.StatusCreated, newTodo)
	})

	// Todo 수정
	r.PATCH("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo models.Todo
		if err := DB.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		var updatedTodo models.Todo
		if err := c.ShouldBindJSON(&updatedTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		DB.Model(&todo).Updates(updatedTodo)
		c.JSON(http.StatusOK, todo)
	})

	// Todo 삭제
	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo models.Todo
		if err := DB.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		DB.Delete(&todo)
		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}