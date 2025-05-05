package service

import (
	"github.com/nakamuranatalia/useful-tools-api/internal/model"
	"github.com/nakamuranatalia/useful-tools-api/internal/repository"
)

type Service interface {
	SaveTool(model.Tool) (*model.Tool, error)
	FindTools() ([]model.Tool, error)
	FindToolByUuid(string) (*model.Tool, error)
}

type ToolsService struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) ToolsService {
	return ToolsService{
		repository: repository,
	}
}

func (s ToolsService) SaveTool(tool model.Tool) (*model.Tool, error) {
	return s.repository.SaveTool(tool)
}

func (s ToolsService) FindTools() ([]model.Tool, error) {
	return s.repository.FindTools()
}

func (s ToolsService) FindToolByUuid(toolUuid string) (*model.Tool, error) {
	return s.repository.FindToolByUuid(toolUuid)
}
