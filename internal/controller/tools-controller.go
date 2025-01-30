package controller

import (
	"github.com/nakamuranatalia/useful-tools-api/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nakamuranatalia/useful-tools-api/internal/model"
)

type Controller interface {
	SaveTool(echo.Context) error
	DeleteTool(echo.Context) error
}

type ToolsController struct {
	service service.Service
}

func NewController(service service.Service) ToolsController {
	return ToolsController{
		service: service,
	}
}

func (c ToolsController) SaveTool(context echo.Context) error {
	tool := model.Tool{}

	err := context.Bind(&tool)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	result, err := c.service.SaveTool(tool)
	if err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, result)
}
