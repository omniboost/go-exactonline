package purchaseorder

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	PurchaseOrderLinesEndpoint = "/v1/{division}/purchaseorder/PurchaseOrderLines"
)

// PurchaseOrderLines endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PurchaseOrderPurchaseOrderLines

func (s *Service) PurchaseOrderLinesGet(requestParams *PurchaseOrderLinesGetParams, ctx context.Context) (*PurchaseOrderLinesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewPurchaseOrderLinesGetResponse()
	path := s.rest.SubPath(PurchaseOrderLinesEndpoint)

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

func (s *Service) NewPurchaseOrderLinesGetResponse() *PurchaseOrderLinesGetResponse {
	return &PurchaseOrderLinesGetResponse{}
}

type PurchaseOrderLinesGetResponse struct {
	Results PurchaseOrderLines `json:"results"`
}

func (s *Service) NewPurchaseOrderLinesGetParams() *PurchaseOrderLinesGetParams {
	selectFields, _ := utils.Fields(&PurchaseOrder{})
	return &PurchaseOrderLinesGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type PurchaseOrderLinesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
