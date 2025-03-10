package crm

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	AccountsEndpoint = "/v1/{division}/crm/Accounts{id}"
)

// Accounts endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=crmAccounts

func (s *Service) AccountsGet(requestParams *AccountsGetParams, ctx context.Context) (*AccountsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewAccountsGetResponse()
	path := s.rest.SubPath(AccountsEndpoint)

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

func (s *Service) NewAccountsGetResponse() *AccountsGetResponse {
	return &AccountsGetResponse{}
}

type AccountsGetResponse struct {
	Results Accounts `json:"results"`
	Next    edm.URL  `json:"__next"`
}

func (s *Service) NewAccountsGetParams() *AccountsGetParams {
	selectFields, _ := utils.Fields(&Account{})
	expandFields := []string{"BankAccounts"}
	return &AccountsGetParams{
		Select:  odata.NewSelect(selectFields),
		Expand:  odata.NewExpand(expandFields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		Skip:    odata.NewSkip(),
		OrderBy: odata.NewOrderBy(selectFields),
	}
}

type AccountsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select  *odata.Select  `schema:"$select,omitempty"`
	Expand  *odata.Expand  `schema:"$expand,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	Skip    *odata.Skip    `schema:"$skip,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}
