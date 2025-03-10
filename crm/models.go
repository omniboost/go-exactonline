package crm

import (
	"encoding/json"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/utils"
)

type Accounts []Account

type Account struct {
	ID                                  edm.GUID        `json:"ID"`                                          // Primary key
	Accountant                          edm.GUID        `json:"Accountant"`                                  // Reference to the accountant of the customer. Conditions: The referred accountant must have value > 0 in the field IsAccountant
	AccountManager                      edm.GUID        `json:"AccountManager"`                              // ID of the account manager
	AccountManagerFullName              edm.String      `json:"AccountManagerFullName"`                      // Name of the account manager
	AccountManagerHID                   edm.Int32       `json:"AccountManagerHID"`                           // Number of the account manager
	ActivitySector                      edm.GUID        `json:"ActivitySector"`                              // Reference to Activity sector of the account
	ActivitySubSector                   edm.GUID        `json:"ActivitySubSector"`                           // Reference to Activity sub-sector of the account
	AddressLine1                        edm.String      `json:"AddressLine1"`                                // Visit address first line
	AddressLine2                        edm.String      `json:"AddressLine2"`                                // Visit address second line
	AddressLine3                        edm.String      `json:"AddressLine3"`                                // Visit address third line
	BankAccounts                        BankAccounts    `json:"BankAccounts,omitempty" http_methods:"-POST"` // Collection of Bank accounts
	Blocked                             edm.Boolean     `json:"Blocked"`                                     // Indicates if the account is blocked
	BRIN                                edm.GUID        `json:"BRIN"`                                        // Obsolete
	BusinessType                        edm.GUID        `json:"BusinessType"`                                // Reference to the business type of the account
	CanDropShip                         edm.Boolean     `json:"CanDropShip"`                                 // Indicates the default for the possibility to drop ship when an item is linked to a supplier
	ChamberOfCommerce                   edm.String      `json:"ChamberOfCommerce"`                           // Chamber of commerce number
	City                                edm.String      `json:"City"`                                        // Visit address City
	Classification                      edm.String      `json:"Classification"`                              // Obsolete
	Classification1                     edm.GUID        `json:"Classification1"`                             // Account classification 1
	Classification2                     edm.GUID        `json:"Classification2"`                             // Account classification 2
	Classification3                     edm.GUID        `json:"Classification3"`                             // Account classification 3
	Classification4                     edm.GUID        `json:"Classification4"`                             // Account classification 4
	Classification5                     edm.GUID        `json:"Classification5"`                             // Account classification 5
	Classification6                     edm.GUID        `json:"Classification6"`                             // Account classification 6
	Classification7                     edm.GUID        `json:"Classification7"`                             // Account classification 7
	Classification8                     edm.GUID        `json:"Classification8"`                             // Account classification 8
	ClassificationDescription           edm.String      `json:"ClassificationDescription"`                   // Obsolete
	Code                                edm.String      `json:"Code,omitempty"`                              // Unique key, fixed length numeric string with leading spaces, length 18. IMPORTANT: When you use OData $filter on this field you have to make sure the filter parameter contains the leading spaces
	CodeAtSupplier                      edm.String      `json:"CodeAtSupplier"`                              // Code under which your own company is known at the account
	CompanySize                         edm.GUID        `json:"CompanySize"`                                 // Reference to Company size of the account
	ConsolidationScenario               edm.Byte        `json:"ConsolidationScenario"`                       // Consolidation scenario (Time & Billing). Values: 0 = No consolidation, 1 = Item, 2 = Item + Project, 3 = Item + Employee, 4 = Item + Employee + Project, 5 = Project + WBS + Item, 6 = Project + WBS + Item + Employee. Item means in this case including Unit and Price, these also have to be the same to consolidate
	ControlledDate                      edm.DateTime    `json:"ControlledDate"`                              // Date of the latest control of account data with external web service
	Costcenter                          edm.String      `json:"Costcenter"`                                  // Obsolete
	CostcenterDescription               edm.String      `json:"CostcenterDescription"`                       // Obsolete
	CostPaid                            edm.Byte        `json:"CostPaid"`                                    // Obsolete
	Country                             edm.String      `json:"Country"`                                     // Country code
	CountryName                         edm.String      `json:"CountryName"`                                 // Country name
	Created                             edm.DateTime    `json:"Created"`                                     // Creation date
	Creator                             edm.GUID        `json:"Creator"`                                     // User ID of creator
	CreatorFullName                     edm.String      `json:"CreatorFullName"`                             // Name of creator
	CreditLinePurchase                  edm.Double      `json:"CreditLinePurchase"`                          // Maximum amount of credit for Purchase. If no value has been defined, there is no credit limit
	CreditLineSales                     edm.Double      `json:"CreditLineSales"`                             // Maximum amount of credit for sales. If no value has been defined, there is no credit limit
	Currency                            edm.String      `json:"Currency"`                                    // Obsolete
	CustomerSince                       edm.DateTime    `json:"CustomerSince"`                               // Obsolete
	DatevCreditorCode                   edm.String      `json:"DatevCreditorCode"`                           // DATEV creditor code for Germany legislation
	DatevDebtorCode                     edm.String      `json:"DatevDebtorCode"`                             // DATEV debtor code for Germany legislation
	DiscountPurchase                    edm.Double      `json:"DiscountPurchase"`                            // Default discount percentage for purchase. This is stored as a fraction. ie 5.5% is stored as .055
	DiscountSales                       edm.Double      `json:"DiscountSales"`                               // Default discount percentage for sales. This is stored as a fraction. ie 5.5% is stored as .055
	Division                            edm.Int32       `json:"Division"`                                    // Division code
	Document                            edm.GUID        `json:"Document"`                                    // Obsolete
	DunsNumber                          edm.String      `json:"DunsNumber"`                                  // Obsolete
	Email                               edm.String      `json:"Email"`                                       // E-Mail address of the account
	EndDate                             edm.DateTime    `json:"EndDate"`                                     // Determines in combination with the start date if the account is active. If the current date is > end date the account is inactive
	EstablishedDate                     edm.DateTime    `json:"EstablishedDate"`                             // RegistrationDate
	Fax                                 edm.String      `json:"Fax"`                                         // Fax number
	GLAccountPurchase                   edm.GUID        `json:"GLAccountPurchase"`                           // Default (corporate) GL offset account for purchase (cost)
	GLAccountSales                      edm.GUID        `json:"GLAccountSales"`                              // Default (corporate) GL offset account for sales (revenue)
	GLAP                                edm.GUID        `json:"GLAP"`                                        // Default GL account for Accounts Payable
	GLAR                                edm.GUID        `json:"GLAR"`                                        // Default GL account for Accounts Receivable
	HasWithholdingTaxSales              edm.Boolean     `json:"HasWithholdingTaxSales"`                      // Indicates whether a customer has withholding tax on sales
	IgnoreDatevWarningMessage           edm.Boolean     `json:"IgnoreDatevWarningMessage"`                   // Suppressed warning message when there is duplication on the DATEV code
	IntraStatArea                       edm.String      `json:"IntraStatArea"`                               // Intrastat Area
	IntraStatDeliveryTerm               edm.String      `json:"IntraStatDeliveryTerm"`                       // Intrastat delivery method
	IntraStatSystem                     edm.String      `json:"IntraStatSystem"`                             // System for Intrastat
	IntraStatTransactionA               edm.String      `json:"IntraStatTransactionA"`                       // Transaction type A for Intrastat
	IntraStatTransactionB               edm.String      `json:"IntraStatTransactionB"`                       // Transaction type B for Intrastat
	IntraStatTransportMethod            edm.String      `json:"IntraStatTransportMethod"`                    // Transport method for Intrastat
	InvoiceAccount                      edm.GUID        `json:"InvoiceAccount"`                              // ID of account to be invoiced instead of this account
	InvoiceAccountCode                  edm.String      `json:"InvoiceAccountCode"`                          // Code of InvoiceAccount
	InvoiceAccountName                  edm.String      `json:"InvoiceAccountName"`                          // Name of InvoiceAccount
	InvoiceAttachmentType               edm.Int32       `json:"InvoiceAttachmentType"`                       // Indicates which attachment types should be sent when a sales invoice is printed. Only values in related table with Invoice=1 are allowed
	InvoicingMethod                     InvoicingMethod `json:"InvoicingMethod"`                             // Method of sending for sales invoices. Values: 1: Paper, 2: EMail, 4: Mailbox (electronic exchange)
	IsAccountant                        IsAccountant    `json:"IsAccountant"`                                // Indicates whether the account is an accountant. Values: 0 = No accountant, 1 = True, but accountant doesn't want his name to be published in the list of accountants, 2 = True, and accountant is published in the list of accountants
	IsAgency                            edm.Byte        `json:"IsAgency"`                                    // Indicates whether the accounti is an agency
	IsBank                              edm.Boolean     `json:"IsBank"`                                      // Obsolete
	IsCompetitor                        edm.Byte        `json:"IsCompetitor"`                                // Indicates whether the account is a competitor
	IsExtraDuty                         edm.Boolean     `json:"IsExtraDuty"`                                 // Indicates whether a customer is eligible for extra duty
	IsMailing                           edm.Byte        `json:"IsMailing"`                                   // Indicates if the account is excluded from mailing marketing information
	IsMember                            edm.Boolean     `json:"IsMember"`                                    // Obsolete
	IsPilot                             edm.Boolean     `json:"IsPilot"`                                     // Indicates whether the account is a pilot account
	IsPurchase                          edm.Boolean     `json:"IsPurchase"`                                  // Obsolete
	IsReseller                          edm.Boolean     `json:"IsReseller"`                                  // Indicates whether the account is a reseller
	IsSales                             edm.Boolean     `json:"IsSales"`                                     // Indicates whether the account is allowed for sales
	IsSupplier                          edm.Boolean     `json:"IsSupplier"`                                  // Indicates whether the account is a supplier
	Language                            edm.String      `json:"Language"`                                    // Language code
	LanguageDescription                 edm.String      `json:"LanguageDescription"`                         // Language description
	Latitude                            edm.Double      `json:"Latitude"`                                    // Latitude (used by Google maps)
	LeadSource                          edm.GUID        `json:"LeadSource"`                                  // Reference to Lead source of an account
	Logo                                edm.Binary      `json:"Logo"`                                        // Bytes of the logo image
	LogoFileName                        edm.String      `json:"LogoFileName"`                                // The file name (without path, but with extension) of the image
	LogoThumbnailUrl                    edm.String      `json:"LogoThumbnailUrl"`                            // Thumbnail url of the logo
	LogoUrl                             edm.String      `json:"LogoUrl"`                                     // Url to retrieve the logo
	Longitude                           edm.Double      `json:"Longitude"`                                   // Longitude (used by Google maps)
	MainContact                         edm.GUID        `json:"MainContact"`                                 // Reference to main contact person
	Modified                            edm.DateTime    `json:"Modified"`                                    // Last modified date
	Modifier                            edm.GUID        `json:"Modifier"`                                    // User ID of modifier
	ModifierFullName                    edm.String      `json:"ModifierFullName"`                            // Name of modifier
	Name                                edm.String      `json:"Name"`                                        // Account name
	Parent                              edm.GUID        `json:"Parent"`                                      // ID of the parent account
	PayAsYouEarn                        edm.String      `json:"PayAsYouEarn"`                                // Indicates the loan repayment plan for UK legislation
	PaymentConditionPurchase            edm.String      `json:"PaymentConditionPurchase"`                    // Code of default payment condition for purchase
	PaymentConditionPurchaseDescription edm.String      `json:"PaymentConditionPurchaseDescription"`         // Description of PaymentConditionPurchase
	PaymentConditionSales               edm.String      `json:"PaymentConditionSales"`                       // Code of default payment condition for sales
	PaymentConditionSalesDescription    edm.String      `json:"PaymentConditionSalesDescription"`            // Description of PaymentConditionSales
	Phone                               edm.String      `json:"Phone"`                                       // Phone number
	PhoneExtension                      edm.String      `json:"PhoneExtension"`                              // Phone number extention
	Postcode                            edm.String      `json:"Postcode"`                                    // Visit address postcode
	PriceList                           edm.GUID        `json:"PriceList"`                                   // Default sales price list for account
	PurchaseCurrency                    edm.String      `json:"PurchaseCurrency"`                            // Currency of purchase
	PurchaseCurrencyDescription         edm.String      `json:"PurchaseCurrencyDescription"`                 // Description of PurchaseCurrency
	PurchaseLeadDays                    edm.Int32       `json:"PurchaseLeadDays"`                            // Indicates number of days required to receive a purchase. Acts as a default
	PurchaseVATCode                     edm.String      `json:"PurchaseVATCode"`                             // Default VAT code used for purchase entries
	PurchaseVATCodeDescription          edm.String      `json:"PurchaseVATCodeDescription"`                  // Description of PurchaseVATCode
	RecepientOfCommissions              edm.Boolean     `json:"RecepientOfCommissions"`                      // Define the relation that should be taken in the official document of the rewarding fiscal fiches Belcotax
	Remarks                             edm.String      `json:"Remarks"`                                     // Remarks
	Reseller                            edm.GUID        `json:"Reseller"`                                    // ID of the reseller account. Conditions: the target account must have the property IsReseller turned on
	ResellerCode                        edm.String      `json:"ResellerCode"`                                // Code of Reseller
	ResellerName                        edm.String      `json:"ResellerName"`                                // Name of Reseller
	RSIN                                edm.String      `json:"RSIN"`                                        // Fiscal number for NL legislation
	SalesCurrency                       edm.String      `json:"SalesCurrency"`                               // Currency of Sales used for Time & Billing
	SalesCurrencyDescription            edm.String      `json:"SalesCurrencyDescription"`                    // Description of SalesCurrency
	SalesTaxSchedule                    edm.GUID        `json:"SalesTaxSchedule"`                            // Obsolete
	SalesTaxScheduleCode                edm.String      `json:"SalesTaxScheduleCode"`                        // Obsolete
	SalesTaxScheduleDescription         edm.String      `json:"SalesTaxScheduleDescription"`                 // Obsolete
	SalesVATCode                        edm.String      `json:"SalesVATCode"`                                // Default VAT code for a sales entry
	SalesVATCodeDescription             edm.String      `json:"SalesVATCodeDescription"`                     // Description of SalesVATCode
	SearchCode                          edm.String      `json:"SearchCode"`                                  // Search code
	SecurityLevel                       edm.Int32       `json:"SecurityLevel"`                               // Security level (0 - 100)
	SeparateInvPerProject               edm.Byte        `json:"SeparateInvPerProject"`                       // Separate invoice per project (Time & Billing)
	SeparateInvPerSubscription          edm.Byte        `json:"SeparateInvPerSubscription"`                  // Indicates how invoices are generated from subscriptions. 0 = subscriptions belonging to the same customer are combined in a single invoice. 1 = each subscription results in one invoice. In both cases, each individual subscription line results in one invoice line
	ShippingLeadDays                    edm.Int32       `json:"ShippingLeadDays"`                            // Indicates the number of days it takes to send goods to the customer. Acts as a default
	ShippingMethod                      edm.GUID        `json:"ShippingMethod"`                              // Default shipping method
	StartDate                           edm.DateTime    `json:"StartDate"`                                   // Indicates in combination with the end date if the account is active
	State                               edm.String      `json:"State"`                                       // State/Province code
	StateName                           edm.String      `json:"StateName"`                                   // Name of State
	Status                              Status          `json:"Status"`                                      // If the status field is filled this means the account is a customer. The value indicates the customer status. Possible values: A=None, S=Suspect, P=Prospect, C=Customer
	StatusSince                         edm.DateTime    `json:"StatusSince"`                                 // Obsolete
	TradeName                           edm.String      `json:"TradeName"`                                   // Trade name can be registered and shown with the client (for all legislations)
	Type                                Type            `json:"Type"`                                        // Account type: Values: A = Relation, D = Division
	UniqueTaxpayerReference             edm.String      `json:"UniqueTaxpayerReference"`                     // Unique taxpayer reference for UK legislation
	VATLiability                        edm.String      `json:"VATLiability"`                                // Indicates the VAT status of an account to be able to identify the relation that should be selected in the VAT debtor listing in Belgium
	VATNumber                           edm.String      `json:"VATNumber"`                                   // The number under which the account is known at the Value Added Tax collection agency
	Website                             edm.String      `json:"Website"`                                     // Website of the account
}

