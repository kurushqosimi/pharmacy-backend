package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_DSN")

	var db *gorm.DB
	var err error

	maxAttempts := 10
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("✅ Успешное подключение к базе данных!")
			DB = db
			return
		}

		log.Printf("⏳ Попытка %d/%d: база ещё не готова. Ошибка: %v\n", attempt, maxAttempts, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("❌ Не удалось подключиться к базе данных после %d попыток: %v", maxAttempts, err)
}
