package purchaseentry

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/omitempty"
)

// POST

func (s *Service) PurchaseEntriesPost(body *PurchaseEntriesPostBody, ctx context.Context) (*PurchaseEntriesPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewPurchaseEntriesPostResponse()
	path := s.rest.SubPath(PurchaseEntriesEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type PurchaseEntriesPostBody NewPurchaseEntry

func (b PurchaseEntriesPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(b)
}

func (s *Service) NewPurchaseEntriesPostBody() *PurchaseEntriesPostBody {
	return &PurchaseEntriesPostBody{
		PurchaseEntryLines: NewPurchaseEntryLines{},
	}
}

func (s *Service) NewPurchaseEntriesPostResponse() *PurchaseEntriesPostResponse {
	return &PurchaseEntriesPostResponse{}
}

type PurchaseEntriesPostResponse struct {
	MetaData edm.MetaData `json:"__metadata"`
	PurchaseEntry
}

type NewPurchaseEntry struct {
	EntryID               edm.GUID              `json:"EntryID,omitempty"`               // Primary key
	BatchNumber           edm.Int32             `json:"BatchNumber,omitempty"`           // The number of the batch of entries. Normally a batch consists of multiple entries. Batchnumbers are filled for invoices created by: - Fixed entries - Prolongation (only available with module hosting)
	Currency              edm.String            `json:"Currency,omitempty"`              // Currency code
	Description           edm.String            `json:"Description,omitempty"`           // Description
	Document              edm.GUID              `json:"Document,omitempty"`              // Reference to document
	DueDate               edm.DateTime          `json:"DueDate,omitempty"`               // Date when payment should be done
	EntryDate             edm.DateTime          `json:"EntryDate,omitempty"`             // Invoice date
	EntryNumber           edm.Int32             `json:"EntryNumber,omitempty"`           // Entry number
	ExternalLinkReference edm.String            `json:"ExternalLinkReference,omitempty"` // External link
	GAccountAmountFC      edm.Double            `json:"GAccountAmountFC,omitempty"`      // A positive value of the amount indicates that the amount is to be paid to the suppliers G bank account.In case of a credit invoice the amount should have negative value when retrieved or posted to Exact.
	InvoiceNumber         edm.Int32             `json:"InvoiceNumber,omitempty"`         // Invoice number
	Journal               edm.String            `json:"Journal"`                         // Journal
	OrderNumber           edm.Int32             `json:"OrderNumber,omitempty"`           // Order number
	PaymentCondition      edm.String            `json:"PaymentCondition,omitempty"`      // Payment condition
	PaymentReference      edm.String            `json:"PaymentReference,omitempty"`      // The payment reference used for bank imports, VAT return and Tax reference
	ProcessNumber         edm.Int32             `json:"ProcessNumber,omitempty"`         // Internal processing number, only relevant for Germany
	PurchaseEntryLines    NewPurchaseEntryLines `json:"PurchaseEntryLines"`              // Collection of lines
	Rate                  edm.Double            `json:"Rate,omitempty"`                  // Currency exchange rate
	ReportingPeriod       edm.Int16             `json:"ReportingPeriod,omitempty"`       // The period of the transaction lines. The period should exist in the period date table
	ReportingYear         edm.Int16             `json:"ReportingYear,omitempty"`         // The financial year to which the entry belongs. The financial year should exist in the period date table
	Reversal              edm.Boolean           `json:"Reversal,omitempty"`              // Indicates that amounts are reversed
	Supplier              edm.GUID              `json:"Supplier"`                        // Reference to supplier (account)
	VATAmountFC           edm.Double            `json:"VATAmountFC,omitempty"`           // Vat Amount in the currency of the transaction
	YourRef               edm.String            `json:"YourRef,omitempty"`               // Your reference
}

func (e NewPurchaseEntry) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(e)
}

type PurchaseEntry struct {
	EntryID                     edm.GUID     `json:"EntryID,omitempty"`                     // Primary key
	AmountDC                    edm.Double   `json:"AmountDC,omitempty"`                    // Amount in the default currency of the company
	AmountFC                    edm.Double   `json:"AmountFC,omitempty"`                    // Amount in the currency of the transaction
	BatchNumber                 edm.Int32    `json:"BatchNumber,omitempty"`                 // The number of the batch of entries. Normally a batch consists of multiple entries. Batchnumbers are filled for invoices created by: - Fixed entries - Prolongation (only available with module hosting)
	Created                     edm.DateTime `json:"Created,omitempty"`                     // Creation date
	Creator                     edm.GUID     `json:"Creator,omitempty"`                     // User ID of creator
	CreatorFullName             edm.String   `json:"CreatorFullName,omitempty"`             // Name of creator
	Currency                    edm.String   `json:"Currency,omitempty"`                    // Currency code
	Description                 edm.String   `json:"Description,omitempty"`                 // Description
	Division                    edm.Int32    `json:"Division,omitempty"`                    // Division code
	Document                    edm.GUID     `json:"Document,omitempty"`                    // Reference to document
	DocumentNumber              edm.Int32    `json:"DocumentNumber,omitempty"`              // Document number
	DocumentSubject             edm.String   `json:"DocumentSubject,omitempty"`             // Document subject
	DueDate                     edm.DateTime `json:"DueDate,omitempty"`                     // Date when payment should be done
	EntryDate                   edm.DateTime `json:"EntryDate,omitempty"`                   // Invoice date
	EntryNumber                 edm.Int32    `json:"EntryNumber,omitempty"`                 // Entry number
	ExternalLinkDescription     edm.String   `json:"ExternalLinkDescription,omitempty"`     // Description of ExternalLink
	ExternalLinkReference       edm.String   `json:"ExternalLinkReference,omitempty"`       // External link
	GAccountAmountFC            edm.Double   `json:"GAccountAmountFC,omitempty"`            // A positive value of the amount indicates that the amount is to be paid to the suppliers G bank account.In case of a credit invoice the amount should have negative value when retrieved or posted to Exact.
	InvoiceNumber               edm.Int32    `json:"InvoiceNumber,omitempty"`               // Invoice number
	Journal                     edm.String   `json:"Journal"`                               // Journal
	JournalDescription          edm.String   `json:"JournalDescription,omitempty"`          // Description of Journal
	Modified                    edm.DateTime `json:"Modified,omitempty"`                    // Last modified date
	Modifier                    edm.GUID     `json:"Modifier,omitempty"`                    // User ID of modifier
	ModifierFullName            edm.String   `json:"ModifierFullName,omitempty"`            // Name of modifier
	OrderNumber                 edm.Int32    `json:"OrderNumber,omitempty"`                 // Order number
	PaymentCondition            edm.String   `json:"PaymentCondition,omitempty"`            // Payment condition
	PaymentConditionDescription edm.String   `json:"PaymentConditionDescription,omitempty"` // Description of PaymentCondition
	PaymentReference            edm.String   `json:"PaymentReference,omitempty"`            // The payment reference used for bank imports, VAT return and Tax reference
	ProcessNumber               edm.Int32    `json:"ProcessNumber,omitempty"`               // Internal processing number, only relevant for Germany
	PurchaseEntryLines          interface{}  `json:"PurchaseEntryLines"`                    // Collection of lines
	Rate                        edm.Double   `json:"Rate,omitempty"`                        // Currency exchange rate
	ReportingPeriod             edm.Int16    `json:"ReportingPeriod,omitempty"`             // The period of the transaction lines. The period should exist in the period date table
	ReportingYear               edm.Int16    `json:"ReportingYear,omitempty"`               // The financial year to which the entry belongs. The financial year should exist in the period date table
	Reversal                    edm.Boolean  `json:"Reversal,omitempty"`                    // Indicates that amounts are reversed
	Status                      edm.Int16    `json:"Status,omitempty"`                      // Status: 5 = Rejected, 20 = Open, 50 = Processed
	StatusDescription           edm.String   `json:"StatusDescription,omitempty"`           // Description of Status
	Supplier                    edm.GUID     `json:"Supplier"`                              // Reference to supplier (account)
	SupplierName                edm.String   `json:"SupplierName,omitempty"`                // Name of supplier
	Type                        edm.Int32    `json:"Type,omitempty"`                        // Type: 30 = Purchase entry, 31 = Purchase credit note
	TypeDescription             edm.String   `json:"TypeDescription,omitempty"`             // Description of Type
	VATAmountDC                 edm.Double   `json:"VATAmountDC,omitempty"`                 // Vat Amount in the default currency of the company
	VATAmountFC                 edm.Double   `json:"VATAmountFC,omitempty"`                 // Vat Amount in the currency of the transaction
	YourRef                     edm.String   `json:"YourRef,omitempty"`                     // Your reference
}

func (e PurchaseEntry) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(e)
}

