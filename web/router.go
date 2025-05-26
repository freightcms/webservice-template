package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) *echo.Router {
	r := echo.NewRouter(e)

	r.Add(http.MethodGet, "/people", getAllPeopleHandler)
	r.Add(http.MethodGet, "/", echo.HandlerFunc(func(c echo.Context) error {
		body := struct {
			Status string `json:"status" xml:"status"`
		}{
			Status: "Ok",
		}
		return c.JSONPretty(http.StatusOK, &body, "    ")
	}))
	return r
}
