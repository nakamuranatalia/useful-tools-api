package service

import (
	"github.com/nakamuranatalia/useful-tools-api/internal/model"
	"github.com/nakamuranatalia/useful-tools-api/internal/repository"
)

type Service interface {
	SaveTool(model.Tool) (*model.Tool, error)
}

type toolsService struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) toolsService {
	return toolsService{
		repository: repository,
	}
}

func (s toolsService) SaveTool(tool model.Tool) (*model.Tool, error) {
	return s.repository.SaveTool(tool)
}
