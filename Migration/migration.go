package migration

import (
	"fmt"

	config "github.com/ankur12345678/shout/Config"
	models "github.com/ankur12345678/shout/Models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	log.Info("Connecting to DB......")
	config := config.LoadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=%s", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT, config.DB_SSL_MODE, config.DB_TIMEZONE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Error connecting DB...Exiting...")
	}
	db.AutoMigrate(&models.Comment{}, &models.Interaction{}, &models.Post{}, &models.User{})
	log.Info("Connected to DB!")
	return db
}
