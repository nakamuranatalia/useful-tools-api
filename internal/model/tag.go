package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Tools      []Tool `gorm:"many2many:tool_tag" json:"-"`
}
