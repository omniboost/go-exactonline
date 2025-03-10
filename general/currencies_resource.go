package general

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/rest"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	CurrenciesEndpoint = "/v1/{division}/general/Currencies"
)

func NewCurrenciesResource(rest *rest.Client) *CurrenciesResource {
	return &CurrenciesResource{
		endpoint: CurrenciesEndpoint,
		rest:     rest,
	}
}

type CurrenciesResource struct {
	endpoint string
	rest     *rest.Client
}

func (r *CurrenciesResource) Get(requestParams *CurrenciesGetParams, ctx context.Context) (*CurrenciesGetResponse, error) {
	method := http.MethodGet
	responseBody := r.NewGetResponse()
	path := r.Path()

	// create a new HTTP request
	httpReq, err := r.rest.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// Process query parameters
	utils.AddQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = r.rest.Do(httpReq, responseBody)
	return responseBody, err
}

func (r *CurrenciesResource) Path() string {
	return r.rest.SubPath(r.endpoint)
}

func (r *CurrenciesResource) NewGetParams() *CurrenciesGetParams {
	selectFields, _ := utils.Fields(&Currency{})
	// expandFields := []string{"BankEntryLines"}
	return &CurrenciesGetParams{
		Select: odata.NewSelect(selectFields),
		// Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

func (r *CurrenciesResource) NewGetResponse() *CurrenciesGetResponse {
	return &CurrenciesGetResponse{}
}

type CurrenciesGetResponse struct {
	Results Currencies `json:"results"`
}

func (s *Service) NewCurrenciesGetParams() *CurrenciesGetParams {
	selectFields, _ := utils.Fields(&Currency{})
	// expandFields := []string{"BankEntryLines"}
	return &CurrenciesGetParams{
		Select: odata.NewSelect(selectFields),
		// Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type CurrenciesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	// Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
