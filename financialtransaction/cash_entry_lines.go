package financialtransaction

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	CashEntryLinesEndpoint = "/v1/{division}/financialtransaction/CashEntryLines"
)

// CashEntryLines endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionCashEntryLines

func (s *Service) CashEntryLinesGet(requestParams *CashEntryLinesGetParams, ctx context.Context) (*CashEntryLinesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewCashEntryLinesGetResponse()
	path := s.rest.SubPath(CashEntryLinesEndpoint)

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

func (s *Service) NewCashEntryLinesGetResponse() *CashEntryLinesGetResponse {
	return &CashEntryLinesGetResponse{}
}

type CashEntryLinesGetResponse struct {
	Results CashEntryLines `json:"results"`
}

func (s *Service) NewCashEntryLinesGetParams() *CashEntryLinesGetParams {
	selectFields, _ := utils.Fields(&CashEntryLine{})
	return &CashEntryLinesGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type CashEntryLinesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
