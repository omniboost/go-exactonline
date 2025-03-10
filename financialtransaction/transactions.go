package financialtransaction

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	TransactionsEndpoint = "/v1/{division}/financialtransaction/Transactions"
)

// Transactions endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionTransactions

func (s *Service) TransactionsGet(requestParams *TransactionsGetParams, ctx context.Context) (*TransactionsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewTransactionsGetResponse()
	path := s.rest.SubPath(TransactionsEndpoint)

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

func (s *Service) NewTransactionsGetResponse() *TransactionsGetResponse {
	return &TransactionsGetResponse{}
}

type TransactionsGetResponse struct {
	Results Transactions `json:"results"`
}

func (s *Service) NewTransactionsGetParams() *TransactionsGetParams {
	selectFields, _ := utils.Fields(&Transaction{})
	expandFields := []string{"TransactionLines"}
	return &TransactionsGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type TransactionsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
