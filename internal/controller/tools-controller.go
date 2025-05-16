package controller

import (
	"github.com/nakamuranatalia/useful-tools-api/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nakamuranatalia/useful-tools-api/internal/model"
)

type Controller interface {
	SaveTool(echo.Context) error
	FindTools(echo.Context) error
	FindToolByUuid(echo.Context) error
	DeleteToolByUuid(echo.Context) error
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

func (c ToolsController) FindTools(context echo.Context) error {
	result, err := c.service.FindTools()
	if err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, result)
}

func (c ToolsController) FindToolByUuid(context echo.Context) error {
	uuid := context.Param("uuid")

	if uuid == "" {
		return context.String(http.StatusBadRequest, "Uuid is required")
	}

	result, err := c.service.FindToolByUuid(uuid)
	if err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, result)
}

func (c ToolsController) DeleteToolByUuid(context echo.Context) error {
	uuid := context.Param("uuid")

	if uuid == "" {
		return context.String(http.StatusBadRequest, "Uuid is required")
	}

	err := c.service.DeleteToolByUuid(uuid)
	if err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}

	return context.NoContent(http.StatusOK)
}
