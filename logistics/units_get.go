package logistics

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	UnitsEndpoint = "/v1/{division}/logistics/Units"
)

// Units endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsUnits

func (s *Service) UnitsGet(requestParams *UnitsGetParams, ctx context.Context) (*UnitsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewUnitsGetResponse()
	path := s.rest.SubPath(UnitsEndpoint)

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

func (s *Service) NewUnitsGetResponse() *UnitsGetResponse {
	return &UnitsGetResponse{}
}

type UnitsGetResponse struct {
	Results Units `json:"results"`
}

func (s *Service) NewUnitsGetParams() *UnitsGetParams {
	selectFields, _ := utils.Fields(&Unit{})
	return &UnitsGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type UnitsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
