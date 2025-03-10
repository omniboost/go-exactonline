package project

import "github.com/omniboost/go-exactonline/edm"

type Projects []Project

type Project struct {
	ID                               edm.GUID     `json:"ID,omitempty"`                               // Primary key
	Account                          edm.GUID     `json:"Account,omitempty"`                          // The account for this project
	AccountCode                      edm.String   `json:"AccountCode,omitempty"`                      // Code of Account
	AccountContact                   edm.GUID     `json:"AccountContact,omitempty"`                   // Contact person of Account
	AccountName                      edm.String   `json:"AccountName,omitempty"`                      // Name of Account
	AllowAdditionalInvoicing         edm.Boolean  `json:"AllowAdditionalInvoicing,omitempty"`         // Indicates if additional invoice is allowed for project
	BlockEntry                       edm.Boolean  `json:"BlockEntry,omitempty"`                       // Block time and cost entries
	BlockPurchasing                  edm.Boolean  `json:"BlockPurchasing,omitempty"`                  // Block purchasing
	BlockRebilling                   edm.Boolean  `json:"BlockRebilling,omitempty"`                   // Block rebilling
	BudgetedAmount                   edm.Double   `json:"BudgetedAmount,omitempty"`                   // Budgeted amount of sales in the default currency of the company
	BudgetedCosts                    edm.Double   `json:"BudgetedCosts,omitempty"`                    // Budgeted amount of costs in the default currency of the company
	BudgetedHoursPerHourType         interface{}  `json:"BudgetedHoursPerHourType,omitempty"`         // Collection of budgeted hours
	BudgetedRevenue                  edm.Double   `json:"BudgetedRevenue,omitempty"`                  // Budgeted amount of revenue in the default currency of the company
	BudgetOverrunHours               edm.Byte     `json:"BudgetOverrunHours,omitempty"`               // BudgetOverrunHours: 10-Allowed, 20-Not Allowed
	BudgetType                       edm.Int16    `json:"BudgetType,omitempty"`                       // Budget type
	BudgetTypeDescription            edm.String   `json:"BudgetTypeDescription,omitempty"`            // Budget type description
	Classification                   edm.GUID     `json:"Classification,omitempty"`                   // Used only for PSA to link a project classification to the project
	ClassificationDescription        edm.String   `json:"ClassificationDescription,omitempty"`        // Description of Classification
	Code                             edm.String   `json:"Code"`                                       // Code
	CostsAmountFC                    edm.Double   `json:"CostsAmountFC,omitempty"`                    // Used only for PSA to store the budgetted costs of a project (except for project type Campaign and Non-billable). Positive quantities only
	Created                          edm.DateTime `json:"Created,omitempty"`                          // Creation date
	Creator                          edm.GUID     `json:"Creator,omitempty"`                          // User ID of creator
	CreatorFullName                  edm.String   `json:"CreatorFullName,omitempty"`                  // Name of creator
	CustomerPOnumber                 edm.String   `json:"CustomerPOnumber,omitempty"`                 // Used only for PSA to store the customer's PO number
	CustomField                      edm.String   `json:"CustomField,omitempty"`                      // Custom field endpoint
	Description                      edm.String   `json:"Description"`                                // Description of the project
	Division                         edm.Int32    `json:"Division,omitempty"`                         // Division code
	DivisionName                     edm.String   `json:"DivisionName,omitempty"`                     // Name of Division
	EndDate                          edm.DateTime `json:"EndDate,omitempty"`                          // End date of the project. In combination with the start date the status is determined
	FixedPriceItem                   edm.GUID     `json:"FixedPriceItem,omitempty"`                   // Item used for fixed price invoicing. To be defined per project. If empty the functionality relies on the setting
	FixedPriceItemDescription        edm.String   `json:"FixedPriceItemDescription,omitempty"`        // Description of FixedPriceItem
	HasWBSLines                      edm.Boolean  `json:"HasWBSLines,omitempty"`                      // Indicates if whether the Project has WBS
	IncludeInvoiceSpecification      edm.Int16    `json:"IncludeInvoiceSpecification,omitempty"`      // Include invoice specification. E.g: 1 = Based on account, 2 = Always, 3 = Never
	IncludeSpecificationInInvoicePdf edm.Boolean  `json:"IncludeSpecificationInInvoicePdf,omitempty"` // Indicates whether to include invoice specification in invoice PDF
	InternalNotes                    edm.String   `json:"InternalNotes,omitempty"`                    // Internal notes not to be printed in invoice
	InvoiceAddress                   edm.GUID     `json:"InvoiceAddress,omitempty"`                   // Invoice address
	InvoiceAsQuoted                  edm.Boolean  `json:"InvoiceAsQuoted,omitempty"`                  // Indicates whether the project is invoice as quoted
	InvoiceTerms                     interface{}  `json:"InvoiceTerms,omitempty"`                     // Collection of invoice terms
	Manager                          edm.GUID     `json:"Manager,omitempty"`                          // Responsible person for this project
	ManagerFullname                  edm.String   `json:"ManagerFullname,omitempty"`                  // Name of Manager
	MarkupPercentage                 edm.Double   `json:"MarkupPercentage,omitempty"`                 // Purchase markup percentage
	Modified                         edm.DateTime `json:"Modified,omitempty"`                         // Last modified date
	Modifier                         edm.GUID     `json:"Modifier,omitempty"`                         // User ID of modifier
	ModifierFullName                 edm.String   `json:"ModifierFullName,omitempty"`                 // Name of modifier
	Notes                            edm.String   `json:"Notes,omitempty"`                            // For additional information about projects
	PrepaidItem                      edm.GUID     `json:"PrepaidItem,omitempty"`                      // Used only for PSA. This item is used for prepaid invoicing. If left empty, the functionality relies on a setting
	PrepaidItemDescription           edm.String   `json:"PrepaidItemDescription,omitempty"`           // Description of PrepaidItem
	PrepaidType                      edm.Int16    `json:"PrepaidType,omitempty"`                      // PrepaidType: 1-Retainer, 2-Hour type bundle
	PrepaidTypeDescription           edm.String   `json:"PrepaidTypeDescription,omitempty"`           // Description of PrepaidType
	ProjectRestrictionEmployees      interface{}  `json:"ProjectRestrictionEmployees,omitempty"`      // Collection of employee restrictions
	ProjectRestrictionItems          interface{}  `json:"ProjectRestrictionItems,omitempty"`          // Collection of item restrictions
	ProjectRestrictionRebillings     interface{}  `json:"ProjectRestrictionRebillings,omitempty"`     // Collection of rebilling restrictions
	SalesTimeQuantity                edm.Double   `json:"SalesTimeQuantity,omitempty"`                // Budgeted time. Total number of hours estimated for the fixed price project
	SourceQuotation                  edm.GUID     `json:"SourceQuotation,omitempty"`                  // Source quotation
	StartDate                        edm.DateTime `json:"StartDate,omitempty"`                        // Start date of a project. In combination with the end date the status is determined
	TimeQuantityToAlert              edm.Double   `json:"TimeQuantityToAlert,omitempty"`              // Alert when exceeding (Hours)
	Type                             edm.Int32    `json:"Type"`                                       // Reference to ProjectTypes. E.g: 1 = Campaign , 2 = Fixed Price, 3 = Time and Material, 4 = Non billable, 5 = Prepaid
	TypeDescription                  edm.String   `json:"TypeDescription,omitempty"`                  // Description of Type
	UseBillingMilestones             edm.Boolean  `json:"UseBillingMilestones,omitempty"`             // Indicates whether the Project is using billing milestones
}

