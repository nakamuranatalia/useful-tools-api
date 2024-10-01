package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleRequest() {
	r := echo.New()
	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	r.Logger.Fatal(r.Start(":3000"))
}
