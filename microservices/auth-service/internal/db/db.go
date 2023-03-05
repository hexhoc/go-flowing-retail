package db

import (
	"github.com/hexhoc/auth-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed connect to db ", err)
	}

	db.AutoMigrate(&models.User{})

	return Handler{DB: db}
}
