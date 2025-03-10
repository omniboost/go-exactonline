package project

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	HourTypesEndpoint = "/v1/{division}/read/project/HourTypes{id}"
)

// HourTypes endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectHourTypes

func (s *Service) HourTypesGet(requestParams *HourTypesGetParams, ctx context.Context) (*HourTypesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewHourTypesGetResponse()
	path := s.rest.SubPath(HourTypesEndpoint)

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

func (s *Service) NewHourTypesGetResponse() *HourTypesGetResponse {
	return &HourTypesGetResponse{}
}

type HourTypesGetResponse struct {
	Results HourTypes `json:"results"`
	Next    edm.URL   `json:"__next"`
}

func (s *Service) NewHourTypesGetParams() *HourTypesGetParams {
	selectFields, _ := utils.Fields(&HourType{})
	expandFields := []string{}
	return &HourTypesGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type HourTypesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}
