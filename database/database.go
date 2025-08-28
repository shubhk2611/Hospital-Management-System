package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hospital/models"
	"log"
)

var DB *gorm.DB

// InitDatabase initializes the database connection and performs auto-migration
func InitDatabase() error {
	var err error
	// DSN string for MySQL connection.
	dsn := "root:pointers@tcp(127.0.0.1:3306)/hospital?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to the MySQL database using GORM.
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	// Auto-migrate the database schema for both Doctor and Patient models.
	err = DB.AutoMigrate(&models.Doctor{}, &models.Patient{})
	if err != nil {
		log.Printf("Failed to migrate database schema: %v", err)
		return err
	}

	log.Println("Database connected and migrated successfully")
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