type PurchaseEntryLine struct {
	ID                         edm.GUID     `json:"ID,omitempty"`                         // Primary key
	AmountDC                   edm.Double   `json:"AmountDC,omitempty"`                   // Amount in the default currency of the company
	AmountFC                   edm.Double   `json:"AmountFC"`                             // Amount in the currency of the transaction
	Asset                      edm.GUID     `json:"Asset,omitempty"`                      // Reference to asset
	AssetDescription           edm.String   `json:"AssetDescription,omitempty"`           // Asset description
	CostCenter                 edm.String   `json:"CostCenter,omitempty"`                 // Reference to cost center
	CostCenterDescription      edm.String   `json:"CostCenterDescription,omitempty"`      // Description of CostCenter
	CostUnit                   edm.String   `json:"CostUnit,omitempty"`                   // Reference to cost unit
	CostUnitDescription        edm.String   `json:"CostUnitDescription,omitempty"`        // Description of CostUnit
	Description                edm.String   `json:"Description,omitempty"`                // Description
	Division                   edm.Int32    `json:"Division,omitempty"`                   // Division code
	EntryID                    edm.GUID     `json:"EntryID"`                              // Reference to header of the purchase entry
	From                       edm.DateTime `json:"From,omitempty"`                       // From date to identify the range for deferred costs. This is used in combination with the property 'To' that defines the end date
	GLAccount                  edm.GUID     `json:"GLAccount"`                            // General ledger account
	GLAccountCode              edm.String   `json:"GLAccountCode,omitempty"`              // Code of GLAccount
	GLAccountDescription       edm.String   `json:"GLAccountDescription,omitempty"`       // Description of GLAccount
	IntraStatArea              edm.String   `json:"IntraStatArea,omitempty"`              // IntraStat area (only relevant when IntraStat for purchase is enabled in the administration)
	IntraStatCountry           edm.String   `json:"IntraStatCountry,omitempty"`           // IntraStatCountry (only relevant when IntraStat for purchase is enabled in the administration)
	IntraStatDeliveryTerm      edm.String   `json:"IntraStatDeliveryTerm,omitempty"`      // IntraStat delivery term (only relevant when IntraStat for purchase is enabled in the administration)
	IntraStatTransactionA      edm.String   `json:"IntraStatTransactionA,omitempty"`      // IntraStat transaction A (only relevant when IntraStat for purchase is enabled in the administration)
	IntraStatTransactionB      edm.String   `json:"IntraStatTransactionB,omitempty"`      // IntraStat transaction B (only relevant when IntraStat for purchase is enabled in the Belgium, Luxembourg, France & United Kingdom administration)
	IntraStatTransportMethod   edm.String   `json:"IntraStatTransportMethod,omitempty"`   // IntraStat transport method (only relevant when IntraStat for purchase is enabled in the administration)
	LineNumber                 edm.Int32    `json:"LineNumber,omitempty"`                 // Line number
	Notes                      edm.String   `json:"Notes,omitempty"`                      // Extra remarks
	PrivateUsePercentage       edm.Double   `json:"PrivateUsePercentage,omitempty"`       // Percentage of re-invoice part of a cost to the owner of the company.
	Project                    edm.GUID     `json:"Project,omitempty"`                    // Reference to project
	ProjectDescription         edm.String   `json:"ProjectDescription,omitempty"`         // Description of Project
	Quantity                   edm.Double   `json:"Quantity,omitempty"`                   // Quantity
	SerialNumber               edm.String   `json:"SerialNumber,omitempty"`               // Serial number
	StatisticalNetWeight       edm.Double   `json:"StatisticalNetWeight,omitempty"`       // Statistical NetWeight (only relevant when IntraStat for purchase is enabled in the administration)
	StatisticalNumber          edm.String   `json:"StatisticalNumber,omitempty"`          // Statistical Number (only relevant when IntraStat for purchase is enabled in the administration)
	StatisticalQuantity        edm.Double   `json:"StatisticalQuantity,omitempty"`        // Statistical Quantity (only relevant when IntraStat for purchase is enabled in the administration)
	StatisticalValue           edm.Double   `json:"StatisticalValue,omitempty"`           // Statistical Value (only relevant when IntraStat for purchase is enabled in the administration)
	Subscription               edm.GUID     `json:"Subscription,omitempty"`               // Reference to subscription
	SubscriptionDescription    edm.String   `json:"SubscriptionDescription,omitempty"`    // Description of Subscription
	To                         edm.DateTime `json:"To,omitempty"`                         // To date to identify the range for deferred costs. This is used in combination with the property 'From' that defines the start date
	TrackingNumber             edm.GUID     `json:"TrackingNumber,omitempty"`             // Reference to tracking number
	TrackingNumberDescription  edm.String   `json:"TrackingNumberDescription,omitempty"`  // Description of TrackingNumber
	Type                       edm.Int32    `json:"Type,omitempty"`                       // Type: 30 = Purchase entry, 31 = Purchase credit note
	VATAmountDC                edm.Double   `json:"VATAmountDC,omitempty"`                // VAT amount in the default currency of the company
	VATAmountFC                edm.Double   `json:"VATAmountFC,omitempty"`                // VAT amount in the currency of the transaction. Use this property to specify a VAT amount that differs from the VAT amount that is automatically calculated.
	VATBaseAmountDC            edm.Double   `json:"VATBaseAmountDC,omitempty"`            // VAT base amount in the default currency of the company
	VATBaseAmountFC            edm.Double   `json:"VATBaseAmountFC,omitempty"`            // VAT base amount in the currency of the transaction
	VATCode                    edm.String   `json:"VATCode,omitempty"`                    // VAT code. If this property is not filled, it will use the default VAT code of the G/L account property
	VATCodeDescription         edm.String   `json:"VATCodeDescription,omitempty"`         // Description of VATCode
	VATNonDeductiblePercentage edm.Double   `json:"VATNonDeductiblePercentage,omitempty"` // If not the full amount of the VAT is deductible, you can indicate a percentage for the non decuctible part. This is used during the entry of purchase invoices.
	VATPercentage              edm.Double   `json:"VATPercentage,omitempty"`              // VAT percentage
	WithholdingAmountDC        edm.Double   `json:"WithholdingAmountDC,omitempty"`        // Withholding tax amount for spanish legislation
	WithholdingTax             edm.String   `json:"WithholdingTax,omitempty"`             // Withholding tax key for spanish legislation debugger eval code:12:9
}

func (l PurchaseEntryLine) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(l)
}
