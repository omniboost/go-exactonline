package financialtransaction

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	CashEntriesEndpoint = "/v1/{division}/financialtransaction/CashEntries"
)

// CashEntries endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionCashEntries

func (s *Service) CashEntriesGet(requestParams *CashEntriesGetParams, ctx context.Context) (*CashEntriesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewCashEntriesGetResponse()
	path := s.rest.SubPath(CashEntriesEndpoint)

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

func (s *Service) NewCashEntriesGetResponse() *CashEntriesGetResponse {
	return &CashEntriesGetResponse{}
}

type CashEntriesGetResponse struct {
	Results CashEntries `json:"results"`
}

func (s *Service) NewCashEntriesGetParams() *CashEntriesGetParams {
	selectFields, _ := utils.Fields(&CashEntry{})
	expandFields := []string{"CashEntryLines"}
	return &CashEntriesGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type CashEntriesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
