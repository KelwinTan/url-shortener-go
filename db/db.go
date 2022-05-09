package db

import (
	log "github.com/sirupsen/logrus"

	"github.com/KelwinTan/url-shortener-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:@localhost:5432/db_url_shortener"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}

	db.AutoMigrate(&models.Url{})

	return db
}
