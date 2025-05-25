package web

import "github.com/labstack/echo/v4"

func Register(e *echo.Echo) {
	e.GET("/people", getAllPeopleHandler)
}
