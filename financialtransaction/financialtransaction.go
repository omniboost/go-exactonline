package financialtransaction

import (
	"encoding/json"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/rest"
	"github.com/omniboost/go-exactonline/utils"
)

func NewService(rest *rest.Client) *Service {
	return &Service{rest: rest}
}

type Service struct {
	rest *rest.Client
}

type Transactions []Transaction

type Transaction struct {
	EntryID                     edm.GUID          `json:"EntryID,omitempty"`
	ClosingBalanceFC            edm.Double        `json:"ClosingBalanceFC"`
	Created                     edm.DateTime      `json:"Created"`
	Date                        edm.DateTime      `json:"Date"`
	Description                 edm.String        `json:"Description"`
	Division                    edm.Int32         `json:"Division"`
	Document                    edm.GUID          `json:"Document"`
	DocumentNumber              edm.Int32         `json:"DocumentNumber"`
	DocumentSubject             edm.String        `json:"DocumentSubject"`
	EntryNumber                 edm.Int32         `json:"EntryNumber"`
	ExternalLinkDescription     edm.String        `json:"ExternalLinkDescription"`
	ExternalLinkReference       edm.String        `json:"ExternalLinkReference"`
	FinancialPeriod             edm.Int16         `json:"FinancialPeriod"`
	FinancialYear               edm.Int16         `json:"FinancialYear"`
	IsExtraDuty                 edm.Boolean       `json:"IsExtraDuty"`
	JournalCode                 edm.String        `json:"JournalCode"`
	JournalDescription          edm.String        `json:"JournalDescription"`
	Modified                    edm.DateTime      `json:"Modified"`
	OpeningBalanceFC            edm.Double        `json:"OpeningBalanceFC"`
	PaymentConditionCode        edm.String        `json:"PaymentConditionCode"`
	PaymentConditionDescription edm.String        `json:"PaymentConditionDescription"`
	PaymentReference            edm.String        `json:"PaymentReference"`
	Status                      TransactionStatus `json:"Status"`
	StatusDescription           edm.String        `json:"StatusDescription"`
	TransactionLines            TransactionLines  `json:"TransactionLines"`
	Type                        TransactionType   `json:"Type"`
	TypeDescription             edm.String        `json:"TypeDescription"`
}

type TransactionLines []TransactionLine

// standalone: "TransactionLines": []
// deferred: "TransactionLines": {"__deferred": {}}
// embedded: "TransactionLines": {"results": []}
func (t *TransactionLines) UnmarshalJSON(data []byte) (err error) {
	type Results TransactionLines

	type Envelope struct {
		Results  Results         `json:"results"`
		Deferred json.RawMessage `json:"__deferred"`
	}

	// create the json tester
	tester := &utils.JsonTester{}
	json.Unmarshal(data, tester)
	if err != nil {
		return err
	}

	// test if json is array (standalone)
	if tester.IsArray() {
		results := &Results{}
		err = json.Unmarshal(data, results)
		if err != nil {
			return err
		}

		*t = TransactionLines(*results)
		return nil
	}

	envelope := &Envelope{Results: Results(*t)}
	err = json.Unmarshal(data, envelope)
	if err != nil {
		return err
	}

	*t = TransactionLines(envelope.Results)
	return nil
}

