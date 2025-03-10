package purchaseentry

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/odata"
	"github.com/omniboost/go-exactonline/omitempty"
	"github.com/omniboost/go-exactonline/utils"
)

// PurchaseEntries endpoint
// - https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesEntryPurchaseEntries

// GET

func (s *Service) PurchaseEntriesGet(requestParams *PurchaseEntriesGetParams, ctx context.Context) (*PurchaseEntriesGetResponse, error) {
	method := http.MethodGet
	responseBody := s.NewPurchaseEntriesGetResponse()
	path := s.rest.SubPath(PurchaseEntriesEndpoint)

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

func (s *Service) NewPurchaseEntriesGetResponse() *PurchaseEntriesGetResponse {
	return &PurchaseEntriesGetResponse{}
}

type PurchaseEntriesGetResponse struct {
	Results PurchaseEntries `json:"results"`
}

func (s *Service) NewPurchaseEntriesGetParams() *PurchaseEntriesGetParams {
	selectFields, _ := utils.Fields(&SalesEntry{})
	expandFields := []string{"SalesEntryLines"}
	return &PurchaseEntriesGetParams{
		Select:  odata.NewSelect(selectFields),
		Expand:  odata.NewExpand(expandFields),
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		OrderBy: odata.NewOrderBy(selectFields),
	}
}

type PurchaseEntriesGetParams struct {
	// @TODO: check if this an OData struct or something
	Select  *odata.Select  `schema:"$select,omitempty"`
	Expand  *odata.Expand  `schema:"$expand,omitempty"`
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}

type PurchaseEntries []SalesEntry

type SalesEntry struct {
	AmountDC                    edm.Double      `json:"AmountDC"`                    // Amount in the default currency of the company. For the header lines (LineNumber = 0) of an entry this is the SUM(AmountDC) of all lines
	AmountFC                    edm.Double      `json:"AmountFC"`                    // Amount in the currency of the transaction. For the header this is the sum of all lines, including VAT
	BatchNumber                 edm.Int32       `json:"BatchNumber"`                 // The number of the batch of entries. Normally a batch consists of multiple entries. Batchnumbers are filled for invoices created by: - Fixed entries - Prolongation (only available with module hosting)
	Created                     edm.DateTime    `json:"Created,omitempty"`           // Creation date
	Creator                     edm.GUID        `json:"Creator"`                     // User ID of creator
	CreatorFullName             edm.String      `json:"CreatorFullName"`             // Name of creator
	Currency                    edm.String      `json:"Currency"`                    // Currency for the invoice. By default this is the currency of the administration
	Customer                    edm.GUID        `json:"Customer"`                    // Reference to customer (account)
	CustomerName                edm.String      `json:"CustomerName"`                // Name of customer
	Description                 edm.String      `json:"Description"`                 // Description. Can be different for header and lines
	Division                    edm.Int32       `json:"Division"`                    // Division code
	Document                    edm.GUID        `json:"Document"`                    // Document that is manually linked to the invoice
	DocumentNumber              edm.Int32       `json:"DocumentNumber"`              // Number of the document
	DocumentSubject             edm.String      `json:"DocumentSubject"`             // Subject of the document
	DueDate                     edm.DateTime    `json:"DueDate,omitempty"`           // The due date for payments. This date is calculated based on the EntryDate and the Paymentcondition
	EntryDate                   edm.DateTime    `json:"EntryDate,omitempty"`         // The date when the invoice is entered
	EntryNumber                 edm.Int32       `json:"EntryNumber"`                 // Entry number
	ExternalLinkDescription     edm.String      `json:"ExternalLinkDescription"`     // Description of ExternalLink
	ExternalLinkReference       edm.String      `json:"ExternalLinkReference"`       // Reference of ExternalLink
	GAccountAmountFC            edm.Double      `json:"GAccountAmountFC"`            // A positive value of the amount indicates that the amount is to be paid by the customer to your G bank account.In case of a credit invoice the amount should have negative value when retrieved or posted to Exact.
	InvoiceNumber               edm.Int32       `json:"InvoiceNumber"`               // Assigned at entry or at printing depending on setting. The number assigned is based on the freenumbers as defined for the Journal. When printing the field InvoiceNumber is copied to the fields EntryNumber and InvoiceNumber of the sales entry
	IsExtraDuty                 edm.Boolean     `json:"IsExtraDuty"`                 // Indicates whether the invoice has extra duty
	Journal                     edm.String      `json:"Journal"`                     // The journal code. Every invoice should be linked to a sales journal
	JournalDescription          edm.String      `json:"JournalDescription"`          // Description of Journal
	Modified                    edm.DateTime    `json:"Modified"`                    // Last modified date
	Modifier                    edm.GUID        `json:"Modifier"`                    // User ID of modifier
	ModifierFullName            edm.String      `json:"ModifierFullName"`            // Name of modifier
	OrderNumber                 edm.Int32       `json:"OrderNumber"`                 // Number to indentify the invoice. Order numbers are not unique. Default the number is based on a setting for the first free number
	PaymentCondition            edm.String      `json:"PaymentCondition"`            // The payment condition used for due date and discount calculation
	PaymentConditionDescription edm.String      `json:"PaymentConditionDescription"` // Description of PaymentCondition
	PaymentReference            edm.String      `json:"PaymentReference"`            // The payment reference used for bank imports, VAT return and Tax reference
	ProcessNumber               edm.Int32       `json:"ProcessNumber"`               //
	Rate                        edm.Double      `json:"Rate"`                        // Foreign currency rate
	ReportingPeriod             edm.Int16       `json:"ReportingPeriod"`             // The period of the transaction lines. The period should exist in the period date table
	ReportingYear               edm.Int16       `json:"ReportingYear"`               // The financial year to which the entry belongs. The financial year should exist in the period date table
	Reversal                    edm.Boolean     `json:"Reversal"`                    // Indicates if amounts are reversed
	SalesEntryLines             SalesEntryLines `json:"SalesEntryLines"`             // Collection of lines
	Status                      edm.Int16       `json:"Status"`                      // Status: 5 = Rejected, 20 = Open, 50 = Processed
	StatusDescription           edm.String      `json:"StatusDescription"`           // Description of Status
	Type                        edm.Int32       `json:"Type"`                        // Type: 20 = Sales entry, 21 = Sales credit note
	TypeDescription             edm.String      `json:"TypeDescription"`             // Description of Type
	VATAmountDC                 edm.Double      `json:"VATAmountDC"`                 // Vat amount in the default currency of the company
	VATAmountFC                 edm.Double      `json:"VATAmountFC"`                 // Vat amount in the currency of the transaction
	YourRef                     edm.String      `json:"YourRef"`                     // The invoice number of the customer
}

type SalesEntryLines []SalesEntryLine

// standalone: "SalesEntryLines": []
// deferred: "SalesEntryLines": {"__deferred": {}}
// embedded: "SalesEntryLines": {"results": []}
func (l *SalesEntryLines) UnmarshalJSON(data []byte) error {
	type Results SalesEntryLines

	type Envelope struct {
		Results  Results         `json:"results"`
		Deferred json.RawMessage `json:"__deferred"`
	}

	// create the json tester
	tester := &utils.JsonTester{}
	err := json.Unmarshal(data, tester)
	if err != nil {
		return err
	}

	// test if json is array (standalone)
	if tester.IsArray() {
		results := &Results{}
		err := json.Unmarshal(data, results)
		if err != nil {
			return err
		}

		*l = SalesEntryLines(*results)
		return nil
	}

	// lines are embedded or deferred: only try embedded
	envelope := &Envelope{Results: Results(*l)}
	err = json.Unmarshal(data, envelope)
	if err != nil {
		return err
	}

	*l = SalesEntryLines(envelope.Results)
	return nil
}

func (e SalesEntry) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(e)
}

