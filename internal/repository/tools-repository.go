package repository

import (
	"github.com/nakamuranatalia/useful-tools-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	SaveTool(model.Tool) (*model.Tool, error)
	FindTools() ([]model.Tool, error)
	FindToolByUuid(string) (*model.Tool, error)
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
