package generaljournalentry

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/utils"
)

const (
	GeneralJournalEntriesEndpoint = "/v1/{division}/generaljournalentry/GeneralJournalEntries"
)

// GeneralJournalEntries endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=GeneralJournalEntryGeneralJournalEntries

// GET

func (s *Service) GeneralJournalEntriesGet(requestParams *GeneralJournalEntriesGetParams, ctx context.Context) (*GeneralJournalEntriesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewGeneralJournalEntriesGetResponse()
	path := s.rest.SubPath(GeneralJournalEntriesEndpoint)

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

func (s *Service) NewGeneralJournalEntriesGetResponse() *GeneralJournalEntriesGetResponse {
	return &GeneralJournalEntriesGetResponse{}
}

type GeneralJournalEntriesGetResponse struct {
	Results GeneralJournalEntries `json:"results"`
}

func (s *Service) NewGeneralJournalEntriesGetParams() *GeneralJournalEntriesGetParams {
	selectFields, _ := utils.Fields(&GeneralJournalEntry{})
	expandFields := []string{"GeneralJournalEntryLines"}
	return &GeneralJournalEntriesGetParams{
		Select:  odata.NewSelect(selectFields),
		Expand:  odata.NewExpand(expandFields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		OrderBy: odata.NewOrderBy(selectFields),
	}
}

type GeneralJournalEntriesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select  *odata.Select  `schema:"$select,omitempty"`
	Expand  *odata.Expand  `schema:"$expand,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}