type TransactionLine struct {
	ID                        edm.GUID     `json:"ID"`
	Account                   edm.GUID     `json:"Account"`
	AccountCode               edm.String   `json:"AccountCode"`
	AccountName               edm.String   `json:"AccountName"`
	AmountDC                  edm.Double   `json:"AmountDC"`
	AmountFC                  edm.Double   `json:"AmountFC"`
	AmountVATBaseFC           edm.Double   `json:"AmountVATBaseFC"`
	AmountVATFC               edm.Double   `json:"AmountVATFC"`
	Asset                     edm.GUID     `json:"Asset"`
	AssetCode                 edm.String   `json:"AssetCode"`
	AssetDescription          edm.String   `json:"AssetDescription"`
	CostCenter                edm.String   `json:"CostCenter"`
	CostCenterDescription     edm.String   `json:"CostCenterDescription"`
	CostUnit                  edm.String   `json:"CostUnit"`
	CostUnitDescription       edm.String   `json:"CostUnitDescription"`
	Created                   edm.DateTime `json:"Created"`
	Creator                   edm.GUID     `json:"Creator"`
	CreatorFullName           edm.String   `json:"CreatorFullName"`
	Currency                  edm.String   `json:"Currency"`
	Date                      edm.DateTime `json:"Date"`
	Description               edm.String   `json:"Description"`
	Division                  edm.Int32    `json:"Division"`
	Document                  edm.GUID     `json:"Document"`
	DocumentNumber            edm.Int32    `json:"DocumentNumber"`
	DocumentSubject           edm.String   `json:"DocumentSubject"`
	DueDate                   edm.DateTime `json:"DueDate"`
	EntryID                   edm.GUID     `json:"EntryID,omitempty"`
	EntryNumber               edm.Int32    `json:"EntryNumber"`
	ExchangeRate              edm.Double   `json:"ExchangeRate"`
	ExtraDutyAmountFC         edm.Double   `json:"ExtraDutyAmountFC"`
	ExtraDutyPercentage       edm.Double   `json:"ExtraDutyPercentage"`
	FinancialPeriod           edm.Int16    `json:"FinancialPeriod"`
	FinancialYear             edm.Int16    `json:"FinancialYear"`
	GLAccount                 edm.GUID     `json:"GLAccount"`
	GLAccountCode             edm.String   `json:"GLAccountCode"`
	GLAccountDescription      edm.String   `json:"GLAccountDescription"`
	InvoiceNumber             edm.Int32    `json:"InvoiceNumber,omitempty"`
	Item                      edm.GUID     `json:"Item"`
	ItemCode                  edm.String   `json:"ItemCode"`
	ItemDescription           edm.String   `json:"ItemDescription"`
	JournalCode               edm.String   `json:"JournalCode"`
	JournalDescription        edm.String   `json:"JournalDescription"`
	LineNumber                edm.Int32    `json:"LineNumber"`
	LineType                  edm.Int16    `json:"LineType"`
	Modified                  edm.DateTime `json:"Modified"`
	Modifier                  edm.GUID     `json:"Modifier"`
	ModifierFullName          edm.String   `json:"ModifierFullName"`
	Notes                     edm.String   `json:"Notes"`
	OffsetID                  edm.GUID     `json:"OffsetID"`
	OrderNumber               edm.Int32    `json:"OrderNumber"`
	PaymentDiscountAmount     edm.Double   `json:"PaymentDiscountAmount"`
	PaymentReference          edm.String   `json:"PaymentReference"`
	Project                   edm.GUID     `json:"Project"`
	ProjectCode               edm.String   `json:"ProjectCode"`
	ProjectDescription        edm.String   `json:"ProjectDescription"`
	Quantity                  edm.Double   `json:"Quantity"`
	SerialNumber              edm.String   `json:"SerialNumber"`
	Status                    edm.Int16    `json:"Status"`
	Subscription              edm.GUID     `json:"Subscription"`
	SubscriptionDescription   edm.String   `json:"SubscriptionDescription"`
	TrackingNumber            edm.String   `json:"TrackingNumber"`
	TrackingNumberDescription edm.String   `json:"TrackingNumberDescription"`
	Type                      edm.Int32    `json:"Type"`
	VATCode                   edm.String   `json:"VATCode"`
	VATCodeDescription        edm.String   `json:"VATCodeDescription"`
	VATPercentage             edm.Double   `json:"VATPercentage"`
	VATType                   edm.String   `json:"VATType"`
	YourRef                   edm.String   `json:"YourRef"`
}

const (
	StatusRejected  TransactionStatus = 5
	StatusOpen      TransactionStatus = 20
	StatusProcessed TransactionStatus = 50
)

type TransactionStatus edm.Int16

