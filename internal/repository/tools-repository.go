package repository

import (
	"github.com/nakamuranatalia/useful-tools-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	SaveTool(tool model.Tool) (*model.Tool, error)
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
