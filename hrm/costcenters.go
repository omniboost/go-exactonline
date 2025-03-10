package hrm

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	CostcentersEndpoint = "/v1/{division}/hrm/Costcenters"
)

// Costcenters endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMCostcenters

func (s *Service) CostcentersGet(requestParams *CostcentersGetParams, ctx context.Context) (*CostcentersGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewCostcentersGetResponse()
	path := s.rest.SubPath(CostcentersEndpoint)

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

func (s *Service) NewCostcentersGetResponse() *CostcentersGetResponse {
	return &CostcentersGetResponse{}
}

type CostcentersGetResponse struct {
	Results Costcenters `json:"results"`
}

func (s *Service) NewCostcentersGetParams() *CostcentersGetParams {
	selectFields, _ := utils.Fields(&Costcenter{})
	return &CostcentersGetParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type CostcentersGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
