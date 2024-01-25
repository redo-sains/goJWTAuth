package initializers

import (
	"jwtAuth/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("Failed to migrate database!!!")
	}
}
