package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	// CreatedOn  int `json:"created_on"`
	// ModifiedOn int `json:"modified_on"`
	// DeletedOn  int `json:"deleted_on"`
}

var db *gorm.DB

func ConnectDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("taskTracker.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Task{})
}
