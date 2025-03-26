package generaljournalentry

import (
	"encoding/json"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/omitempty"
	"github.com/omniboost/go-exactonline/utils"
)

type GeneralJournalEntries []GeneralJournalEntry

type GeneralJournalEntry struct {
	EntryID                  edm.GUID                 `json:"EntryID"`                  // Primary key
	Created                  edm.DateTime             `json:"Created"`                  // Creation date
	Currency                 edm.String               `json:"Currency"`                 // Currency code
	Division                 edm.Int32                `json:"Division"`                 // Division code
	EntryNumber              edm.Int32                `json:"EntryNumber,omitempty"`    // Entry number
	ExchangeRate             edm.Double               `json:"ExchangeRate"`             // Exchange rate
	FinancialPeriod          edm.Int16                `json:"FinancialPeriod"`          // Financial period
	FinancialYear            edm.Int16                `json:"FinancialYear"`            // Financial year
	GeneralJournalEntryLines GeneralJournalEntryLines `json:"GeneralJournalEntryLines"` // Collection of lines
	JournalCode              edm.String               `json:"JournalCode"`              // Code of Journal
	JournalDescription       edm.String               `json:"JournalDescription"`       // Description of Journal
	Modified                 edm.DateTime             `json:"Modified"`                 // Last modified date
	Reversal                 edm.Boolean              `json:"Reversal"`                 // Indicates that amounts are reversed
	Status                   edm.Int16                `json:"Status"`                   // Status: 5 = Rejected, 20 = Open, 50 = Processed
	StatusDescription        edm.String               `json:"StatusDescription"`        // Description of Status
	Type                     edm.Int32                `json:"Type"`                     // Type: 10 = Opening balance, 90 = Other
	TypeDescription          edm.String               `json:"TypeDescription"`          // Description of Type
}

type GeneralJournalEntryLines []GeneralJournalEntryLine

// standalone: "GeneralJournalEntryLines": []
// deferred: "GeneralJournalEntryLines": {"__deferred": {}}
// embedded: "GeneralJournalEntryLines": {"results": []}
func (e *GeneralJournalEntryLines) UnmarshalJSON(data []byte) error {
	type Results GeneralJournalEntryLines

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

		*e = GeneralJournalEntryLines(*results)
		return nil
	}

	// lines are embedded or deferred: only try embedded
	envelope := &Envelope{Results: Results(*e)}
	err = json.Unmarshal(data, envelope)
	if err != nil {
		return err
	}

	*e = GeneralJournalEntryLines(envelope.Results)
	return nil
}

type GeneralJournalEntryLine struct {
	ID                    edm.GUID     `json:"ID"`                    // Primary key
	Account               edm.GUID     `json:"Account"`               // Reference to account
	AccountCode           edm.String   `json:"AccountCode"`           // Code of Account
	AccountName           edm.String   `json:"AccountName"`           // Name of Account
	AmountDC              edm.Double   `json:"AmountDC"`              // Amount in the default currency of the company, If an 'Including' VAT code is used this amount includes the VAT amount.
	AmountFC              edm.Double   `json:"AmountFC"`              // Amount in the currency of the transaction. If an 'Including' VAT code is used this amount includes the VAT amount.
	AmountVATDC           edm.Double   `json:"AmountVATDC"`           // Vat amount in the default currency of the company.
	AmountVATFC           edm.Double   `json:"AmountVATFC"`           // Vat amount in the currency of the transaction. If you want to set this in a POST you have to specify VATCode as well.
	Asset                 edm.GUID     `json:"Asset"`                 // Reference to asset
	AssetCode             edm.String   `json:"AssetCode"`             // Code of Asset
	AssetDescription      edm.String   `json:"AssetDescription"`      // Description of Asset
	CostCenter            edm.String   `json:"CostCenter"`            // Reference to cost center
	CostCenterDescription edm.String   `json:"CostCenterDescription"` // Description of CostCenter
	CostUnit              edm.String   `json:"CostUnit"`              // Reference to cost unit
	CostUnitDescription   edm.String   `json:"CostUnitDescription"`   // Description of CostUnit
	Created               edm.DateTime `json:"Created"`               // Creation date
	Creator               edm.GUID     `json:"Creator"`               // User ID of creator
	CreatorFullName       edm.String   `json:"CreatorFullName"`       // Name of creator
	Date                  edm.DateTime `json:"Date"`                  // Entry date
	Description           edm.String   `json:"Description"`           // Description
	Division              edm.Int32    `json:"Division"`              // Division code
	Document              edm.GUID     `json:"Document"`              // Reference to document
	DocumentNumber        edm.Int32    `json:"DocumentNumber"`        // Document number
	DocumentSubject       edm.String   `json:"DocumentSubject"`       // Document subject
	EntryID               edm.GUID     `json:"EntryID,omitempty"`     // Reference to header of the entry
	EntryNumber           edm.Int32    `json:"EntryNumber,omitempty"` // Entry number of the header
	GLAccount             edm.GUID     `json:"GLAccount"`             // General ledger account
	GLAccountCode         edm.String   `json:"GLAccountCode"`         // Code of GLAccount
	GLAccountDescription  edm.String   `json:"GLAccountDescription"`  // Description of GLAccount
	LineNumber            edm.Int32    `json:"LineNumber"`            // Line number
	Modified              edm.DateTime `json:"Modified"`              // Last modified date
	Modifier              edm.GUID     `json:"Modifier"`              // User ID of modifier
	ModifierFullName      edm.String   `json:"ModifierFullName"`      // Name of modifier
	Notes                 edm.String   `json:"Notes"`                 // Extra remarks
	OffsetID              edm.GUID     `json:"OffsetID"`              //
	OurRef                edm.Int32    `json:"OurRef"`                // Our ref of general journal entry
	Project               edm.GUID     `json:"Project"`               // Reference to project
	ProjectCode           edm.String   `json:"ProjectCode"`           // Code of Project
	ProjectDescription    edm.String   `json:"ProjectDescription"`    // Description of Project
	Quantity              edm.Double   `json:"Quantity"`              // Quantity
	VATBaseAmountDC       edm.Double   `json:"VATBaseAmountDC"`       // VAT base amount in the default currency of the company
	VATBaseAmountFC       edm.Double   `json:"VATBaseAmountFC"`       // VAT base amount in the currency of the transaction
	VATCode               edm.String   `json:"VATCode"`               // VATCode can only be used if the general journal has VAT enabled. VAT Lines will be automatically created if the VATCode is specified when creating a new general journal entry.
	VATCodeDescription    edm.String   `json:"VATCodeDescription"`    // Description of VATCode
	VATPercentage         edm.Double   `json:"VATPercentage"`         // Vat percentage
	VATType               VATType      `json:"VATType"`               // The VAT type determines what the values are in relation to VAT returns. The following values are supported:A Sales VAT to pay,D Credit note extra duty to claim,I Purchase basis,M Credit note purchase non-deductible,N Purchase non-deductible,O Purchase VAT to claim,P Purchase VAT to pay,Q Credit note purchase VAT to claim,R Extra duty to pay,S No VAT,V Sales basis,W Credit note purchase basis,X Credit note sales basis,Y Credit note purchase VAT to pay,Z Credit note sales VAT to claim
}

