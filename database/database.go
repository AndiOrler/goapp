package database

import (
	"fmt"

	"github.com/AndiOrler/goapp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dbURL := "postgres://postgres:ndee_postgres@localhost:5432/ndee_test"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
