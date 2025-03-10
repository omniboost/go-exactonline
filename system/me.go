package system

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/utils"
)

// Me endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SystemSystemDivisions

const (
	MeEndpoint = "/v1/current/Me"
)

func (s *Service) MeGet(requestParams *MeGetParams, ctx context.Context) (*MeGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewMeGetResponse()
	path := s.rest.SubPath(MeEndpoint)

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

func (s *Service) NewMeGetResponse() *MeGetResponse {
	return &MeGetResponse{}
}

type MeGetResponse struct {
	Results []MeUser `json:"results"`
}

type MeUser struct {
	MetaData               edm.MetaData `json:"__metadata"`
	UserID                 edm.GUID     `json:"UserID"`
	CurrentDivision        edm.Int32    `json:"CurrentDivision"`
	DivisionCustomer       edm.GUID     `json:"DivisionCustomer"`
	DivisionCustomerCode   edm.String   `json:"DivisionCustomerCode"`
	Email                  edm.String   `json:"Email"`
	EmployeeID             edm.GUID     `json:"EmployeeID"`
	FirstName              edm.String   `json:"FirstName"`
	FullName               edm.String   `json:"FullName"`
	Gender                 Gender       `json:"Gender"`
	Initials               edm.String   `json:"Initials"`
	Language               edm.String   `json:"Language"`
	LanguageCode           edm.String   `json:"LanguageCode"`
	LastName               edm.String   `json:"LastName"`
	Legislation            edm.Int64    `json:"Legislation,string"`
	MiddleName             edm.String   `json:"MiddleName"`
	Mobile                 edm.String   `json:"Mobile"`
	Nationality            edm.String   `json:"Nationality"`
	Phone                  edm.String   `json:"Phone"`
	PhoneExtension         edm.String   `json:"PhoneExtension"`
	PictureURL             utils.URL    `json:"PictureUrl"`
	ServerTime             edm.String   `json:"ServerTime"`
	ServerUtcOffset        edm.Double   `json:"ServerUtcOffset"`
	ThumbnailPicture       edm.Binary   `json:"ThumbnailPicture"`
	ThumbnailPictureFormat edm.String   `json:"ThumbnailPictureFormat"`
	Title                  edm.String   `json:"Title"`
	UserName               edm.String   `json:"UserName"`
}

func (s *Service) NewMeGetParams() *MeGetParams {
	return &MeGetParams{}
}

type MeGetParams struct {
	// @TODO: check if this an OData struct or something
	Select string `schema:"$select"`
	Filter string `schema:"$filter"`
}

func (p *MeGetParams) FromQueryParams(queryParams url.Values) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(p, queryParams)
}
