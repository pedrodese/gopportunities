package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	var err error

	//Initiliaze Database
	db, err = InitializePostgres()

	if err != nil {
		return fmt.Errorf("Error initializing Postgres: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize Logger
	logger = NewLogger(p)
	return logger
}
