package database

import (
	"fmt"
	"github.com/olzhas-b/hello_world_crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)
const (
	HOST="localhost"
	DBPORT="5436"
	USER="postgres"
	PASSWORD="postgres"
	NAME="postgres"
)
var db *gorm.DB
var err error

func InitDataBase() *gorm.DB {
	fmt.Println("00000000000000000000000000000000000000000000000000000000")
	postgresConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", HOST, PASSWORD, USER, NAME, DBPORT)
	db, err = gorm.Open(postgres.Open(postgresConfig), &gorm.Config{})

	if err != nil {
		log.Println("failed open database", err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("failed migration")
	}
	return db
}

func GetDataBase() *gorm.DB {
	if db == nil {
		db = InitDataBase()
	}
	return db
}
