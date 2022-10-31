package initializers

import (
	"log"

	"github.com/Puyodead1/fosscord-server-go/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.UserSettings{})
	DB.AutoMigrate(&models.Channel{})
	DB.AutoMigrate(&models.Guild{})
	DB.AutoMigrate(&models.ReadState{})
}
