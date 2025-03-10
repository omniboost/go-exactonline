package system

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/utils"
)

// Divisions endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SystemSystemDivisions

const (
	DivisionsEndpoint = "/v1/{division}/system/Divisions"
)

func (s *Service) DivisionsGet(requestParams *DivisionsGetParams, ctx context.Context) (*DivisionsGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewDivisionsGetResponse()
	path := s.rest.SubPath(DivisionsEndpoint)

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

func (s *Service) NewDivisionsGetResponse() *DivisionsGetResponse {
	return &DivisionsGetResponse{}
}

type DivisionsGetResponse []DivisionsUser

type DivisionsUser struct {
	MetaData                edm.MetaData   `json:"__metadata"`
	Code                    edm.Int32      `json:"Code"`
	AddressLine1            edm.String     `json:"AddressLine1"`
	AddressLine2            edm.String     `json:"AddressLine2"`
	AddressLine3            edm.String     `json:"AddressLine3"`
	ChamberOfCommerceNumber edm.String     `json:"ChamberOfCommerceNumber"`
	City                    edm.String     `json:"City"`
	Country                 edm.String     `json:"Country"`
	Created                 edm.DateTime   `json:"Created"`
	Currency                edm.String     `json:"Currency"`
	Current                 edm.Boolean    `json:"Current"`
	Customer                edm.GUID       `json:"Customer"`
	CustomerCode            edm.String     `json:"CustomerCode"`
	CustomerName            edm.String     `json:"CustomerName"`
	Description             edm.String     `json:"Description"`
	Email                   edm.String     `json:"Email"`
	Hid                     edm.Int64      `json:"Hid,string"`
	IsMainDivision          edm.Boolean    `json:"IsMainDivision"`
	Modified                edm.DateTime   `json:"Modified"`
	Phone                   edm.String     `json:"Phone"`
	Postcode                edm.String     `json:"Postcode"`
	State                   edm.String     `json:"State"`
	Status                  DivisionStatus `json:"Status"`
}

func (s *Service) NewDivisionsGetParams() *DivisionsGetParams {
	return &DivisionsGetParams{}
}

type DivisionsGetParams struct {
	// @TODO: check if this an OData struct or something
	Select string `schema:"$select,omitempty"`
	Filter string `schema:"$filter,omitempty"`
}

func (p *DivisionsGetParams) FromQueryParams(queryParams url.Values) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(p, queryParams)
}
