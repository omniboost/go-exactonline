package financial

import "github.com/omniboost/go-exactonline/edm"

type Journals []Journal

type Journal struct {
	ID                              edm.GUID     `json:"ID"`
	AllowVariableCurrency           edm.Boolean  `json:"AllowVariableCurrency"`
	AllowVariableExchangeRate       edm.Boolean  `json:"AllowVariableExchangeRate"`
	AllowVAT                        edm.Boolean  `json:"AllowVAT"`
	AutoSave                        edm.Boolean  `json:"AutoSave"`
	Bank                            edm.GUID     `json:"Bank"`
	BankAccountBICCode              edm.String   `json:"BankAccountBICCode"`
	BankAccountCountry              edm.String   `json:"BankAccountCountry"`
	BankAccountDescription          edm.String   `json:"BankAccountDescription"`
	BankAccountIBAN                 edm.String   `json:"BankAccountIBAN"`
	BankAccountID                   edm.GUID     `json:"BankAccountID"`
	BankAccountIncludingMask        edm.String   `json:"BankAccountIncludingMask"`
	BankAccountUseSEPA              edm.Boolean  `json:"BankAccountUseSEPA"`
	BankAccountUseSepaDirectDebit   edm.Boolean  `json:"BankAccountUseSepaDirectDebit"`
	BankName                        edm.String   `json:"BankName"`
	Code                            edm.String   `json:"Code"`
	Created                         edm.DateTime `json:"Created"`
	Creator                         edm.GUID     `json:"Creator"`
	CreatorFullName                 edm.String   `json:"CreatorFullName"`
	Currency                        edm.String   `json:"Currency"`
	CurrencyDescription             edm.String   `json:"CurrencyDescription"`
	Description                     edm.String   `json:"Description"`
	Division                        edm.Int32    `json:"Division"`
	GLAccount                       edm.GUID     `json:"GLAccount"`
	GLAccountCode                   edm.String   `json:"GLAccountCode"`
	GLAccountDescription            edm.String   `json:"GLAccountDescription"`
	GLAccountType                   edm.Int32    `json:"GLAccountType"`
	Modified                        edm.DateTime `json:"Modified"`
	Modifier                        edm.GUID     `json:"Modifier"`
	ModifierFullName                edm.String   `json:"ModifierFullName"`
	PaymentInTransitAccount         edm.GUID     `json:"PaymentInTransitAccount"`
	PaymentServiceAccountIdentifier edm.String   `json:"PaymentServiceAccountIdentifier"`
	PaymentServiceProvider          edm.Int32    `json:"PaymentServiceProvider"`
	PaymentServiceProviderName      edm.String   `json:"PaymentServiceProviderName"`
	Type                            JournalType  `json:"Type"`
}

const (
	JournalTypeCash                  JournalType = 10
	JournalTypeBank                  JournalType = 12
	JournalTypePaymentService        JournalType = 16
	JournalTypeSales                 JournalType = 20
	JournalTypeReturnInvoce          JournalType = 21
	JournalTypePurchase              JournalType = 22
	JournalTypeReceivedReturnInvoice JournalType = 23
	JournalTypeGeneralJournal        JournalType = 90
)

type JournalType edm.Int32

func (j *JournalType) String() string {
	switch int(*j) {
	case 10:
		return "Cash"
	case 12:
		return "Bank"
	case 16:
		return "Payment service"
	case 20:
		return "Sales"
	case 21:
		return "Return invoice"
	case 22:
		return "Purchase"
	case 23:
		return "Received return invoice"
	case 90:
		return "General journal"
	}
	return ""
}

type GLSchemes []GLScheme

type GLScheme struct {
	ID               edm.GUID     `json:"ID"`
	Code             edm.String   `json:"Code"`
	Created          edm.DateTime `json:"Created"`
	Creator          edm.GUID     `json:"Creator"`
	CreatorFullName  edm.String   `json:"CreatorFullName"`
	Description      edm.String   `json:"Description"`
	Division         edm.Int32    `json:"Division"`
	Main             edm.Byte     `json:"Main"`
	Modified         edm.DateTime `json:"Modified"`
	Modifier         edm.GUID     `json:"Modifier"`
	ModifierFullName edm.String   `json:"ModifierFullName"`
	TargetNamespace  edm.String   `json:"TargetNamespace"`
}

type GLAccounts []GLAccount