type NewAccount Account

type BankAccounts []BankAccount

// standalone: "BankAccounts": []
// deferred: "BankAccounts": {"__deferred": {}}
// embedded: "BankAccounts": {"results": []}
func (b *BankAccounts) UnmarshalJSON(data []byte) error {
	type Results BankAccounts

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

		*b = BankAccounts(*results)
		return nil
	}

	envelope := &Envelope{Results: Results(*b)}
	err = json.Unmarshal(data, envelope)
	if err != nil {
		return err
	}

	*b = BankAccounts(envelope.Results)
	return nil
}

type BankAccount struct {
	ID                    edm.GUID     `json:"ID"`                    // Primary key
	Account               edm.GUID     `json:"Account"`               // Account (customer, supplier) to which the bank account belongs
	AccountName           edm.String   `json:"AccountName"`           // The name of the account
	Bank                  edm.GUID     `json:"Bank"`                  // Obsolete
	BankAccount           edm.String   `json:"BankAccount"`           // The bank account number
	BankAccountHolderName edm.String   `json:"BankAccountHolderName"` // Name of the holder of the bank account, as known by the bank
	BankDescription       edm.String   `json:"BankDescription"`       // Obsolete
	BankName              edm.String   `json:"BankName"`              // Obsolete
	BICCode               edm.String   `json:"BICCode"`               // BIC code of the bank where the bank account is held
	Created               edm.DateTime `json:"Created"`               // Creation date
	Creator               edm.GUID     `json:"Creator"`               // User ID of creator
	CreatorFullName       edm.String   `json:"CreatorFullName"`       // Name of creator
	Description           edm.String   `json:"Description"`           // Description of the bank account
	Division              edm.Int32    `json:"Division"`              // Division code
	Format                edm.String   `json:"Format"`                // Format that belongs to the bank account number
	IBAN                  edm.String   `json:"IBAN"`                  // Obsolete
	Main                  edm.Boolean  `json:"Main"`                  // Indicates if the bank account is the main bank account
	Modified              edm.DateTime `json:"Modified"`              // Last modified date
	Modifier              edm.GUID     `json:"Modifier"`              // User ID of modifier
	ModifierFullName      edm.String   `json:"ModifierFullName"`      // Name of modifier
	PaymentServiceAccount edm.GUID     `json:"PaymentServiceAccount"` // ID of the Payment service account. Used when Type is 'P' (Payment service)
	Type                  edm.String   `json:"Type"`                  // The type indicates what entity the bank account is used for. A = Account (default), E = Employee, K = Cash, P = Payment service, R = Bank, S = Student, U = Unknown. Currently it's only possible to create 'Account' type bank accounts.
	TypeDescription       edm.String   `json:"TypeDescription"`       // Description of the Type
}

