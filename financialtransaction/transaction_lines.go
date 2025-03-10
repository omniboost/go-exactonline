package financialtransaction

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	TransactionLinesEndpoint = "/v1/{division}/financialtransaction/TransactionLines"
)

// TransactionLines endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionTransactionLines

func (s *Service) TransactionLinesGet(requestParams *TransactionLinesGetParams, ctx context.Context) (*TransactionLinesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewTransactionLinesGetResponse()
	path := s.rest.SubPath(TransactionLinesEndpoint)

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

func (s *Service) NewTransactionLinesGetResponse() *TransactionLinesGetResponse {
	return &TransactionLinesGetResponse{}
}

type TransactionLinesGetResponse struct {
	Results TransactionLines `json:"results"`
}

func (s *Service) NewTransactionLinesGetParams() *TransactionLinesGetParams {
	selectFields, _ := utils.Fields(&TransactionLine{})
	return &TransactionLinesGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type TransactionLinesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
