package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/nakamuranatalia/useful-tools-api/internal/controller"
)

func HandleRequest(c controller.Controller) {
	r := echo.New()
	r.POST("/tools", c.SaveTool)
	r.GET("/tools", c.FindTools)
	r.GET("tool/:uuid", c.FindToolByUuid)
	r.DELETE("tool/:uuid", c.DeleteToolByUuid)
	r.Logger.Fatal(r.Start(":3000"))
}