const (
	TypeOpeningBalance              TransactionType = 10
	TypeSalesEntry                  TransactionType = 20
	TypeSalesCreditNote             TransactionType = 21
	TypePurchaseEntry               TransactionType = 30
	TypePurchaseCreditNote          TransactionType = 31
	TypeCashFlow                    TransactionType = 40
	TypeVatReturn                   TransactionType = 50
	TypeAssetDepreciation           TransactionType = 70
	TypeAssetInvestment             TransactionType = 71
	TypeAssetRevaluation            TransactionType = 72
	TypeAssetTransfer               TransactionType = 73
	TypeAssetSplit                  TransactionType = 74
	TypeAssetDiscontinue            TransactionType = 75
	TypeAssetSales                  TransactionType = 76
	TypeRevaluation                 TransactionType = 80
	TypeExchangeRateDifference      TransactionType = 82
	TypePaymentDifference           TransactionType = 83
	TypeDeferredRevenue             TransactionType = 84
	TypeTrackingNumberRevaluation   TransactionType = 85
	TypeDeferredCost                TransactionType = 86
	TypeVatOnPrepayment             TransactionType = 87
	TypeOther                       TransactionType = 90
	TypeDelivery                    TransactionType = 120
	TypeSalesReturn                 TransactionType = 121
	TypeReceipt                     TransactionType = 130
	TypePurchaseReturn              TransactionType = 131
	TypeShopOrderStockReceipt       TransactionType = 140
	TypeShopOrderStockReversal      TransactionType = 141
	TypeIssueToParent               TransactionType = 142
	TypeShopOrderTimeEntry          TransactionType = 145
	TypeShopOrderTimeEntryReversal  TransactionType = 146
	TypeShopOrderByProductReceipt   TransactionType = 147
	TypeShopOrderByProductReversal  TransactionType = 148
	TypeRequirementIssue            TransactionType = 150
	TypeRequirementReversal         TransactionType = 151
	TypeReturnedFromParent          TransactionType = 152
	TypeSubcontractIssue            TransactionType = 155
	TypeSubcontractReversal         TransactionType = 156
	TypeShopOrderCompleted          TransactionType = 158
	TypeFinishAssembly              TransactionType = 162
	TypePayroll                     TransactionType = 170
	TypeStockRevaluation            TransactionType = 180
	TypeFinancialRevaluation        TransactionType = 181
	TypeStockCount                  TransactionType = 195
	TypeCorrectionEntry             TransactionType = 290
	TypePeriodClosing               TransactionType = 310
	TypeYearEndReflection           TransactionType = 320
	TypeYearEndCosting              TransactionType = 321
	TypeYearEndProfitsToGrossProfit TransactionType = 322
	TypeYearEndCostsToGrossProfit   TransactionType = 323
	TypeYearEndTax                  TransactionType = 324
	TypeYearEndGrossProfitToNetPL   TransactionType = 325
	TypeYearEndNetPLToBalanceSheet  TransactionType = 326
	TypeYearEndClosingBalance       TransactionType = 327
	TypeYearStartOpeningBalance     TransactionType = 328
	TypeBudget                      TransactionType = 3000
)

type TransactionType edm.Int32

