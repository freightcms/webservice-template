package web

import "github.com/labstack/echo/v4"

func getAllPeopleHandler(c echo.Context) error {
	var r GetAllPeopleRequest
	if err := c.Bind(&r); err != nil {
		return err
	}
	c.JSON()
	return nil
}
