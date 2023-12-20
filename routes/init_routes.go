package routes

import "github.com/labstack/echo/v4"

func Init() (*echo.Group, *echo.Echo) {
	e := echo.New()
	apiv1 := e.Group("/api/v1")
	e.Debug = true
	return apiv1, e
}