func (t *TransactionType) String() string {
	switch int(*t) {
	case 10:
		return "Opening balance"
	case 20:
		return "Sales entry"
	case 21:
		return "Sales credit note"
	case 30:
		return "Purchase entry"
	case 31:
		return "Purchase credit note"
	case 40:
		return "Cash flow"
	case 50:
		return "VAT return"
	case 70:
		return "Asset - Depreciation"
	case 71:
		return "Asset - Investment"
	case 72:
		return "Asset - Revaluation"
	case 73:
		return "Asset - Transfer"
	case 74:
		return "Asset - Split"
	case 75:
		return "Asset - Discontinue"
	case 76:
		return "Asset - Sales"
	case 80:
		return "Revaluation"
	case 82:
		return "Exchange rate difference"
	case 83:
		return "Payment difference"
	case 84:
		return "Deferred revenue"
	case 85:
		return "Tracking number:Revaluation"
	case 86:
		return "Deferred cost"
	case 87:
		return "VAT on prepayment"
	case 90:
		return "Other"
	case 120:
		return "Delivery"
	case 121:
		return "Sales Return"
	case 130:
		return "Receipt"
	case 131:
		return "Purchase return"
	case 140:
		return "Shop order stock receipt"
	case 141:
		return "Shop order stock reversal"
	case 142:
		return "Issue to parent"
	case 145:
		return "Shop order time entry"
	case 146:
		return "Shop order time entry reversal"
	case 147:
		return "Shop order by-product receipt"
	case 148:
		return "Shop order by-product reversal"
	case 150:
		return "Requirement issue"
	case 151:
		return "Requirement reversal"
	case 152:
		return "Returned from parent"
	case 155:
		return "Subcontract Issue"
	case 156:
		return "Subcontract reversal"
	case 158:
		return "Shop order completed"
	case 162:
		return "Finish assembly"
	case 170:
		return "Payroll"
	case 180:
		return "Stock revaluation"
	case 181:
		return "Financial revaluation"
	case 195:
		return "Stock count"
	case 290:
		return "Correction entry"
	case 310:
		return "Period closing"
	case 320:
		return "Year end reflection"
	case 321:
		return "Year end costing"
	case 322:
		return "Year end profits to gross profit"
	case 323:
		return "Year end costs to gross profit"
	case 324:
		return "Year end tax"
	case 325:
		return "Year end gross profit to net p/l"
	case 326:
		return "Year end net p/l to balance sheet"
	case 327:
		return "Year end closing balance"
	case 328:
		return "Year start opening balance"
	case 3000:
		return "Budget"
	}
	return ""
}

type BankEntries []BankEntry

type BankEntry struct {
	EntryID                      edm.GUID       `json:"EntryID,omitempty"`
	BankEntryLines               BankEntryLines `json:"BankEntryLines"`
	BankStatementDocument        edm.GUID       `json:"BankStatementDocument"`
	BankStatementDocumentNumber  edm.Int32      `json:"BankStatementDocumentNumber"`
	BankStatementDocumentSubject edm.String     `json:"BankStatementDocumentSubject"`
	ClosingBalanceFC             edm.Double     `json:"ClosingBalanceFC"`
	Created                      edm.DateTime   `json:"Created"`
	Currency                     edm.String     `json:"Currency"`
	Division                     edm.Int32      `json:"Division"`
	EntryNumber                  edm.Int32      `json:"EntryNumber"`
	FinancialPeriod              edm.Int16      `json:"FinancialPeriod"`
	FinancialYear                edm.Int16      `json:"FinancialYear"`
	JournalCode                  edm.String     `json:"JournalCode"`
	JournalDescription           edm.String     `json:"JournalDescription"`
	Modified                     edm.DateTime   `json:"Modified"`
	OpeningBalanceFC             edm.Double     `json:"OpeningBalanceFC"`
	Status                       edm.Int16      `json:"Status"`
	StatusDescription            edm.String     `json:"StatusDescription"`
}

type BankEntryLines []BankEntryLine

// standalone: "BankEntryLines": []
// deferred: "BankEntryLines": {"__deferred": {}}
// embedded: "BankEntryLines": {"results": []}
func (b *BankEntryLines) UnmarshalJSON(data []byte) (err error) {
	type Results BankEntryLines

	type Envelope struct {
		Results  Results         `json:"results"`
		Deferred json.RawMessage `json:"__deferred"`
	}

	tester := &utils.JsonTester{}
	json.Unmarshal(data, tester)
	if err != nil {
		return err
	}

	// test if json is array (standalone)
	if tester.IsArray() {
		results := &Results{}
		err = json.Unmarshal(data, results)
		if err != nil {
			return err
		}

		*b = BankEntryLines(*results)
		return nil
	}

	envelope := &Envelope{Results: Results(*b)}

	*b = BankEntryLines(envelope.Results)
	return nil
}

