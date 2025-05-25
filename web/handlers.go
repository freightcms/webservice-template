package web

import (
	"net/http"

	"github.com/freightcms/webservice-template/db"
	"github.com/labstack/echo/v4"
)

func getAllPeopleHandler(c echo.Context) error {
	var r GetAllPeopleRequest
	if err := c.Bind(&r); err != nil {
		return err
	}
	q := db.NewQuery().SetPage(r.Page).SetPageSize(r.Limit)
	people, err := c.(AppContext).PersonResourceManager.Get(q)
	if err != nil {
		return err
	}
	res := GetAllPeopleResponse{
		Total:  -1,
		People: people,
	}
	c.JSON(http.StatusOK, res)
	return nil
}
