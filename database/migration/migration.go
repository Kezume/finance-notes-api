package migration

import (
	"fmt"

	"github.com/Kezume/finance-notes/database"
	"github.com/Kezume/finance-notes/models"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&models.Transaction{})
	if err != nil {
		panic(err)
	}

	fmt.Println("success migrated")
}