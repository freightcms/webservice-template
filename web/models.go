package web

import "github.com/freightcms/webservice-template/models"

type (
	GetAllPeopleRequest struct {
		// Limit the numer of results
		Limit int `json:"limit" query:"limit"`
		// Page number of the query
		Page int `json:"page" query:"page"`
	}

	// GetAllPeopleResponse is provided as the JSON or XML bindable repsonse
	// to an HTTP Request
	GetAllPeopleResponse struct {
		// Total is the number of results that are in the query
		Total  int64            `json:"total" xml:"total"`
		People []*models.Person `json:"people" xml:"people"`
	}
)
