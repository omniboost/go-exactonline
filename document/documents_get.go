package document

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	DocumentsEndpoint = "/v1/{division}/documents/Documents"
)

func (s *Service) DocumentsGet(requestParams *DocumentsGetParams, ctx context.Context) (*DocumentsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewDocumentsGetResponse()
	path := s.rest.SubPath(DocumentsEndpoint)

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

func (s *Service) NewDocumentsGetResponse() *DocumentsGetResponse {
	return &DocumentsGetResponse{}
}

type DocumentsGetResponse struct {
	Results Documents `json:"results"`
}

func (s *Service) NewDocumentsGetParams() *DocumentsGetParams {
	selectFields, _ := utils.Fields(&Document{})
	expandFields := []string{}
	return &DocumentsGetParams{
		Select:  odata.NewSelect(selectFields),
		Expand:  odata.NewExpand(expandFields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		OrderBy: odata.NewOrderBy(selectFields),
	}
}

type DocumentsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select  *odata.Select  `schema:"$select,omitempty"`
	Expand  *odata.Expand  `schema:"$expand,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}
