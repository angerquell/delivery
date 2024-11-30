package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"delivery/config"
	"delivery/database/postgres"
	"delivery/models"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключаемся к базе данных
	err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer postgres.DB.Close() // Закрываем соединение с базой при завершении программы

	// Создаем сервер Gin
	r := gin.Default()

	// Регистрируем маршрут
	r.GET("/ping", func(c *gin.Context) {
		// Получаем список продуктов из базы данных
		products := models.GetAllProducts()

		// Отправляем ответ в формате JSON
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"products": products,
		})
	})

	// Получаем порт из конфигурации
	port := cfg.ServerPort
	if port == "" {
		port = "4040" // Значение по умолчанию
	}

	// Запускаем сервер
	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)
}
