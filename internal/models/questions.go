package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Question struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Ask  string `json:"ask"`
}


var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Question{})
	DB = database
}
