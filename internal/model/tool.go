package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tool struct {
	gorm.Model
	Uuid        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"uuid"`
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Tags        []Tag     `gorm:"many2many:tool_tag"`
}
