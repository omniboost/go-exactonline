package document

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	DocumentTypesEndpoint = "/v1/{division}/documents/DocumentTypes"
)

func (s *Service) DocumentTypesGet(requestParams *DocumentTypesGetParams, ctx context.Context) (*DocumentTypesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewDocumentTypesGetResponse()
	path := s.rest.SubPath(DocumentTypesEndpoint)

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

func (s *Service) NewDocumentTypesGetResponse() *DocumentTypesGetResponse {
	return &DocumentTypesGetResponse{}
}

type DocumentTypesGetResponse struct {
	Results DocumentTypes `json:"results"`
}

func (s *Service) NewDocumentTypesGetParams() *DocumentTypesGetParams {
	selectFields, _ := utils.Fields(&DocumentAttachment{})
	expandFields := []string{}
	return &DocumentTypesGetParams{
		Select:  odata.NewSelect(selectFields),
		Expand:  odata.NewExpand(expandFields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		OrderBy: odata.NewOrderBy(selectFields),
	}
}

type DocumentTypesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select  *odata.Select  `schema:"$select,omitempty"`
	Expand  *odata.Expand  `schema:"$expand,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}
