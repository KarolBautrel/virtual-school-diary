package db

import (
	"fmt"
	"virtual-diary/internal/class/classdao"
	"virtual-diary/internal/student/studentdao"
	envUtils "virtual-diary/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func syncDb(db *gorm.DB) error {
	err := db.AutoMigrate(&classdao.Class{}, &studentdao.Student{})
	if err != nil {
		return err
	}
	return nil
}
func DBConnector() (*gorm.DB, error) {

	host, user, dbName, password, sslMode := envUtils.GetDbEnvironmentVariables()
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", host, user, dbName, sslMode, password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("Connected to DB")
	if err := syncDb(db); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	return db, nil
}
