package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string `json:"name"`
	Tools []Tool `gorm:"many2many:tag_tools;"`
}
