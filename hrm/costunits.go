package hrm

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	CostunitsEndpoint = "/v1/{division}/hrm/Costunits"
)

// Costunits endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMCostunits

func (s *Service) CostunitsGet(requestParams *CostunitsGetParams, ctx context.Context) (*CostunitsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewCostunitsGetResponse()
	path := s.rest.SubPath(CostunitsEndpoint)

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

func (s *Service) NewCostunitsGetResponse() *CostunitsGetResponse {
	return &CostunitsGetResponse{}
}

type CostunitsGetResponse struct {
	Results Costunits `json:"results"`
}

func (s *Service) NewCostunitsGetParams() *CostunitsGetParams {
	selectFields, _ := utils.Fields(&Costunit{})
	return &CostunitsGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type CostunitsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
