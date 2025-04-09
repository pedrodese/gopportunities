package config

import (
	"github.com/pedrodese/gopportunities/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	dsn := "host=localhost user=gopportunities password=gopass dbname=gopportunities port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("PostgreSQL opening error: %v", err)
		return nil, err
	}

	if err := db.AutoMigrate(&schemas.Opening{}); err != nil {
		logger.Errorf("PostgreSQL automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
