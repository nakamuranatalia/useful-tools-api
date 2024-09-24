package routes

import "github.com/labstack/echo/v4"

func HandleRequest() {
	r := echo.New()
	r.Logger.Fatal(r.Start("3000"))
}
