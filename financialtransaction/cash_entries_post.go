package financialtransaction

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/omitempty"
)

// POST

func (s *Service) CashEntriesPost(body *CashEntriesPostBody, ctx context.Context) (*CashEntriesPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewCashEntriesPostResponse()
	path := s.rest.SubPath(CashEntriesEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type CashEntriesPostBody NewCashEntry

func (s *Service) NewCashEntriesPostBody() *CashEntriesPostBody {
	return &CashEntriesPostBody{
		CashEntryLines: NewCashEntryLines{},
	}
}

func (s *Service) NewCashEntriesPostResponse() *CashEntriesPostResponse {
	return &CashEntriesPostResponse{}
}

type CashEntriesPostResponse struct{}

type NewCashEntry struct {
	CashEntryLines   NewCashEntryLines `json:"CashEntryLines"`             // Collection of lines
	ClosingBalanceFC edm.Double        `json:"ClosingBalanceFC,omitempty"` // Closing balance in the currency of the transaction
	Currency         edm.String        `json:"Currency"`                   // Currency code
	EntryNumber      edm.Int32         `json:"EntryNumber,omitempty"`      // Entry number
	FinancialPeriod  edm.Int16         `json:"FinancialPeriod"`            // Fiancial period
	FinancialYear    edm.Int16         `json:"FinancialYear"`              // Fiancial year
	JournalCode      edm.String        `json:"JournalCode"`                // Code of Journal
	OpeningBalanceFC edm.Double        `json:"OpeningBalanceFC,omitempty"` // Opening balance in the currency of the transaction
}

type NewCashEntryLines []NewCashEntryLine

type NewCashEntryLine struct {
	Account       edm.GUID     `json:"Account,omitempty"`       // Reference to Account
	AmountFC      edm.Double   `json:"AmountFC"`                // Amount in the currency of the transaction
	AmountVATFC   edm.Double   `json:"AmountVATFC"`             // Vat amount in the currency of the transaction
	Asset         edm.GUID     `json:"Asset,omitempty"`         // Reference to an asset
	CostCenter    edm.String   `json:"CostCenter,omitempty"`    // Reference to a cost center
	CostUnit      edm.String   `json:"CostUnit,omitempty"`      // Reference to a cost unit
	Date          edm.DateTime `json:"Date"`                    // Date
	Description   edm.String   `json:"Description"`             // Description
	Document      edm.GUID     `json:"Document,omitempty"`      // Reference to a document
	EntryID       edm.GUID     `json:"EntryID,omitempty"`       // Reference to the header
	ExchangeRate  edm.Double   `json:"ExchangeRate,omitempty"`  // Exchange rate
	GLAccount     edm.GUID     `json:"GLAccount"`               // General ledger account
	Notes         edm.String   `json:"Notes,omitempty"`         // Extra remarks
	OurRef        edm.Int32    `json:"OurRef,omitempty"`        // Invoice number
	Project       edm.GUID     `json:"Project,omitempty"`       // Reference to a project
	Quantity      edm.Double   `json:"Quantity,omitempty"`      // Quantity
	VATCode       edm.String   `json:"VATCode"`                 // Reference to vat code
	VATPercentage edm.Double   `json:"VATPercentage,omitempty"` // Vat code percentage
	VATType       edm.String   `json:"VATType,omitempty"`       // Type of vat code
}

func (l NewCashEntryLine) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(l)
}
