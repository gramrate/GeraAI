package database

import (
	"gera-ai/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dbURL := config.Config.DBConnectionString
	return gorm.Open(postgres.Open(dbURL), &gorm.Config{})
}
