package project

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	ProjectRestrictionItemsEndpoint = "/v1/{division}/project/ProjectRestrictionItems"
)

// ProjectRestrictionItems endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectProjectRestrictionItems

func (s *Service) ProjectRestrictionItemsGet(requestParams *ProjectRestrictionItemsGetParams, ctx context.Context) (*ProjectRestrictionItemsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewProjectRestrictionItemsGetResponse()
	path := s.rest.SubPath(ProjectRestrictionItemsEndpoint)

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

func (s *Service) NewProjectRestrictionItemsGetResponse() *ProjectRestrictionItemsGetResponse {
	return &ProjectRestrictionItemsGetResponse{}
}

type ProjectRestrictionItemsGetResponse struct {
	Results ProjectRestrictionItems `json:"results"`
	Next    edm.URL                 `json:"__next"`
}

func (s *Service) NewProjectRestrictionItemsGetParams() *ProjectRestrictionItemsGetParams {
	selectFields, _ := utils.Fields(&ProjectRestrictionItem{})
	expandFields := []string{}
	return &ProjectRestrictionItemsGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type ProjectRestrictionItemsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}
