package repository

import (
	"github.com/nakamuranatalia/useful-tools-api/internal/model"
	"gorm.io/gorm"
)

type toolsRepository struct {
	gorm *gorm.DB
}

func NewRepository(gorm *gorm.DB) Repository {
	return toolsRepository{
		gorm: gorm,
	}
}

type Repository interface {
	SaveTool(tool model.Tool) (*model.Tool, error)
}

func (r toolsRepository) SaveTool(tool model.Tool) (*model.Tool, error) {
	result := r.gorm.Create(&tool)
	return &tool, result.Error
}
