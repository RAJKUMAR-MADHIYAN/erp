package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitDB creates and migrates the database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:rajk@567@tcp(127.0.0.1:3306)/enterprise?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err // Return the error instead of panicking
	}

	// Enable  debug mode for GORM to log SQL queries
	db = db.Debug()

	// Auto-migrate the Company struct to create the table if it doesn't exist
	if err := db.AutoMigrate(&Company{}).Error; err != nil {
		return nil, err // Return the error if auto-migration fails
	}

	return db, nil
}

/*
unc InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=moeloet dbname=edan_golang sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	db.Debug().AutoMigrate(&Company{})
	return db, nil
}*/
