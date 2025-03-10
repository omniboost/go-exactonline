package salesinvoice

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	SalesInvoicesEndpoint = "/v1/{division}/salesinvoice/SalesInvoices"
)

// SalesInvoices endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesInvoiceSalesInvoices

// GET

func (s *Service) SalesInvoicesGet(requestParams *SalesInvoicesGetParams, ctx context.Context) (*SalesInvoicesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewSalesInvoicesGetResponse()
	path := s.rest.SubPath(SalesInvoicesEndpoint)

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

func (s *Service) NewSalesInvoicesGetResponse() *SalesInvoicesGetResponse {
	return &SalesInvoicesGetResponse{}
}

type SalesInvoicesGetResponse struct {
	Results SalesInvoices `json:"results"`
}

func (s *Service) NewSalesInvoicesGetParams() *SalesInvoicesGetParams {
	selectFields, _ := utils.Fields(&SalesInvoice{})
	expandFields := []string{"SalesInvoiceLines"}
	return &SalesInvoicesGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type SalesInvoicesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}

// POST

func (s *Service) SalesInvoicesPost(body *SalesInvoicesPostBody, ctx context.Context) (*SalesInvoicesPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewSalesInvoicesPostResponse()
	path := s.rest.SubPath(SalesInvoicesEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type SalesInvoicesPostBody NewSalesInvoice

func (s *Service) NewSalesInvoicesPostBody() *SalesInvoicesPostBody {
	return &SalesInvoicesPostBody{
		SalesInvoiceLines: NewSalesInvoiceLines{},
	}
}

func (s *Service) NewSalesInvoicesPostResponse() *SalesInvoicesPostResponse {
	return &SalesInvoicesPostResponse{}
}

type SalesInvoicesPostResponse struct {
	Results SalesInvoices `json:"results"`
}
