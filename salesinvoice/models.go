package salesinvoice

import (
	"encoding/json"
	"errors"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/utils"
)

type SalesInvoices []SalesInvoice

type SalesInvoice struct {
	InvoiceID                            edm.GUID          `json:"InvoiceID"`                            // Primary key
	AmountDC                             edm.Double        `json:"AmountDC"`                             // For the header lines (LineNumber = 0) of an entry this is the SUM(AmountDC) of all lines
	AmountFC                             edm.Double        `json:"AmountFC"`                             // For the header this is the sum of all lines, including VAT
	Created                              edm.DateTime      `json:"Created"`                              // Creation date
	Creator                              edm.GUID          `json:"Creator"`                              // User ID of creator
	CreatorFullName                      edm.String        `json:"CreatorFullName"`                      // Name of creator
	Currency                             edm.String        `json:"Currency"`                             // Currency for the invoice. Default this is the currency of the administration
	DeliverTo                            edm.GUID          `json:"DeliverTo"`                            // Delivery account for invoice
	DeliverToAddress                     edm.GUID          `json:"DeliverToAddress"`                     // Address of delivery as per invoice delivery account
	DeliverToContactPerson               edm.GUID          `json:"DeliverToContactPerson"`               // Delivery account person for invoice
	DeliverToContactPersonFullName       edm.String        `json:"DeliverToContactPersonFullName"`       // Name of delivery account's contact person as per invoice
	DeliverToName                        edm.String        `json:"DeliverToName"`                        // Name of the delivery account's customer as per invoice
	Description                          edm.String        `json:"Description"`                          // Description. Can be different for header and lines
	Division                             edm.Int32         `json:"Division"`                             // Division code
	Document                             edm.GUID          `json:"Document"`                             // Document that is manually linked to the invoice
	DocumentNumber                       edm.Int32         `json:"DocumentNumber"`                       // Number of the document
	DocumentSubject                      edm.String        `json:"DocumentSubject"`                      // Subject of the document
	DueDate                              edm.DateTime      `json:"DueDate"`                              // The due date for payments. This date is calculated based on the EntryDate and the Paymentcondition
	ExtraDutyAmountFC                    edm.Double        `json:"ExtraDutyAmountFC"`                    // Extra duty amount in the currency of the transaction. Both extra duty amount and VAT amount need to be specified in order to differ this property from automatically calculated.
	InvoiceDate                          edm.DateTime      `json:"InvoiceDate"`                          // Official date for the invoice. When the invoice is entered it's equal to the field 'EntryDate'. During the printing process the invoice date can be entered
	InvoiceNumber                        edm.Int32         `json:"InvoiceNumber"`                        // Assigned at entry or at printing depending on setting. The number assigned is based on the freenumbers as defined for the Journal. When printing the field InvoiceNumber is copied to the fields EntryNumber and InvoiceNumber of the sales entry
	InvoiceTo                            edm.GUID          `json:"InvoiceTo"`                            // Reference to the Customer who will receive the invoice
	InvoiceToContactPerson               edm.GUID          `json:"InvoiceToContactPerson"`               // Reference to the Contact person of the customer who will receive the invoice
	InvoiceToContactPersonFullName       edm.String        `json:"InvoiceToContactPersonFullName"`       // Name of the contact person of the customer who will receive the invoice
	InvoiceToName                        edm.String        `json:"InvoiceToName"`                        // Name of the customer who will receive the invoice
	IsExtraDuty                          edm.Boolean       `json:"IsExtraDuty"`                          // Indicates whether the invoice has extra duty
	Journal                              edm.String        `json:"Journal"`                              // The journal code. Every invoice should be linked to a sales journal
	JournalDescription                   edm.String        `json:"JournalDescription"`                   // Description of Journal
	Modified                             edm.DateTime      `json:"Modified"`                             // Last modified date
	Modifier                             edm.GUID          `json:"Modifier"`                             // User ID of modifier
	ModifierFullName                     edm.String        `json:"ModifierFullName"`                     // Name of modifier
	OrderDate                            edm.DateTime      `json:"OrderDate"`                            // Order date
	OrderedBy                            edm.GUID          `json:"OrderedBy"`                            // Customer who ordered the invoice
	OrderedByContactPerson               edm.GUID          `json:"OrderedByContactPerson"`               // Contact person of customer who ordered the invoice
	OrderedByContactPersonFullName       edm.String        `json:"OrderedByContactPersonFullName"`       // Name of contact person of customer who ordered the invoice
	OrderedByName                        edm.String        `json:"OrderedByName"`                        // Name of customer who ordered the invoice
	OrderNumber                          edm.Int32         `json:"OrderNumber"`                          // Number to identify the order. By default the number is based on a setting for the first free number, but you can post your own number.
	PaymentCondition                     edm.String        `json:"PaymentCondition"`                     // The payment condition used for due date and discount calculation
	PaymentConditionDescription          edm.String        `json:"PaymentConditionDescription"`          // Description of PaymentCondition
	PaymentReference                     edm.String        `json:"PaymentReference"`                     // Payment reference for sales invoice
	Remarks                              edm.String        `json:"Remarks"`                              // Extra remarks
	SalesInvoiceLines                    SalesInvoiceLines `json:"SalesInvoiceLines"`                    // Collection of lines
	Salesperson                          edm.GUID          `json:"Salesperson"`                          // Sales representative
	SalespersonFullName                  edm.String        `json:"SalespersonFullName"`                  // Name of sales representative
	StarterSalesInvoiceStatus            edm.Int16         `json:"StarterSalesInvoiceStatus"`            // Starter Sales invoice status (for starter functionality)
	StarterSalesInvoiceStatusDescription edm.String        `json:"StarterSalesInvoiceStatusDescription"` // Description of StarterSalesInvoiceStatus
	Status                               edm.Int16         `json:"Status"`                               // The status of the entry. 10 = draft. During the creation of an invoice draft records occur in the draft modus if during an invoice a new page with lines is triggered. If the user leaves the invoice in an abnormal way the draft invoices can be recovered. Draft invoices are not included in financial reports, balances etc. 20 = open. Open invoices can be changed. New invoices get the status open by default. 50 = processed. Processed invoices can't be changed anymore. Processing is done via printing. Processed invoices can't be reopened
	StatusDescription                    edm.String        `json:"StatusDescription"`                    // Description of Status
	TaxSchedule                          edm.GUID          `json:"TaxSchedule"`                          // Obsolete
	TaxScheduleCode                      edm.String        `json:"TaxScheduleCode"`                      // Obsolete
	TaxScheduleDescription               edm.String        `json:"TaxScheduleDescription"`               // Obsolete
	Type                                 InvoiceType       `json:"Type"`                                 // Indicates the type of invoice Values: 8020 - Sales invoices, 8021 - Sales credit note
	TypeDescription                      edm.String        `json:"TypeDescription"`                      // Description of the type
	VATAmountDC                          edm.Double        `json:"VATAmountDC"`                          // Total VAT amount in the default currency of the company
	VATAmountFC                          edm.Double        `json:"VATAmountFC"`                          // Total VAT amount in the currency of the transaction
	WithholdingTaxAmountFC               edm.Double        `json:"WithholdingTaxAmountFC"`               // Withholding tax amount applied to sales invoice
	WithholdingTaxBaseAmount             edm.Double        `json:"WithholdingTaxBaseAmount"`             // Withholding tax base amount to calculate withholding amount
	WithholdingTaxPercentage             edm.Double        `json:"WithholdingTaxPercentage"`             // Withholding tax percentage applied to sales invoice
	YourRef                              edm.String        `json:"YourRef"`                              // The invoice number of the customer
}

