package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"school-be/pkg/models"
	"school-be/utils"
)

var DB *gorm.DB

func InitializeDB(config *utils.PostgresConfig) error {
	var err error
	fmt.Println("Initializing PostgreSQL database connection...")

	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASS,
		config.DB_NAME,
	)

	DB, err = gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
	if err != nil {
		fmt.Printf("failed to connect to PostgreSQL database, error : %v", err)
		return err
	}

	DB.AutoMigrate(&models.Student{})

	return nil
}

func AutoMigrate(models ...interface{}) {
	DB.AutoMigrate(models...)
}
