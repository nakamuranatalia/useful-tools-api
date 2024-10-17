package model

import "gorm.io/gorm"

type Tool struct {
	gorm.Model
	Title       string
	Link        string
	Description string
	Tags        []string
}
