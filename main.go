package main

import (
	"github.com/gin-gonic/gin"
	"pharmacy-backend/config"
	"pharmacy-backend/routes"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080") // Запуск сервера
}
