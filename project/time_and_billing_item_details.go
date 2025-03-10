package project

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	TimeAndBillingItemDetailsEndpoint = "/v1/{division}/read/project/TimeAndBillingItemDetails"
)

// TimeAndBillingItemDetails endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectTimeAndBillingItemDetails

func (s *Service) TimeAndBillingItemDetailsGet(requestParams *TimeAndBillingItemDetailsGetParams, ctx context.Context) (*TimeAndBillingItemDetailsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewTimeAndBillingItemDetailsGetResponse()
	path := s.rest.SubPath(TimeAndBillingItemDetailsEndpoint)

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

func (s *Service) NewTimeAndBillingItemDetailsGetResponse() *TimeAndBillingItemDetailsGetResponse {
	return &TimeAndBillingItemDetailsGetResponse{}
}

type TimeAndBillingItemDetailsGetResponse struct {
	Results TimeAndBillingItemDetails `json:"results"`
	Next    edm.URL                   `json:"__next"`
}

func (s *Service) NewTimeAndBillingItemDetailsGetParams() *TimeAndBillingItemDetailsGetParams {
	// selectFields, _ := utils.Fields(&TimeAndBillingItemDetail{})
	// expandFields := []string{}
	return &TimeAndBillingItemDetailsGetParams{
		// Select: odata.NewSelect(selectFields),
		// Expand: odata.NewExpand(expandFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type TimeAndBillingItemDetailsGetParams struct {
	// @TODO: check if this an OData struct or something
	// Select *odata.Select `schema:"$select,omitempty"`
	// Expand *odata.Expand `schema:"$expand,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}
