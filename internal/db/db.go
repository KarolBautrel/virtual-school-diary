package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnector() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres dbname=postgres sslmode=disable password=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("Connected to DB")
	return db, nil
}