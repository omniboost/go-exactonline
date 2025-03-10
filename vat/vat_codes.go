package vat

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	VatCodesEndpoint = "/v1/{division}/vat/VATCodes"
)

// VatCodes endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=vatVatCodes

func (s *Service) VatCodesGet(requestParams *VatCodesGetParams, ctx context.Context) (*VatCodesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewVatCodesGetResponse()
	path := s.rest.SubPath(VatCodesEndpoint)

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

func (s *Service) NewVatCodesGetResponse() *VatCodesGetResponse {
	return &VatCodesGetResponse{}
}

type VatCodesGetResponse struct {
	Results VatCodes `json:"results"`
}

func (s *Service) NewVatCodesGetParams() *VatCodesGetParams {
	selectFields, _ := utils.Fields(&VatCode{})
	expandFields := []string{"VATPercentages"}
	return &VatCodesGetParams{
		Select: odata.NewSelect(selectFields),
		Expand: odata.NewExpand(expandFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
	}
}

type VatCodesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Expand *odata.Expand `schema:"$expand,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}
