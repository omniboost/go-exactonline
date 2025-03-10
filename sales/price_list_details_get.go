package sales

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	PriceListDetailsEndpoint = "/v1/{division}/sales/SalesPriceListDetails"
)

// PriceListDetails endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsPriceListDetails

func (s *Service) PriceListDetailsGet(requestParams *PriceListDetailsGetParams, ctx context.Context) (*PriceListDetailsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewPriceListDetailsGetResponse()
	path := s.rest.SubPath(PriceListDetailsEndpoint)

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

func (s *Service) NewPriceListDetailsGetResponse() *PriceListDetailsGetResponse {
	return &PriceListDetailsGetResponse{}
}

type PriceListDetailsGetResponse struct {
	Results PriceListDetails `json:"results"`
}

func (s *Service) NewPriceListDetailsGetParams() *PriceListDetailsGetParams {
	selectFields, _ := utils.Fields(&PriceListDetail{})
	return &PriceListDetailsGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type PriceListDetailsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
