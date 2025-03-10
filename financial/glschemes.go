package financial

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	GLSchemesEndpoint = "/v1/{division}/financial/GLSchemes"
)

// GLSchemes endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialGLSchemes

func (s *Service) GLSchemesGet(requestParams *GLSchemesGetParams, ctx context.Context) (*GLSchemesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewGLSchemesGetResponse()
	path := s.rest.SubPath(GLSchemesEndpoint)

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

func (s *Service) NewGLSchemesGetResponse() *GLSchemesGetResponse {
	return &GLSchemesGetResponse{}
}

type GLSchemesGetResponse struct {
	Results GLSchemes `json:"results"`
}

func (s *Service) NewGLSchemesGetParams() *GLSchemesGetParams {
	selectFields, _ := utils.Fields(&GLScheme{})
	return &GLSchemesGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type GLSchemesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
