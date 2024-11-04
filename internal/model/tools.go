package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Tool struct {
	gorm.Model
	Title       string
	Link        string
	Description string
	Tags        pq.StringArray `gorm:"type:text[]"`
}
