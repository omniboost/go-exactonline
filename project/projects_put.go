package project

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/omitempty"
)

// POST

func (s *Service) ProjectsPut(body *ProjectsPutBody, ctx context.Context) error {
	method := http.MethodPut
	path := s.rest.SubPathWithID(ProjectsEndpoint, body.ID.String())

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, nil)
	return err
}

type ProjectsPutBody struct {
	ID                               edm.GUID     `json:"ID,omitempty"`                               // Primary key
	Account                          edm.GUID     `json:"Account,omitempty"`                          // The account for this project
	AccountContact                   edm.GUID     `json:"AccountContact,omitempty"`                   // Contact person of Account
	AllowAdditionalInvoicing         edm.Boolean  `json:"AllowAdditionalInvoicing,omitempty"`         // Indicates if additional invoice is allowed for project
	BlockEntry                       edm.Boolean  `json:"BlockEntry,omitempty"`                       // Block time and cost entries
	BlockPurchasing                  edm.Boolean  `json:"BlockPurchasing,omitempty"`                  // Block purchasing
	BlockRebilling                   edm.Boolean  `json:"BlockRebilling,omitempty"`                   // Block rebilling
	BudgetedAmount                   edm.Double   `json:"BudgetedAmount,omitempty"`                   // Budgeted amount of sales in the default currency of the company
	BudgetedCosts                    edm.Double   `json:"BudgetedCosts,omitempty"`                    // Budgeted amount of costs in the default currency of the company
	BudgetedRevenue                  edm.Double   `json:"BudgetedRevenue,omitempty"`                  // Budgeted amount of revenue in the default currency of the company
	BudgetOverrunHours               edm.Byte     `json:"BudgetOverrunHours,omitempty"`               // BudgetOverrunHours: 10-Allowed, 20-Not Allowed
	BudgetType                       edm.Int16    `json:"BudgetType,omitempty"`                       // Budget type
	Classification                   edm.GUID     `json:"Classification,omitempty"`                   // Used only for PSA to link a project classification to the project
	Code                             edm.String   `json:"Code"`                                       // Code
	CustomerPOnumber                 edm.String   `json:"CustomerPOnumber,omitempty"`                 // Used only for PSA to store the customer's PO number
	Description                      edm.String   `json:"Description"`                                // Description of the project
	EndDate                          edm.DateTime `json:"EndDate,omitempty"`                          // End date of the project. In combination with the start date the status is determined
	FixedPriceItem                   edm.GUID     `json:"FixedPriceItem,omitempty"`                   // Item used for fixed price invoicing. To be defined per project. If empty the functionality relies on the setting
	IncludeInvoiceSpecification      edm.Int16    `json:"IncludeInvoiceSpecification,omitempty"`      // Include invoice specification. E.g: 1 = Based on account, 2 = Always, 3 = Never
	IncludeSpecificationInInvoicePdf edm.Boolean  `json:"IncludeSpecificationInInvoicePdf,omitempty"` // Indicates whether to include invoice specification in invoice PDF
	InternalNotes                    edm.String   `json:"InternalNotes,omitempty"`                    // Internal notes not to be printed in invoice
	InvoiceAddress                   edm.GUID     `json:"InvoiceAddress,omitempty"`                   // Invoice address
	Manager                          edm.GUID     `json:"Manager,omitempty"`                          // Responsible person for this project
	MarkupPercentage                 edm.Double   `json:"MarkupPercentage,omitempty"`                 // Purchase markup percentage
	Notes                            edm.String   `json:"Notes,omitempty"`                            // For additional information about projects
	PrepaidItem                      edm.GUID     `json:"PrepaidItem,omitempty"`                      // Used only for PSA. This item is used for prepaid invoicing. If left empty, the functionality relies on a setting
	PrepaidType                      edm.Int16    `json:"PrepaidType,omitempty"`                      // PrepaidType: 1-Retainer, 2-Hour type bundle
	SalesTimeQuantity                edm.Double   `json:"SalesTimeQuantity,omitempty"`                // Budgeted time. Total number of hours estimated for the fixed price project
	StartDate                        edm.DateTime `json:"StartDate,omitempty"`                        // Start date of a project. In combination with the end date the status is determined
	TimeQuantityToAlert              edm.Double   `json:"TimeQuantityToAlert,omitempty"`              // Alert when exceeding (Hours)
	Type                             edm.Int32    `json:"Type"`                                       // Reference to ProjectTypes. E.g: 1 = Campaign , 2 = Fixed Price, 3 = Time and Material, 4 = Non billable, 5 = Prepaid
	UseBillingMilestones             edm.Boolean  `json:"UseBillingMilestones,omitempty"`             // Indicates whether the Project is using billing milestones
}

func (s *Service) NewProjectsPutBody() *ProjectsPutBody {
	return &ProjectsPutBody{}
}

func (b ProjectsPutBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(b)
}
