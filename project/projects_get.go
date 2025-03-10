package project

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	ProjectsEndpoint = "/v1/{division}/project/Projects{id}"
)

// Projects endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectProjects

func (s *Service) ProjectsGet(requestParams *ProjectsGetParams, ctx context.Context) (*ProjectsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewProjectsGetResponse()
	path := s.rest.SubPath(ProjectsEndpoint)

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

func (s *Service) NewProjectsGetResponse() *ProjectsGetResponse {
	return &ProjectsGetResponse{}
}

type ProjectsGetResponse struct {
	Results Projects `json:"results"`
	Next    edm.URL  `json:"__next"`
}

func (s *Service) NewProjectsGetParams() *ProjectsGetParams {
	selectFields, _ := utils.Fields(&Project{})
	expandFields := []string{}
	return &ProjectsGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type ProjectsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select    *odata.Select `schema:"$select,omitempty"`
	Expand    *odata.Expand `schema:"$expand,omitempty"`
	Filter    *odata.Filter `schema:"$filter,omitempty"`
	Top       *odata.Top    `schema:"$top,omitempty"`
	Skip      *odata.Skip   `schema:"$skip,omitempty"`
	SkipToken string        `schema:"$skiptoken,omitempty"`
}
