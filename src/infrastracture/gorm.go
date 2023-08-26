package infrastracture

import (
	"bm/src/entities"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB = NewGormDB("root", "secret",
	"localhost", "article")

func NewGormDB(username string, password string, host string, dbName string) *gorm.DB {
	// Define the database connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, username, password, dbName)

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&entities.User{}, &entities.Article{})
	return db
}
