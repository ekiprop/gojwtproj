package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Setup() (*gorm.DB, error) {

	dbUrl := fmt.Sprint(os.Getenv("DATABASE_URL"))

	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	if err = db.AutoMigrate(&User{}); err != nil {
		log.Println(err)
	}

	if err = db.AutoMigrate(&Grocery{}); err != nil {
		log.Println(err)
	}

	return db, err
}
