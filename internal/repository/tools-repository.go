package repository

import (
	"github.com/nakamuranatalia/useful-tools-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	SaveTool(model.Tool) (*model.Tool, error)
	FindTools() ([]model.Tool, error)
	FindToolByUuid(string) (*model.Tool, error)
	DeleteToolByUuid(string) error
}

type ToolsRepository struct {
	gorm *gorm.DB
}

func NewRepository(gorm *gorm.DB) ToolsRepository {
	return ToolsRepository{
		gorm: gorm,
	}
}

func (r ToolsRepository) SaveTool(tool model.Tool) (*model.Tool, error) {
	result := r.gorm.Create(&tool)
	return &tool, result.Error
}

func (r ToolsRepository) FindTools() ([]model.Tool, error) {
	var tools []model.Tool
	result := r.gorm.Preload("Tags").Find(&tools)

	return tools, result.Error
}

func (r ToolsRepository) FindToolByUuid(toolUuid string) (*model.Tool, error) {
	var tool model.Tool
	result := r.gorm.Preload("Tags").First(&tool, "uuid", toolUuid)

	return &tool, result.Error
}

func (r ToolsRepository) DeleteToolByUuid(toolUuid string) error {
	tool, err := r.FindToolByUuid(toolUuid)
	if err != nil {
		return err
	}

	if err = r.gorm.Model(&tool).Association("Tags").Clear(); err != nil {
		return err
	}

	if err = r.gorm.Delete(&tool).Error; err != nil {
		return err
	}

	return nil
}

/*
var user User
db.First(&user, userID)

var newProjects []Project
db.Where("id IN ?", []uint{2, 3}).Find(&newProjects)

// Replace user’s projects
db.Model(&user).Association("Projects").Replace(&newProjects)

*/
