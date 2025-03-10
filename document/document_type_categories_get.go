package document

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	DocumentTypeCategoriesEndpoint = "/v1/{division}/documents/DocumentTypeCategories"
)

func (s *Service) DocumentTypeCategoriesGet(requestParams *DocumentTypeCategoriesGetParams, ctx context.Context) (*DocumentTypeCategoriesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewDocumentTypeCategoriesGetResponse()
	path := s.rest.SubPath(DocumentTypeCategoriesEndpoint)

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

func (s *Service) NewDocumentTypeCategoriesGetResponse() *DocumentTypeCategoriesGetResponse {
	return &DocumentTypeCategoriesGetResponse{}
}

type DocumentTypeCategoriesGetResponse struct {
	Results DocumentTypeCategories `json:"results"`
}

func (s *Service) NewDocumentTypeCategoriesGetParams() *DocumentTypeCategoriesGetParams {
	selectFields, _ := utils.Fields(&DocumentAttachment{})
	expandFields := []string{}
	return &DocumentTypeCategoriesGetParams{
		Select:  odata.NewSelect(selectFields),
		Expand:  odata.NewExpand(expandFields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		OrderBy: odata.NewOrderBy(selectFields),
	}
}

type DocumentTypeCategoriesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select  *odata.Select  `schema:"$select,omitempty"`
	Expand  *odata.Expand  `schema:"$expand,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}
