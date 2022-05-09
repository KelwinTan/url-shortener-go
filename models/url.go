package models

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Url      string `gorm:"unique" json:"url"`
	ShortUrl string `gorm:"unique" json:"short_url"`
}
