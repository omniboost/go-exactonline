package salesentry

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
)

// POST

func (s *Service) SalesEntriesPost(body *SalesEntriesPostBody, ctx context.Context) (*SalesEntriesPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewSalesEntriesPostResponse()
	path := s.rest.SubPath(SalesEntriesEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type SalesEntriesPostBody NewSalesEntry

func (s *Service) NewSalesEntriesPostBody() *SalesEntriesPostBody {
	return &SalesEntriesPostBody{
		SalesEntryLines: NewSalesEntryLines{},
	}
}

func (s *Service) NewSalesEntriesPostResponse() *SalesEntriesPostResponse {
	return &SalesEntriesPostResponse{}
}

type SalesEntriesPostResponse struct {
	MetaData edm.MetaData `json:"__metadata"`
	SalesEntry
}

type NewSalesEntry struct {
	EntryID          edm.GUID     `json:"EntryID,omitempty"`     // The unique ID of the entry. Via this ID all transaction lines of a single entry can be retrieved
	BatchNumber      edm.Int32    `json:"BatchNumber"`           // The number of the batch of entries. Normally a batch consists of multiple entries. Batchnumbers are filled for invoices created by: - Fixed entries - Prolongation (only available with module hosting)
	Currency         edm.String   `json:"Currency"`              // Currency for the invoice. By default this is the currency of the administration
	Customer         edm.GUID     `json:"Customer"`              // Reference to customer (account)
	Description      edm.String   `json:"Description"`           // Description. Can be different for header and lines
	Document         edm.GUID     `json:"Document"`              // Document that is manually linked to the invoice
	DueDate          edm.DateTime `json:"DueDate"`               // The due date for payments. This date is calculated based on the EntryDate and the Paymentcondition
	EntryDate        edm.DateTime `json:"EntryDate"`             // The date when the invoice is entered
	EntryNumber      edm.Int32    `json:"EntryNumber,omitempty"` // Entry number
	GAccountAmountFC edm.Double   `json:"GAccountAmountFC"`      // A positive value of the amount indicates that the amount is to be paid by the customer to your G bank account.In case of a credit invoice the amount should have negative value when retrieved or posted to Exact.
	InvoiceNumber    edm.Int32    `json:"InvoiceNumber"`         // Assigned at entry or at printing depending on setting. The number assigned is based on the freenumbers as defined for the Journal. When printing the field InvoiceNumber is copied to the fields EntryNumber and InvoiceNumber of the sales entry
	IsExtraDuty      edm.Boolean  `json:"IsExtraDuty"`           // Indicates whether the invoice has extra duty
	Journal          edm.String   `json:"Journal"`               // The journal code. Every invoice should be linked to a sales journal
	OrderNumber      edm.Int32    `json:"OrderNumber"`           // Number to indentify the invoice. Order numbers are not unique. Default the number is based on a setting for the first free number
	PaymentCondition edm.String   `json:"PaymentCondition"`      // The payment condition used for due date and discount calculation
	PaymentReference edm.String   `json:"PaymentReference"`      // The payment reference used for bank imports, VAT return and Tax reference
	ProcessNumber    edm.Int32    `json:"ProcessNumber"`         //
	Rate             edm.Double   `json:"Rate"`                  // Foreign currency rate
	// The period of the transaction lines. The period should exist in the
	// period date table. If it the value is empty, Exact Online will determine
	// the period based on the date table
	ReportingPeriod edm.Int16          `json:"ReportingPeriod,omitempty"`
	ReportingYear   edm.Int16          `json:"ReportingYear,omitempty"` // The financial year to which the entry belongs. The financial year should exist in the period date table
	Reversal        edm.Boolean        `json:"Reversal"`                // Indicates if amounts are reversed
	SalesEntryLines NewSalesEntryLines `json:"SalesEntryLines"`         // Collection of lines
	VATAmountFC     edm.Double         `json:"VATAmountFC"`             // Vat amount in the currency of the transaction
	YourRef         edm.String         `json:"YourRef"`                 // The invoice number of the customer
}