const (
	NoConsolidationScenario                     ConsolidationScenario = 0
	ItemConsolidationScenario                   ConsolidationScenario = 1
	ItemProjectConsolidationScenario            ConsolidationScenario = 2
	ItemEmployeeConsolidationScenario           ConsolidationScenario = 3
	ItemEmployeeProjectConsolidationScenario    ConsolidationScenario = 4
	ProjectWbsItemConsolidationScenario         ConsolidationScenario = 5
	ProjectWbsItemEmployeeConsolidationScenario ConsolidationScenario = 6
)

type ConsolidationScenario edm.Byte

func (c *ConsolidationScenario) String() string {
	switch int(*c) {
	case 0:
		return "No"
	case 1:
		return "Item"
	case 2:
		return "Item + Project"
	case 3:
		return "Item + Employee"
	case 4:
		return "Item + Employee + Project"
	case 5:
		return "Project + WBS + Item"
	case 6:
		return "Project + WBS + Item + Employee"
	}
	return ""
}

const (
	PaperInvoicingMethod   InvoicingMethod = 1
	EmailInvoicingMethod   InvoicingMethod = 2
	MailboxInvoicingMethod InvoicingMethod = 4
)

type InvoicingMethod edm.Int32

const (
	NoIsAccountant          IsAccountant = 0
	PublishedIsAccountantUn IsAccountant = 1
	PublishedIsAccountant   IsAccountant = 2
)

type IsAccountant edm.Byte

type SeparateInvPerSubscription edm.Byte

const (
	NoneStatus     Status = "A"
	SuspectStatus  Status = "S"
	ProspectStatus Status = "P"
	CustomerStatus Status = "C"
)

type Status edm.String

const (
	RelationType Type = "A"
	DivisionType Type = "D"
)

type Type edm.String