type NewSalesInvoice struct {
	InvoiceID                edm.GUID             `json:"InvoiceID"`                // Primary key
	Currency                 edm.String           `json:"Currency"`                 // Currency for the invoice. Default this is the currency of the administration
	DeliverTo                edm.GUID             `json:"DeliverTo"`                // Delivery account for invoice
	DeliverToAddress         edm.GUID             `json:"DeliverToAddress"`         // Address of delivery as per invoice delivery account
	DeliverToContactPerson   edm.GUID             `json:"DeliverToContactPerson"`   // Delivery account person for invoice
	Description              edm.String           `json:"Description"`              // Description. Can be different for header and lines
	Document                 edm.GUID             `json:"Document"`                 // Document that is manually linked to the invoice
	DueDate                  edm.DateTime         `json:"DueDate"`                  // The due date for payments. This date is calculated based on the EntryDate and the Paymentcondition
	ExtraDutyAmountFC        edm.Double           `json:"ExtraDutyAmountFC"`        // Extra duty amount in the currency of the transaction. Both extra duty amount and VAT amount need to be specified in order to differ this property from automatically calculated.
	InvoiceTo                edm.GUID             `json:"InvoiceTo"`                // Reference to the Customer who will receive the invoice
	InvoiceToContactPerson   edm.GUID             `json:"InvoiceToContactPerson"`   // Reference to the Contact person of the customer who will receive the invoice
	IsExtraDuty              edm.Boolean          `json:"IsExtraDuty"`              // Indicates whether the invoice has extra duty
	Journal                  edm.String           `json:"Journal"`                  // The journal code. Every invoice should be linked to a sales journal
	OrderDate                edm.DateTime         `json:"OrderDate"`                // Order date
	OrderedBy                edm.GUID             `json:"OrderedBy"`                // Customer who ordered the invoice
	OrderedByContactPerson   edm.GUID             `json:"OrderedByContactPerson"`   // Contact person of customer who ordered the invoice
	OrderNumber              edm.Int32            `json:"OrderNumber"`              // Number to identify the order. By default the number is based on a setting for the first free number, but you can post your own number.
	PaymentCondition         edm.String           `json:"PaymentCondition"`         // The payment condition used for due date and discount calculation
	PaymentReference         edm.String           `json:"PaymentReference"`         // Payment reference for sales invoice
	Remarks                  edm.String           `json:"Remarks"`                  // Extra remarks
	SalesInvoiceLines        NewSalesInvoiceLines `json:"SalesInvoiceLines"`        // Collection of lines
	Salesperson              edm.GUID             `json:"Salesperson"`              // Sales representative
	TaxSchedule              edm.GUID             `json:"TaxSchedule"`              // Obsolete
	Type                     InvoiceType          `json:"Type"`                     // Indicates the type of invoice Values: 8020 - Sales invoices, 8021 - Sales credit note
	WithholdingTaxAmountFC   edm.Double           `json:"WithholdingTaxAmountFC"`   // Withholding tax amount applied to sales invoice
	WithholdingTaxBaseAmount edm.Double           `json:"WithholdingTaxBaseAmount"` // Withholding tax base amount to calculate withholding amount
	WithholdingTaxPercentage edm.Double           `json:"WithholdingTaxPercentage"` // Withholding tax percentage applied to sales invoice
	YourRef                  edm.String           `json:"YourRef"`                  // The invoice number of the customer
}