type NewProject Project

type HourTypes []HourType

type HourType struct {
	ItemId          edm.GUID   `json:"ItemId,omitempty"`          // GUID id of the item that is linked to the project
	ItemDescription edm.String `json:"ItemDescription,omitempty"` // Description of the item that is linked to the project
}

type TimeTransactions []TimeTransaction

type TimeTransaction struct {
	ID                      edm.GUID     `json:"ID,omitempty"`                      // Primary key
	Account                 edm.GUID     `json:"Account,omitempty"`                 // Guid ID of account that is linked to the project
	AccountName             edm.String   `json:"AccountName,omitempty"`             // Name of account that is linked to the project
	Activity                edm.GUID     `json:"Activity,omitempty"`                // Guid ID of activity that is linked to project WBS (work breakdown structure)
	ActivityDescription     edm.String   `json:"ActivityDescription,omitempty"`     // Name of activity that is linked to project WBS (work breakdown structure)
	AmountFC                edm.Double   `json:"AmountFC,omitempty"`                // Calculated amount of the transaction based on (Quantity * PriceFC)
	Attachment              edm.GUID     `json:"Attachment,omitempty"`              // Attachment linked to the transaction (not mandatory)
	Created                 edm.DateTime `json:"Created,omitempty"`                 // Date and time the transaction was created
	Creator                 edm.GUID     `json:"Creator,omitempty"`                 // The Guid ID of user that created the transaction
	CreatorFullName         edm.String   `json:"CreatorFullName,omitempty"`         // The full name of the user that created the record
	Currency                edm.String   `json:"Currency,omitempty"`                // Currency of amount FC
	Date                    edm.DateTime `json:"Date,omitempty"`                    // Date and time the time transaction was done
	Division                edm.Int32    `json:"Division,omitempty"`                // Division code
	DivisionDescription     edm.String   `json:"DivisionDescription,omitempty"`     // Description of Division
	Employee                edm.GUID     `json:"Employee,omitempty"`                // Guid ID of the employee that is linked to the time transaction
	EndTime                 edm.DateTime `json:"EndTime,omitempty"`                 // End time of the time transaction
	EntryNumber             edm.Int32    `json:"EntryNumber,omitempty"`             // Number that represents the grouping of time transactions
	ErrorText               edm.String   `json:"ErrorText,omitempty"`               // (Only used by backgroundjobs) To determine which transaction has an error
	HourStatus              edm.Int16    `json:"HourStatus,omitempty"`              // Status of the transaction: 1 = Draft, 2 = Rejected, 10 = Submitted, 11 = Failed on approval, 14 = Processing, 16 = Processing, 19 = Failed while undoing approval, 20 = Final
	Item                    edm.GUID     `json:"Item"`                              // Item that is linked to the transaction, which provides the time information
	ItemDescription         edm.String   `json:"ItemDescription,omitempty"`         // Description of the item that is linked to the transaction
	ItemDivisable           edm.Boolean  `json:"ItemDivisable,omitempty"`           // Indicates if fractional quantities of the item can be used, for example quantity = 0.4
	Modified                edm.DateTime `json:"Modified,omitempty"`                // The date and time transaction record was modified
	Modifier                edm.GUID     `json:"Modifier,omitempty"`                // The Guid ID of the user that modified the records
	ModifierFullName        edm.String   `json:"ModifierFullName,omitempty"`        // The full name of the user that modified the record
	Notes                   edm.String   `json:"Notes,omitempty"`                   // Notes linked to the transaction for providing additional information (not mandatory)
	PriceFC                 edm.Double   `json:"PriceFC,omitempty"`                 // For use in AmountFC (Quantiy * Price FC)
	Project                 edm.GUID     `json:"Project"`                           // Guid ID of project that is linked to the transaction
	ProjectAccount          edm.GUID     `json:"ProjectAccount,omitempty"`          // Project account ID that is linked to the transaction (not mandatory)
	ProjectAccountCode      edm.String   `json:"ProjectAccountCode,omitempty"`      // Project account code that is linked to the transaction
	ProjectAccountName      edm.String   `json:"ProjectAccountName,omitempty"`      // Project account name that is linked to the transaction
	ProjectCode             edm.String   `json:"ProjectCode,omitempty"`             // Project code that is linked to the transaction
	ProjectDescription      edm.String   `json:"ProjectDescription,omitempty"`      // Project description that is linked to the transaction
	Quantity                edm.Double   `json:"Quantity"`                          // Quantity of the item that is linked to the transaction
	StartTime               edm.DateTime `json:"StartTime,omitempty"`               // Start time of the time transaction
	Subscription            edm.GUID     `json:"Subscription,omitempty"`            // Guid ID of subscription that is linked to the transaction
	SubscriptionAccount     edm.GUID     `json:"SubscriptionAccount,omitempty"`     // Subscription account ID that is linked to the transaction, this is to identify the referenced subscription
	SubscriptionAccountCode edm.String   `json:"SubscriptionAccountCode,omitempty"` // Subscription account code that is linked to the transaction
	SubscriptionAccountName edm.String   `json:"SubscriptionAccountName,omitempty"` // Subscription account name that is linked to the transaction
	SubscriptionDescription edm.String   `json:"SubscriptionDescription,omitempty"` // Subscription description that is linked to the transaction
	SubscriptionNumber      edm.Int32    `json:"SubscriptionNumber,omitempty"`      // Subscription number that is linked to the transaction
	Type                    edm.Int16    `json:"Type,omitempty"`                    // The type of transaction. E.g: 1 = Time, 2 = Cost
}

