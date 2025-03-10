package financial

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	JournalsEndpoint = "/v1/{division}/financial/Journals"
)

// Journals endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialJournals

func (s *Service) JournalsGet(requestParams *JournalsGetParams, ctx context.Context) (*JournalsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewJournalsGetResponse()
	path := s.rest.SubPath(JournalsEndpoint)

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

func (s *Service) NewJournalsGetResponse() *JournalsGetResponse {
	return &JournalsGetResponse{}
}

type JournalsGetResponse struct {
	Results Journals `json:"results"`
}

func (s *Service) NewJournalsGetParams() *JournalsGetParams {
	selectFields, _ := utils.Fields(&Journal{})
	return &JournalsGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type JournalsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