func (i *SalesInvoice) Validate() error {

	if i.Journal == "" {
		return errors.New("Journal is a required field")
	}

	if i.OrderedBy.String() == "" {
		return errors.New("OrderedBy is a required field")
	}

	return nil
}

type SalesInvoiceLines []SalesInvoiceLine

// standalone: "SalesInvoiceLines": []
// deferred: "SalesInvoiceLines": {"__deferred": {}}
// embedded: "SalesInvoiceLines": {"results": []}
func (i *SalesInvoiceLines) UnmarshalJSON(data []byte) error {
	type Results SalesInvoiceLines

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

		*i = SalesInvoiceLines(*results)
		return nil
	}

	// lines are embedded or deferred: only try embedded
	envelope := &Envelope{Results: Results(*i)}
	err = json.Unmarshal(data, envelope)
	if err != nil {
		return err
	}

	*i = SalesInvoiceLines(envelope.Results)
	return nil
}

type SalesInvoiceLine struct {
	ID                      edm.GUID     `json:"ID"`
	AmountDC                edm.Double   `json:"AmountDC"`
	AmountFC                edm.Double   `json:"AmountFC"`
	CostCenter              edm.String   `json:"CostCenter"`
	CostCenterDescription   edm.String   `json:"CostCenterDescription"`
	CostUnit                edm.String   `json:"CostUnit"`
	CostUnitDescription     edm.String   `json:"CostUnitDescription"`
	DeliveryDate            edm.DateTime `json:"DeliveryDate"`
	Description             edm.String   `json:"Description"`
	Discount                edm.Double   `json:"Discount"`
	Division                edm.Int32    `json:"Division"`
	Employee                edm.GUID     `json:"Employee"`
	EmployeeFullName        edm.String   `json:"EmployeeFullName"`
	EndTime                 edm.DateTime `json:"EndTime"`
	ExtraDutyAmountFC       edm.Double   `json:"ExtraDutyAmountFC"`
	ExtraDutyPercentage     edm.Double   `json:"ExtraDutyPercentage"`
	GLAccount               edm.GUID     `json:"GLAccount"`
	GLAccountDescription    edm.String   `json:"GLAccountDescription"`
	InvoiceID               edm.GUID     `json:"InvoiceID"`
	Item                    edm.GUID     `json:"Item"`
	ItemCode                edm.String   `json:"ItemCode"`
	ItemDescription         edm.String   `json:"ItemDescription"`
	LineNumber              edm.Int32    `json:"LineNumber"`
	NetPrice                edm.Double   `json:"NetPrice"`
	Notes                   edm.String   `json:"Notes"`
	Pricelist               edm.GUID     `json:"Pricelist"`
	PricelistDescription    edm.String   `json:"PricelistDescription"`
	Project                 edm.GUID     `json:"Project"`
	ProjectDescription      edm.String   `json:"ProjectDescription"`
	ProjectWBS              edm.GUID     `json:"ProjectWBS"`
	ProjectWBSDescription   edm.String   `json:"ProjectWBSDescription"`
	Quantity                edm.Double   `json:"Quantity"`
	SalesOrder              edm.GUID     `json:"SalesOrder"`
	SalesOrderLine          edm.GUID     `json:"SalesOrderLine"`
	SalesOrderLineNumber    edm.Int32    `json:"SalesOrderLineNumber"`
	SalesOrderNumber        edm.Int32    `json:"SalesOrderNumber"`
	StartTime               edm.DateTime `json:"StartTime"`
	Subscription            edm.GUID     `json:"Subscription"`
	SubscriptionDescription edm.String   `json:"SubscriptionDescription"`
	TaxSchedule             edm.GUID     `json:"TaxSchedule"`
	TaxScheduleCode         edm.String   `json:"TaxScheduleCode"`
	TaxScheduleDescription  edm.String   `json:"TaxScheduleDescription"`
	UnitCode                edm.String   `json:"UnitCode"`
	UnitDescription         edm.String   `json:"UnitDescription"`
	UnitPrice               edm.Double   `json:"UnitPrice"`
	VATAmountDC             edm.Double   `json:"VATAmountDC"`
	VATAmountFC             edm.Double   `json:"VATAmountFC"`
	VATCode                 edm.String   `json:"VATCode"`
	VATCodeDescription      edm.String   `json:"VATCodeDescription"`
	VATPercentage           edm.Double   `json:"VATPercentage"`
}