type NewTimeTransaction struct {
	ID                  edm.GUID     `json:"ID,omitempty"`                  // Primary key
	Account             edm.GUID     `json:"Account,omitempty"`             // Guid ID of account that is linked to the project
	Activity            edm.GUID     `json:"Activity,omitempty"`            // Guid ID of activity that is linked to project WBS (work breakdown structure)
	AmountFC            edm.Double   `json:"AmountFC,omitempty"`            // Calculated amount of the transaction based on (Quantity * PriceFC)
	Attachment          edm.GUID     `json:"Attachment,omitempty"`          // Attachment linked to the transaction (not mandatory)
	Currency            edm.String   `json:"Currency,omitempty"`            // Currency of amount FC
	Date                edm.DateTime `json:"Date,omitempty"`                // Date and time the time transaction was done
	Employee            edm.GUID     `json:"Employee,omitempty"`            // Guid ID of the employee that is linked to the time transaction
	EndTime             edm.DateTime `json:"EndTime,omitempty"`             // End time of the time transaction
	EntryNumber         edm.Int32    `json:"EntryNumber,omitempty"`         // Number that represents the grouping of time transactions
	ErrorText           edm.String   `json:"ErrorText,omitempty"`           // (Only used by backgroundjobs) To determine which transaction has an error
	HourStatus          edm.Int16    `json:"HourStatus,omitempty"`          // Status of the transaction: 1 = Draft, 2 = Rejected, 10 = Submitted, 11 = Failed on approval, 14 = Processing, 16 = Processing, 19 = Failed while undoing approval, 20 = Final
	Item                edm.GUID     `json:"Item"`                          // Item that is linked to the transaction, which provides the time information
	Notes               edm.String   `json:"Notes,omitempty"`               // Notes linked to the transaction for providing additional information (not mandatory)
	Project             edm.GUID     `json:"Project"`                       // Guid ID of project that is linked to the transaction
	Quantity            edm.Double   `json:"Quantity"`                      // Quantity of the item that is linked to the transaction
	StartTime           edm.DateTime `json:"StartTime,omitempty"`           // Start time of the time transaction
	Subscription        edm.GUID     `json:"Subscription,omitempty"`        // Guid ID of subscription that is linked to the transaction
	SubscriptionAccount edm.GUID     `json:"SubscriptionAccount,omitempty"` // Subscription account ID that is linked to the transaction, this is to identify the referenced subscription
}

type TimeAndBillingItemDetails []TimeAndBillingItemDetail

type TimeAndBillingItemDetail struct {
}

type ProjectRestrictionItems []ProjectRestrictionItem

type ProjectRestrictionItem struct {
}
