package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB(dbname string) {
	_, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

// func CreateTask() {}
// func ReadTask() {}
// func UpdateTask() {}
// func DeleteTask() {}
