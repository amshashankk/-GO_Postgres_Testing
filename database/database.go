package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Conntect() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=shashank sslmode=disable password=mobcoder@123")

	if err != nil {
		panic("failed to connect database")
	} else {

		fmt.Println("Successfully connected to the database")
	}

	defer db.Close()

}
