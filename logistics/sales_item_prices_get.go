package logistics

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	SalesItemPricesEndpoint = "/v1/{division}/logistics/SalesItemPrices"
)

// SalesItemPrices endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsSalesItemPrices

func (s *Service) SalesItemPricesGet(requestParams *SalesItemPricesGetParams, ctx context.Context) (*SalesItemPricesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewSalesItemPricesGetResponse()
	path := s.rest.SubPath(SalesItemPricesEndpoint)

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

func (s *Service) NewSalesItemPricesGetResponse() *SalesItemPricesGetResponse {
	return &SalesItemPricesGetResponse{}
}

type SalesItemPricesGetResponse struct {
	Results SalesItemPrices `json:"results"`
}

func (s *Service) NewSalesItemPricesGetParams() *SalesItemPricesGetParams {
	selectFields, _ := utils.Fields(&Item{})
	return &SalesItemPricesGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type SalesItemPricesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
