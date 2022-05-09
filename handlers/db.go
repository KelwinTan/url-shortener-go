package handlers

import "gorm.io/gorm"

type DBHandler struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) DBHandler {
	return DBHandler{DB}
}
