package document

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	DocumentAttachmentsEndpoint = "/v1/{division}/documents/DocumentAttachments"
)

func (s *Service) DocumentAttachmentsGet(requestParams *DocumentAttachmentsGetParams, ctx context.Context) (*DocumentAttachmentsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewDocumentAttachmentsGetResponse()
	path := s.rest.SubPath(DocumentAttachmentsEndpoint)

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

func (s *Service) NewDocumentAttachmentsGetResponse() *DocumentAttachmentsGetResponse {
	return &DocumentAttachmentsGetResponse{}
}

type DocumentAttachmentsGetResponse struct {
	Results DocumentAttachments `json:"results"`
}

func (s *Service) NewDocumentAttachmentsGetParams() *DocumentAttachmentsGetParams {
	selectFields, _ := utils.Fields(&DocumentAttachment{})
	expandFields := []string{}
	return &DocumentAttachmentsGetParams{
		Select:  odata.NewSelect(selectFields),
		Expand:  odata.NewExpand(expandFields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		OrderBy: odata.NewOrderBy(selectFields),
	}
}

type DocumentAttachmentsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select  *odata.Select  `schema:"$select,omitempty"`
	Expand  *odata.Expand  `schema:"$expand,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}
