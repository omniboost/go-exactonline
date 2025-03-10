package project

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	TimeTransactionsEndpoint = "/v1/{division}/project/TimeTransactions{id}"
)

// TimeTransactions endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectTimeTransactions

func (s *Service) TimeTransactionsGet(requestParams *TimeTransactionsGetParams, ctx context.Context) (*TimeTransactionsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewTimeTransactionsGetResponse()
	path := s.rest.SubPath(TimeTransactionsEndpoint)

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

func (s *Service) NewTimeTransactionsGetResponse() *TimeTransactionsGetResponse {
	return &TimeTransactionsGetResponse{}
}

type TimeTransactionsGetResponse struct {
	Results TimeTransactions `json:"results"`
	Next    edm.URL          `json:"__next"`
}

func (s *Service) NewTimeTransactionsGetParams() *TimeTransactionsGetParams {
	selectFields, _ := utils.Fields(&TimeTransaction{})
	expandFields := []string{}
	return &TimeTransactionsGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type TimeTransactionsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}
