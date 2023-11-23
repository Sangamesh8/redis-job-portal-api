package database

import (
	"fmt"
	"job-portal-api/internal/models"

	"job-portal-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
	cfg.PostgresConfig.PostgresHost, cfg.PostgresConfig.PostgresUser, cfg.PostgresConfig.PostgresPassword, cfg.PostgresConfig.PostgresName,cfg.PostgresConfig.PostgresPort,cfg.PostgresConfig.PostgresSSLMode,cfg.PostgresConfig.PostgresTimezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err = db.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.Company{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.Jobs{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return nil, err
	}
	return db, nil
}