type NewSalesInvoiceLines []NewSalesInvoiceLine

type NewSalesInvoiceLine struct {
	ID                  edm.GUID     `json:"ID"`                  // Primary key
	AmountFC            edm.Double   `json:"AmountFC"`            // For normal lines it's the amount excluding VAT
	CostCenter          edm.String   `json:"CostCenter"`          // Reference to Cost center
	CostUnit            edm.String   `json:"CostUnit"`            // Reference to Cost unit
	DeliveryDate        edm.DateTime `json:"DeliveryDate"`        // Delivery date of an item in a sales invoice. This is used for VAT on prepayments, only if sales order is not used in the license.
	Description         edm.String   `json:"Description"`         // Description. Can be different for header and lines
	Discount            edm.Double   `json:"Discount"`            // Discount given on the default price. Discount = (DefaultPrice of Item - PriceItem in line) / DefaultPrice of Item
	Employee            edm.GUID     `json:"Employee"`            // Link to Employee originating from time and cost transactions
	EndTime             edm.DateTime `json:"EndTime"`             // EndTime is used to store the last date of a period. EndTime is used in combination with StartTime
	ExtraDutyAmountFC   edm.Double   `json:"ExtraDutyAmountFC"`   // Extra duty amount in the currency of the transaction. Both extra duty amount and VAT amount need to be specified in order to differ this property from automatically calculated.
	ExtraDutyPercentage edm.Double   `json:"ExtraDutyPercentage"` // Extra duty percentage
	GLAccount           edm.GUID     `json:"GLAccount"`           // The GL Account of the sales invoice line. This field is mandatory. This field is generated based on the revenue account of the item (or the related item group). G/L Account is also used to determine whether the costcenter / costunit is mandatory
	InvoiceID           edm.GUID     `json:"InvoiceID"`           // The InvoiceID identifies the sales invoice. All the lines of a sales invoice have the same InvoiceID
	Item                edm.GUID     `json:"Item"`                // Reference to the item that is sold in this sales invoice line
	NetPrice            edm.Double   `json:"NetPrice"`            // Net price of the sales invoice line
	Notes               edm.String   `json:"Notes"`               // Extra notes
	Pricelist           edm.GUID     `json:"Pricelist"`           // Price list
	Project             edm.GUID     `json:"Project"`             // The project to which the sales transaction line is linked. The project can be different per line. Sometimes also the project in the header is filled although this is not really used
	ProjectWBS          edm.GUID     `json:"ProjectWBS"`          // WBS linked to the sales invoice
	Quantity            edm.Double   `json:"Quantity"`            // The number of items sold in default units. The quantity shown in the entry screen is Quantity * UnitFactor
	StartTime           edm.DateTime `json:"StartTime"`           // StartTime is used to store the first date of a period. StartTime is used in combination with EndTime
	TaxSchedule         edm.GUID     `json:"TaxSchedule"`         // Obsolete
	UnitCode            edm.String   `json:"UnitCode"`            // Code of Unit
	UnitPrice           edm.Double   `json:"UnitPrice"`           // Price per unit
	VATAmountDC         edm.Double   `json:"VATAmountDC"`         // VAT amount in the default currency of the company
	VATAmountFC         edm.Double   `json:"VATAmountFC"`         // VAT amount in the currency of the transaction
	VATCode             edm.String   `json:"VATCode"`             // The VAT code that is used when the invoice is registered
	VATPercentage       edm.Double   `json:"VATPercentage"`       // The vat percentage of the VAT code. This is the percentage at the moment the invoice is created. It's also used for the default calculation of VAT amounts and VAT base amounts
}