type GLAccount struct {
	ID                             edm.GUID      `json:"ID"`
	AssimilatedVATBox              edm.Int16     `json:"AssimilatedVATBox"`
	BalanceSide                    edm.String    `json:"BalanceSide"`
	BalanceType                    edm.String    `json:"BalanceType"`
	BelcotaxType                   edm.Int32     `json:"BelcotaxType"`
	Code                           edm.String    `json:"Code"`
	Compress                       edm.Boolean   `json:"Compress"`
	Costcenter                     edm.String    `json:"Costcenter"`
	CostcenterDescription          edm.String    `json:"CostcenterDescription"`
	Costunit                       edm.String    `json:"Costunit"`
	CostunitDescription            edm.String    `json:"CostunitDescription"`
	Created                        edm.DateTime  `json:"Created"`
	Creator                        edm.GUID      `json:"Creator"`
	CreatorFullName                edm.String    `json:"CreatorFullName"`
	Description                    edm.String    `json:"Description"`
	Division                       edm.Int32     `json:"Division"`
	ExcludeVATListing              edm.Byte      `json:"ExcludeVATListing"`
	ExpenseNonDeductiblePercentage edm.Double    `json:"ExpenseNonDeductiblePercentage"`
	IsBlocked                      edm.Boolean   `json:"IsBlocked"`
	Matching                       edm.Boolean   `json:"Matching"`
	Modified                       edm.DateTime  `json:"Modified"`
	Modifier                       edm.GUID      `json:"Modifier"`
	ModifierFullName               edm.String    `json:"ModifierFullName"`
	PrivateGLAccount               edm.GUID      `json:"PrivateGLAccount"`
	PrivatePercentage              edm.Double    `json:"PrivatePercentage"`
	ReportingCode                  edm.String    `json:"ReportingCode"`
	RevalueCurrency                edm.Boolean   `json:"RevalueCurrency"`
	SearchCode                     edm.String    `json:"SearchCode"`
	Type                           GLAccountType `json:"Type"`
	TypeDescription                edm.String    `json:"TypeDescription"`
	UseCostcenter                  edm.Byte      `json:"UseCostcenter"`
	UseCostunit                    edm.Byte      `json:"UseCostunit"`
	VATCode                        edm.String    `json:"VATCode"`
	VATDescription                 edm.String    `json:"VATDescription"`
	VATGLAccountType               edm.String    `json:"VATGLAccountType"`
	VATNonDeductibleGLAccount      edm.GUID      `json:"VATNonDeductibleGLAccount"`
	VATNonDeductiblePercentage     edm.Double    `json:"VATNonDeductiblePercentage"`
	VATSystem                      edm.String    `json:"VATSystem"`
	YearEndCostGLAccount           edm.GUID      `json:"YearEndCostGLAccount"`
	YearEndReflectionGLAccount     edm.GUID      `json:"YearEndReflectionGLAccount"`
}

const (
	GLAccountTypeCash                               GLAccountType = 10  // Cash
	GLAccountTypeBank                               GLAccountType = 12  // Bank
	GLAccountTypeCreditCard                         GLAccountType = 14  // Credit card
	GLAccountTypePaymentServices                    GLAccountType = 16  // Payment services
	GLAccountTypeAccountReceivable                  GLAccountType = 20  // Accounts receivable
	GLAccountTypePrepaymentAccountsReceivable       GLAccountType = 21  // Prepayment accounts receivable
	GLAccountTypeAccountsPayable                    GLAccountType = 22  // Accounts payable
	GLAccountTypeVAT                                GLAccountType = 24  // VAT
	GLAccountTypeEmployeesPayable                   GLAccountType = 25  // Employees payable
	GLAccountTypePrepaidExpenses                    GLAccountType = 26  // Prepaid expenses
	GLAccountTypeAccruedExpenses                    GLAccountType = 27  // Accrued expenses
	GLAccountTypeIncomeTaxesPayable                 GLAccountType = 29  // Income taxes payable
	GLAccountTypeFixedAssets                        GLAccountType = 30  // Fixed assets
	GLAccountTypedOtherAssets                       GLAccountType = 32  // Other assets
	GLAccountTypeAccumulatedDepreciation            GLAccountType = 35  // Accumulated depreciation
	GLAccountTypeInventory                          GLAccountType = 40  // Inventory
	GLAccountTypeCapitalStock                       GLAccountType = 50  // Capital stock
	GLAccountTypeRetainedEarnings                   GLAccountType = 52  // Retained earnings
	GLAccountTypeLongTermDebt                       GLAccountType = 55  // Long term debt
	GLAccountTypeCurrentPortionOfDebt               GLAccountType = 60  // Current portion of debt
	GLAccountTypeGeneral                            GLAccountType = 90  // General
	GLAccountTypeTaxPayable                         GLAccountType = 100 // Tax payable
	GLAccountTypeRevenue                            GLAccountType = 110 // Revenue
	GLAccountTypeCostOfGoods                        GLAccountType = 111 // Cost of goods
	GLAccountTypeOtherCosts                         GLAccountType = 120 // Other costs
	GLAccountTypeSalesGeneralAdministrativeExpenses GLAccountType = 121 // Sales, general administrative expenses
	GLAccountTypeDepreciationCosts                  GLAccountType = 122 // Depreciation costs
	GLAccountTypeResearchAndDevelopment             GLAccountType = 123 // Research and development
	GLAccountTypeEmployeeCosts                      GLAccountType = 125 // Employee costs
	GLAccountTypeEmploymentCosts                    GLAccountType = 126 // Employment costs
	GLAccountTypeExceptionalCosts                   GLAccountType = 130 // Exceptional costs
	GLAccountTypeExceptionalIncome                  GLAccountType = 140 // Exceptional income
	GLAccountTypeIncomeTaxes                        GLAccountType = 150 // Income taxes
	GLAccountTypeInterestIncome                     GLAccountType = 160 // Interest income
	GLAccountTypeYearEndReflection                  GLAccountType = 300 // Year end reflection
	GLAccountTypeIndirectYearEndCosting             GLAccountType = 301 // Indirect year end costing
	GLAccountTypeDirectYearEndCosting               GLAccountType = 302 // Direct year end costing
)

