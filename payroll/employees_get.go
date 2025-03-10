package payroll

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	EmployeesEndpoint = "/v1/{division}/payroll/Employees{id}"
)

// Employees endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=EmployeeEmployees

func (s *Service) EmployeesGet(requestParams *EmployeesGetParams, ctx context.Context) (*EmployeesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewEmployeesGetResponse()
	path := s.rest.SubPath(EmployeesEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// Process query parameters
	utils.AddQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewEmployeesGetResponse() *EmployeesGetResponse {
	return &EmployeesGetResponse{}
}

type EmployeesGetResponse struct {
	Results Employees `json:"results"`
	Next    edm.URL   `json:"__next"`
}

func (s *Service) NewEmployeesGetParams() *EmployeesGetParams {
	selectFields, _ := utils.Fields(&Employee{})
	expandFields := []string{}
	return &EmployeesGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type EmployeesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}
