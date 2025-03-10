package logistics

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	ItemsEndpoint = "/v1/{division}/logistics/Items"
)

// Items endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsItems

func (s *Service) ItemsGet(requestParams *ItemsGetParams, ctx context.Context) (*ItemsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewItemsGetResponse()
	path := s.rest.SubPath(ItemsEndpoint)

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

func (s *Service) NewItemsGetResponse() *ItemsGetResponse {
	return &ItemsGetResponse{}
}

type ItemsGetResponse struct {
	Results Items `json:"results"`
}

func (s *Service) NewItemsGetParams() *ItemsGetParams {
	selectFields, _ := utils.Fields(&Item{})
	return &ItemsGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type ItemsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
