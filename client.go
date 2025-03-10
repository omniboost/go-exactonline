package exact

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactonline/crm"
	"github.com/omniboost/go-exactonline/document"
	"github.com/omniboost/go-exactonline/financial"
	"github.com/omniboost/go-exactonline/financialtransaction"
	"github.com/omniboost/go-exactonline/general"
	"github.com/omniboost/go-exactonline/generaljournalentry"
	"github.com/omniboost/go-exactonline/hrm"
	"github.com/omniboost/go-exactonline/logistics"
	"github.com/omniboost/go-exactonline/payroll"
	"github.com/omniboost/go-exactonline/project"
	"github.com/omniboost/go-exactonline/purchaseentry"
	"github.com/omniboost/go-exactonline/purchaseorder"
	"github.com/omniboost/go-exactonline/rest"
	"github.com/omniboost/go-exactonline/sales"
	"github.com/omniboost/go-exactonline/salesentry"
	"github.com/omniboost/go-exactonline/salesinvoice"
	"github.com/omniboost/go-exactonline/salesorder"
	"github.com/omniboost/go-exactonline/system"
	"github.com/omniboost/go-exactonline/vat"
)

const (
	DefaultBaseURL = "https://start.exactonline.nl/api"

	libraryVersion = "0.0.1"
	userAgent      = "go-exactonline/" + libraryVersion
)

// Client manages communication with Exact Online API
type Client struct {
	// REST client used to communicate with the API.
	rest.Client

	// Services
	// Accountancy          *Accountancy
	// Activities           *Activities
	// Assets               *Assets
	// Budget               *Budget
	// Bulk                 *Bulk
	// Cashflow             *Cashflow
	// ContinuousMonitoring *ContinuousMonitoring
	CRM                  *crm.Service
	Document             *document.Service
	Financial            *financial.Service
	FinancialTransaction *financialtransaction.Service
	General              *general.Service
	GeneralJournalEntry  *generaljournalentry.Service
	HRM                  *hrm.Service
	// Inventory            *Inventory
	Logistics *logistics.Service
	// Mailbox              *Mailbox
	// Manufacturing        *Manufacturing
	// OpeningBalance       *OpeningBalance
	Payroll       *payroll.Service
	Project       *project.Service
	PurchaseEntry *purchaseentry.Service
	PurchaseOrder *purchaseorder.Service
	Sales         *sales.Service
	SalesEntry    *salesentry.Service
	SalesInvoice  *salesinvoice.Service
	SalesOrder    *salesorder.Service
	// Subscription         *Subscription
	System *system.Service
	// Users                *Users
	VAT *vat.Service
	// Workflow             *Workflow
}

// NewClient returns a new Exact Online API client
func NewClient(httpClient *http.Client, divisionID int) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		Client: *rest.New(httpClient),
	}

	// set default options
	baseURL, _ := url.Parse(DefaultBaseURL)
	c.SetBaseURL(baseURL)
	c.SetDivisionID(divisionID)
	c.SetUserAgent(userAgent)
	c.SetDebug(false)

	c.CRM = crm.NewService(&c.Client)
	c.Document = document.NewService(&c.Client)
	c.Financial = financial.NewService(&c.Client)
	c.FinancialTransaction = financialtransaction.NewService(&c.Client)
	c.General = general.NewService(&c.Client)
	c.GeneralJournalEntry = generaljournalentry.NewService(&c.Client)
	c.HRM = hrm.NewService(&c.Client)
	c.Logistics = logistics.NewService(&c.Client)
	c.Payroll = payroll.NewService(&c.Client)
	c.Project = project.NewService(&c.Client)
	c.PurchaseEntry = purchaseentry.NewService(&c.Client)
	c.PurchaseOrder = purchaseorder.NewService(&c.Client)
	c.Sales = sales.NewService(&c.Client)
	c.SalesEntry = salesentry.NewService(&c.Client)
	c.SalesInvoice = salesinvoice.NewService(&c.Client)
	c.SalesOrder = salesorder.NewService(&c.Client)
	c.System = system.NewService(&c.Client)
	c.VAT = vat.NewService(&c.Client)

	return c
}

func (c *Client) SetDebug(debug bool) {
	c.Client.SetDebug(debug)
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.Client.SetDisallowUnknownFields(disallowUnknownFields)
}

func (c *Client) SetBaseURL(baseURL *url.URL) {
	// set base url for use in http client
	c.Client.SetBaseURL(baseURL)
}

func (c *Client) SetDivisionID(divisionID int) {
	// set base url for use in http client
	c.Client.SetDivisionID(divisionID)
}
