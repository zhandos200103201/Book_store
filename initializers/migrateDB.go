package initializers

import (
	"Test2/pkg/models"
	"log"
)

func SyncDB() {
	err := db.AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		log.Fatal("Error during migration of User model")
	}
}