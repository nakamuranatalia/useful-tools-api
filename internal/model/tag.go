package model

type Tag struct {
	Id    uint   `gorm:"primarykey" json:"id"`
	Name  string `json:"name"`
	Tools []Tool `gorm:"many2many:tool_tag" json:"-"`
}
