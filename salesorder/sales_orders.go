package salesorder

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	SalesOrdersEndpoint = "/v1/{division}/salesorder/SalesOrders"
)

// SalesOrders endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesOrderSalesOrders

func (s *Service) SalesOrdersGet(requestParams *SalesOrdersGetParams, ctx context.Context) (*SalesOrdersGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewSalesOrdersGetResponse()
	path := s.rest.SubPath(SalesOrdersEndpoint)

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

func (s *Service) NewSalesOrdersGetResponse() *SalesOrdersGetResponse {
	return &SalesOrdersGetResponse{}
}

type SalesOrdersGetResponse struct {
	Results SalesOrders `json:"results"`
}

func (s *Service) NewSalesOrdersGetParams() *SalesOrdersGetParams {
	selectFields, _ := utils.Fields(&SalesOrder{})
	expandFields := []string{"SalesOrderLines"}
	return &SalesOrdersGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type SalesOrdersGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
