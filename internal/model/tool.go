package model

import (
	"gorm.io/gorm"
)

type Tool struct {
	gorm.Model
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Tags        []Tag  `gorm:"many2many:tool_tags"`
}