type BankEntryLine struct {
	ID                    edm.GUID     `json:"ID"`
	Account               edm.GUID     `json:"Account"`
	AccountCode           edm.String   `json:"AccountCode"`
	AccountName           edm.String   `json:"AccountName"`
	AmountDC              edm.Double   `json:"AmountDC"`
	AmountFC              edm.Double   `json:"AmountFC"`
	AmountVATFC           edm.Double   `json:"AmountVATFC"`
	Asset                 edm.GUID     `json:"Asset"`
	AssetCode             edm.String   `json:"AssetCode"`
	AssetDescription      edm.String   `json:"AssetDescription"`
	CostCenter            edm.String   `json:"CostCenter"`
	CostCenterDescription edm.String   `json:"CostCenterDescription"`
	CostUnit              edm.String   `json:"CostUnit"`
	CostUnitDescription   edm.String   `json:"CostUnitDescription"`
	Created               edm.DateTime `json:"Created"`
	Creator               edm.GUID     `json:"Creator"`
	CreatorFullName       edm.String   `json:"CreatorFullName"`
	Date                  edm.DateTime `json:"Date"`
	Description           edm.String   `json:"Description"`
	Division              edm.Int32    `json:"Division"`
	Document              edm.GUID     `json:"Document"`
	DocumentNumber        edm.Int32    `json:"DocumentNumber"`
	DocumentSubject       edm.String   `json:"DocumentSubject"`
	EntryID               edm.GUID     `json:"EntryID,omitempty"`
	EntryNumber           edm.Int32    `json:"EntryNumber,omitempty"`
	ExchangeRate          edm.Double   `json:"ExchangeRate"`
	GLAccount             edm.GUID     `json:"GLAccount"`
	GLAccountCode         edm.String   `json:"GLAccountCode"`
	GLAccountDescription  edm.String   `json:"GLAccountDescription"`
	LineNumber            edm.Int32    `json:"LineNumber"`
	Modified              edm.DateTime `json:"Modified"`
	Modifier              edm.GUID     `json:"Modifier"`
	ModifierFullName      edm.String   `json:"ModifierFullName"`
	Notes                 edm.String   `json:"Notes"`
	OffsetID              edm.GUID     `json:"OffsetID"`
	OurRef                edm.Int32    `json:"OurRef,omitempty"`
	Project               edm.GUID     `json:"Project"`
	ProjectCode           edm.String   `json:"ProjectCode"`
	ProjectDescription    edm.String   `json:"ProjectDescription"`
	Quantity              edm.Double   `json:"Quantity"`
	VATCode               edm.String   `json:"VATCode"`
	VATCodeDescription    edm.String   `json:"VATCodeDescription"`
	VATPercentage         edm.Double   `json:"VATPercentage"`
	VATType               edm.String   `json:"VATType"`
}

type CashEntries []CashEntry

type CashEntry struct {
	EntryID            edm.GUID       `json:"EntryID"`            // Primary key
	CashEntryLines     CashEntryLines `json:"CashEntryLines"`     // Collection of lines
	ClosingBalanceFC   edm.Double     `json:"ClosingBalanceFC"`   // Closing balance in the currency of the transaction
	Created            edm.DateTime   `json:"Created"`            // Creation date
	Currency           edm.String     `json:"Currency"`           // Currency code
	Division           edm.Int32      `json:"Division"`           // Division code
	EntryNumber        edm.Int32      `json:"EntryNumber"`        // Entry number
	FinancialPeriod    edm.Int16      `json:"FinancialPeriod"`    // Fiancial period
	FinancialYear      edm.Int16      `json:"FinancialYear"`      // Fiancial year
	JournalCode        edm.String     `json:"JournalCode"`        // Code of Journal
	JournalDescription edm.String     `json:"JournalDescription"` // Description of Journal
	Modified           edm.DateTime   `json:"Modified"`           // Last modified date
	OpeningBalanceFC   edm.Double     `json:"OpeningBalanceFC"`   // Opening balance in the currency of the transaction
	Status             edm.Int16      `json:"Status"`             // Status: 20 = Open, 50 = Processed
	StatusDescription  edm.String     `json:"StatusDescription"`  // Description of Status}
}

type CashEntryLines []CashEntryLine

