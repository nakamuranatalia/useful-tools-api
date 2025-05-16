package model

import (
	"github.com/google/uuid"
)

type Tool struct {
	Id          uint      `gorm:"primarykey" json:"id"`
	Uuid        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"uuid"`
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Tags        []Tag     `gorm:"many2many:tool_tag"`
}
