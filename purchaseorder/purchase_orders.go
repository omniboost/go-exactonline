package purchaseorder

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	PurchaseOrdersEndpoint = "/v1/{division}/purchaseorder/PurchaseOrders"
)

// PurchaseOrders endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PurchaseOrderPurchaseOrders

func (s *Service) PurchaseOrdersGet(requestParams *PurchaseOrdersGetParams, ctx context.Context) (*PurchaseOrdersGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewPurchaseOrdersGetResponse()
	path := s.rest.SubPath(PurchaseOrdersEndpoint)

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

func (s *Service) NewPurchaseOrdersGetResponse() *PurchaseOrdersGetResponse {
	return &PurchaseOrdersGetResponse{}
}

type PurchaseOrdersGetResponse struct {
	Results PurchaseOrders `json:"results"`
}

func (s *Service) NewPurchaseOrdersGetParams() *PurchaseOrdersGetParams {
	selectFields, _ := utils.Fields(&PurchaseOrder{})
	// expandFields := []string{"PurchaseOrderLines"}
	return &PurchaseOrdersGetParams{
		Select: odata.NewSelect(selectFields),
		// Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
	}
}

type PurchaseOrdersGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	// Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
}