type NewGeneralJournalEntry struct {
	// EntryID                  edm.GUID                    `json:"EntryID,omitempty"`                  // Primary key
	Currency                 edm.String                  `json:"Currency"`                 // Currency code
	EntryNumber              edm.Int32                   `json:"EntryNumber,omitempty"`    // Entry number
	ExchangeRate             edm.Double                  `json:"ExchangeRate"`             // Exchange rate
	FinancialPeriod          edm.Int16                   `json:"FinancialPeriod"`          // Financial period
	FinancialYear            edm.Int16                   `json:"FinancialYear"`            // Financial year
	GeneralJournalEntryLines NewGeneralJournalEntryLines `json:"GeneralJournalEntryLines"` // Collection of lines
	JournalCode              edm.String                  `json:"JournalCode"`              // Code of Journal
	Reversal                 edm.Boolean                 `json:"Reversal"`                 // Indicates that amounts are reversed}
}

type NewGeneralJournalEntryLine struct {
	ID            edm.GUID     `json:"ID,omitempty"`            // Primary key
	Account       edm.GUID     `json:"Account,omitempty"`       // Reference to account
	AmountFC      edm.Double   `json:"AmountFC,omitempty"`      // Amount in the currency of the transaction. If an 'Including' VAT code is used this amount includes the VAT amount.
	AmountVATFC   *edm.Double  `json:"AmountVATFC,omitempty"`   // Vat amount in the currency of the transaction. If you want to set this in a POST you have to specify VATCode as well.
	Asset         edm.GUID     `json:"Asset,omitempty"`         // Reference to asset
	CostCenter    edm.String   `json:"CostCenter,omitempty"`    // Reference to cost center
	CostUnit      edm.String   `json:"CostUnit,omitempty"`      // Reference to cost unit
	Date          edm.DateTime `json:"Date"`                    // Entry date
	Description   edm.String   `json:"Description"`             // Description
	Document      edm.GUID     `json:"Document,omitempty"`      // Reference to document
	EntryID       edm.GUID     `json:"EntryID,omitempty"`       // Reference to header of the entry
	GLAccount     edm.GUID     `json:"GLAccount"`               // General ledger account
	Notes         edm.String   `json:"Notes"`                   // Extra remarks
	OffsetID      edm.GUID     `json:"OffsetID,omitempty"`      //
	OurRef        edm.Int32    `json:"OurRef,omitempty"`        // Our ref of general journal entry
	Project       edm.GUID     `json:"Project,omitempty"`       // Reference to project
	Quantity      edm.Double   `json:"Quantity,omitempty"`      // Quantity
	VATCode       *edm.String  `json:"VATCode,omitempty"`       // VATCode can only be used if the general journal has VAT enabled. VAT Lines will be automatically created if the VATCode is specified when creating a new general journal entry.
	VATPercentage edm.Double   `json:"VATPercentage,omitempty"` // Vat percentage
}

func (l NewGeneralJournalEntryLine) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(l)
}

type NewGeneralJournalEntryLines []NewGeneralJournalEntryLine

type VATType edm.String
