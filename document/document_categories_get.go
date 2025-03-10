package document

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	DocumentCategoriesEndpoint = "/v1/{division}/documents/DocumentCategories"
)

func (s *Service) DocumentCategoriesGet(requestParams *DocumentCategoriesGetParams, ctx context.Context) (*DocumentCategoriesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewDocumentCategoriesGetResponse()
	path := s.rest.SubPath(DocumentCategoriesEndpoint)

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

func (s *Service) NewDocumentCategoriesGetResponse() *DocumentCategoriesGetResponse {
	return &DocumentCategoriesGetResponse{}
}

type DocumentCategoriesGetResponse struct {
	Results DocumentCategories `json:"results"`
}

func (s *Service) NewDocumentCategoriesGetParams() *DocumentCategoriesGetParams {
	selectFields, _ := utils.Fields(&DocumentAttachment{})
	expandFields := []string{}
	return &DocumentCategoriesGetParams{
		Select:  odata.NewSelect(selectFields),
		Expand:  odata.NewExpand(expandFields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		OrderBy: odata.NewOrderBy(selectFields),
	}
}

type DocumentCategoriesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select  *odata.Select  `schema:"$select,omitempty"`
	Expand  *odata.Expand  `schema:"$expand,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}
