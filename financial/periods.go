package financial

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	FinancialPeriodsEndpoint = "/v1/{division}/financial/FinancialPeriods"
)

// Periods endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialPeriods

func (s *Service) PeriodsGet(requestParams *PeriodsGetParams, ctx context.Context) (*PeriodsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewPeriodsGetResponse()
	path := s.rest.SubPath(FinancialPeriodsEndpoint)

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

func (s *Service) NewPeriodsGetResponse() *PeriodsGetResponse {
	return &PeriodsGetResponse{}
}

type PeriodsGetResponse struct {
	Results Periods `json:"results"`
	Next    edm.URL `json:"__next"`
}

func (s *Service) NewPeriodsGetParams() *PeriodsGetParams {
	selectFields, _ := utils.Fields(&Journal{})
	return &PeriodsGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type PeriodsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}
