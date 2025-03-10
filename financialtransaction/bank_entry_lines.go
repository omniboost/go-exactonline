package financialtransaction

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	BankEntryLinesEndpoint = "/v1/{division}/financialtransaction/BankEntryLines"
)

// BankEntryLines endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionBankEntryLines

func (s *Service) BankEntryLinesGet(requestParams *BankEntryLinesGetParams, ctx context.Context) (*BankEntryLinesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewBankEntryLinesGetResponse()
	path := s.rest.SubPath(BankEntryLinesEndpoint)

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

func (s *Service) NewBankEntryLinesGetResponse() *BankEntryLinesGetResponse {
	return &BankEntryLinesGetResponse{}
}

type BankEntryLinesGetResponse struct {
	Results BankEntryLines `json:"results"`
}

func (s *Service) NewBankEntryLinesGetParams() *BankEntryLinesGetParams {
	selectFields, _ := utils.Fields(&BankEntryLine{})
	return &BankEntryLinesGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type BankEntryLinesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