type SalesEntryLine struct {
	ID                        edm.GUID     `json:"ID,omitempty"`              // Primary key
	AmountDC                  edm.Double   `json:"AmountDC"`                  // Amount in the default currency of the company. For almost all lines this can be calculated like: AmountDC = AmountFC * RateFC.
	AmountFC                  edm.Double   `json:"AmountFC"`                  // For normal lines it's the amount excluding VAT
	Asset                     edm.GUID     `json:"Asset"`                     // Reference to Asset
	AssetDescription          edm.String   `json:"AssetDescription"`          // Description of Asset
	CostCenter                edm.String   `json:"CostCenter"`                // Reference to Cost center
	CostCenterDescription     edm.String   `json:"CostCenterDescription"`     // Description of CostCenter
	CostUnit                  edm.String   `json:"CostUnit"`                  // Reference to Cost unit
	CostUnitDescription       edm.String   `json:"CostUnitDescription"`       // Description of CostUnit
	Description               edm.String   `json:"Description"`               // Description. Can be different for header and lines
	Division                  edm.Int32    `json:"Division"`                  // Division code
	EntryID                   edm.GUID     `json:"EntryID,omitempty"`         // The unique ID of the entry. Via this ID all transaction lines of a single entry can be retrieved
	ExtraDutyAmountFC         edm.Double   `json:"ExtraDutyAmountFC"`         // Extra duty amount in the currency of the transaction. Both extra duty amount and VAT amount need to be specified in order to differ this property from automatically calculated.
	ExtraDutyPercentage       edm.Double   `json:"ExtraDutyPercentage"`       // Extra duty percentage for the item
	From                      edm.DateTime `json:"From,omitempty"`            // From date for deferred revenue
	GLAccount                 edm.GUID     `json:"GLAccount"`                 // The GL Account of the invoice line. This field is generated based on the revenue account of the item (or the related item group). G/L Account is also used to determine whether the costcenter / costunit is mandatory
	GLAccountCode             edm.String   `json:"GLAccountCode"`             // Code of GLAccount
	GLAccountDescription      edm.String   `json:"GLAccountDescription"`      // Description of GLAccount
	IntraStatArea             edm.String   `json:"IntraStatArea"`             // IntraStat area
	IntraStatCountry          edm.String   `json:"IntraStatCountry"`          // IntraStatCountry
	IntraStatDeliveryTerm     edm.String   `json:"IntraStatDeliveryTerm"`     // IntraStat delivery term
	IntraStatTransactionA     edm.String   `json:"IntraStatTransactionA"`     // IntraStat transaction a
	IntraStatTransportMethod  edm.String   `json:"IntraStatTransportMethod"`  // IntraStat transport method
	LineNumber                edm.Int32    `json:"LineNumber"`                // Indicates the sequence of the lines within one entry
	Notes                     edm.String   `json:"Notes"`                     // Extra notes
	Project                   edm.GUID     `json:"Project"`                   // The project to which the sales transaction line is linked. The project can be different per line. Sometimes also the project in the header is filled although this is not really used
	ProjectDescription        edm.String   `json:"ProjectDescription"`        // Description of Project
	Quantity                  edm.Double   `json:"Quantity"`                  // The number of items sold in default units. The quantity shown in the entry screen is Quantity * UnitFactor
	SerialNumber              edm.String   `json:"SerialNumber"`              // Serial number
	StatisticalNetWeight      edm.Double   `json:"StatisticalNetWeight"`      // Statistical NetWeight
	StatisticalNumber         edm.String   `json:"StatisticalNumber"`         // Statistical Number
	StatisticalQuantity       edm.Double   `json:"StatisticalQuantity"`       // Statistical Quantity
	StatisticalValue          edm.Double   `json:"StatisticalValue"`          // Statistical Value
	Subscription              edm.GUID     `json:"Subscription"`              // When generating invoices from subscriptions, this field records the link between invoice lines and subscription lines
	SubscriptionDescription   edm.String   `json:"SubscriptionDescription"`   // Description of Subscription
	TaxSchedule               edm.GUID     `json:"TaxSchedule"`               // Obsolete
	To                        edm.DateTime `json:"To,omitempty"`              // To date for deferred revenue
	TrackingNumber            edm.GUID     `json:"TrackingNumber"`            // Reference to TrackingNumber
	TrackingNumberDescription edm.String   `json:"TrackingNumberDescription"` // Description of TrackingNumber
	Type                      edm.Int32    `json:"Type"`                      // Type: 20 = Sales entry, 21 = Sales credit note
	VATAmountDC               edm.Double   `json:"VATAmountDC"`               // VAT amount in the default currency of the company
	VATAmountFC               edm.Double   `json:"VATAmountFC"`               // VAT amount in the currency of the transaction. Use this property to specify a VAT amount that differs from the VAT amount that is automatically calculated. However if the transaction uses extra duty, extra duty amount also needs to be specified.
	VATBaseAmountDC           edm.Double   `json:"VATBaseAmountDC"`           // The VAT base amount in the default currency of the company. This is calculated based on the VATBaseAmountFC
	VATBaseAmountFC           edm.Double   `json:"VATBaseAmountFC"`           // The VAT base amount in invoice currency. This is calculated with the use of VAT codes. It's an internal value
	VATCode                   edm.String   `json:"VATCode"`                   // The VAT code used when the invoice was registered
	VATCodeDescription        edm.String   `json:"VATCodeDescription"`        // Description of VATCode
	VATPercentage             edm.Double   `json:"VATPercentage"`             // The VAT percentage of the VAT code. This is the percentage at the moment the invoice is created. It's also used by the default calculation of VAT amounts and VAT base amounts
}

func (l SalesEntryLine) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(l)
}