type GLAccountType edm.Int32

func (g *GLAccountType) String() string {
	switch int(*g) {
	case 10:
		return "Cash"
	case 12:
		return "Bank"
	case 14:
		return "Credit card"
	case 16:
		return "Payment services"
	case 20:
		return "Accounts receivable"
	case 21:
		return "Prepayment accounts receivable"
	case 22:
		return "Accounts payable"
	case 24:
		return "VAT"
	case 25:
		return "Employees payable"
	case 26:
		return "Prepaid expenses"
	case 27:
		return "Accrued expenses"
	case 29:
		return "Income taxes payable"
	case 30:
		return "Fixed assets"
	case 32:
		return "Other assets"
	case 35:
		return "Accumulated depreciation"
	case 40:
		return "Inventory"
	case 50:
		return "Capital stock"
	case 52:
		return "Retained earnings"
	case 55:
		return "Long term debt"
	case 60:
		return "Current portion of debt"
	case 90:
		return "General"
	case 100:
		return "Tax payable"
	case 110:
		return "Revenue"
	case 111:
		return "Cost of goods"
	case 120:
		return "Other costs"
	case 121:
		return "Sales, general administrative expenses"
	case 122:
		return "Depreciation costs"
	case 123:
		return "Research and development"
	case 125:
		return "Employee costs"
	case 126:
		return "Employment costs"
	case 130:
		return "Exceptional costs"
	case 140:
		return "Exceptional income"
	case 150:
		return "Income taxes"
	case 160:
		return "Interest income"
	case 300:
		return "Year end reflection"
	case 301:
		return "Indirect year end costing"
	case 302:
		return "Direct year end costing"
	}
	return ""
}

type Periods []Period

type Period struct {
	ID               edm.GUID     `json:"ID"`               // Primary key
	Created          edm.DateTime `json:"Created"`          // Creation date
	Creator          edm.GUID     `json:"Creator"`          // User ID of creator
	CreatorFullName  edm.String   `json:"CreatorFullName"`  // Name of creator
	Division         edm.Int32    `json:"Division"`         // Division code
	EndDate          edm.DateTime `json:"EndDate"`          // The end date of the period
	FinPeriod        edm.Int16    `json:"FinPeriod"`        // The financial period. Usually the period is a month or quarter with period 1 starting on the first of January.
	FinYear          edm.Int16    `json:"FinYear"`          // The financial year. The financial year and calendar year are not always aligned.
	Modified         edm.DateTime `json:"Modified"`         // Last modified date
	Modifier         edm.GUID     `json:"Modifier"`         // User ID of modifier
	ModifierFullName edm.String   `json:"ModifierFullName"` // Name of modifier
	StartDate        edm.DateTime `json:"StartDate"`        // The start date of a period. A start date should always succeed a previous end date. Except for the first year/period combination
}