// standalone: "CashEntryLines": []
// deferred: "CashEntryLines": {"__deferred": {}}
// embedded: "CashEntryLines": {"results": []}
func (b *CashEntryLines) UnmarshalJSON(data []byte) (err error) {
	type Results CashEntryLines

	type Envelope struct {
		Results  Results         `json:"results"`
		Deferred json.RawMessage `json:"__deferred"`
	}

	tester := &utils.JsonTester{}
	json.Unmarshal(data, tester)
	if err != nil {
		return err
	}

	// test if json is array (standalone)
	if tester.IsArray() {
		results := &Results{}
		err = json.Unmarshal(data, results)
		if err != nil {
			return err
		}

		*b = CashEntryLines(*results)
		return nil
	}

	envelope := &Envelope{Results: Results(*b)}

	*b = CashEntryLines(envelope.Results)
	return nil
}

type CashEntryLine struct {
	ID                    edm.GUID     `json:"ID"`                    // Primary key
	Account               edm.GUID     `json:"Account"`               // Reference to Account
	AccountCode           edm.String   `json:"AccountCode"`           // Code of Account
	AccountName           edm.String   `json:"AccountName"`           // Name of Account
	AmountDC              edm.Double   `json:"AmountDC"`              // Amount in the default currency of the company
	AmountFC              edm.Double   `json:"AmountFC"`              // Amount in the currency of the transaction
	AmountVATFC           edm.Double   `json:"AmountVATFC"`           // Vat amount in the currency of the transaction
	Asset                 edm.GUID     `json:"Asset"`                 // Reference to an asset
	AssetCode             edm.String   `json:"AssetCode"`             // Code of Asset
	AssetDescription      edm.String   `json:"AssetDescription"`      // Description of Asset
	CostCenter            edm.String   `json:"CostCenter"`            // Reference to a cost center
	CostCenterDescription edm.String   `json:"CostCenterDescription"` // Description of CostCenter
	CostUnit              edm.String   `json:"CostUnit"`              // Reference to a cost unit
	CostUnitDescription   edm.String   `json:"CostUnitDescription"`   // Description of CostUnit
	Created               edm.DateTime `json:"Created"`               // Creation date
	Creator               edm.GUID     `json:"Creator"`               // User ID of creator
	CreatorFullName       edm.String   `json:"CreatorFullName"`       // Name of creator
	Date                  edm.DateTime `json:"Date"`                  // Date
	Description           edm.String   `json:"Description"`           // Description
	Division              edm.Int32    `json:"Division"`              // Division code
	Document              edm.GUID     `json:"Document"`              // Reference to a document
	DocumentNumber        edm.Int32    `json:"DocumentNumber"`        // Number of Document
	DocumentSubject       edm.String   `json:"DocumentSubject"`       // Subject of Document
	EntryID               edm.GUID     `json:"EntryID"`               // Reference to the header
	EntryNumber           edm.Int32    `json:"EntryNumber"`           // Entry number of the header
	ExchangeRate          edm.Double   `json:"ExchangeRate"`          // Exchange rate
	GLAccount             edm.GUID     `json:"GLAccount"`             // General ledger account
	GLAccountCode         edm.String   `json:"GLAccountCode"`         // Code of GLAccount
	GLAccountDescription  edm.String   `json:"GLAccountDescription"`  // Description of GLAccount
	LineNumber            edm.Int32    `json:"LineNumber"`            // Line number
	Modified              edm.DateTime `json:"Modified"`              // Last modified date
	Modifier              edm.GUID     `json:"Modifier"`              // User ID of modifier
	ModifierFullName      edm.String   `json:"ModifierFullName"`      // Name of modifier
	Notes                 edm.String   `json:"Notes"`                 // Extra remarks
	OffsetID              edm.GUID     `json:"OffsetID"`              // Reference to offset line
	OurRef                edm.Int32    `json:"OurRef"`                // Invoice number
	Project               edm.GUID     `json:"Project"`               // Reference to a project
	ProjectCode           edm.String   `json:"ProjectCode"`           // Code of Project
	ProjectDescription    edm.String   `json:"ProjectDescription"`    // Description of Project
	Quantity              edm.Double   `json:"Quantity"`              // Quantity
	VATCode               edm.String   `json:"VATCode"`               // Reference to vat code
	VATCodeDescription    edm.String   `json:"VATCodeDescription"`    // Description of VATCode
	VATPercentage         edm.Double   `json:"VATPercentage"`         // Vat code percentage
	VATType               edm.String   `json:"VATType"`               // Type of vat code}
}
