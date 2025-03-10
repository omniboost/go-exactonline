package vat

import (
	"encoding/json"
	"strings"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/utils"
)

type VatCodes []VatCode

type VatCode struct {
	ID                            edm.GUID       `json:"ID"`                            // Primary key
	Account                       edm.GUID       `json:"Account"`                       // Tax account
	AccountCode                   edm.String     `json:"AccountCode"`                   // Code of Account
	AccountName                   edm.String     `json:"AccountName"`                   // Name of Account
	CalculationBasis              edm.Byte       `json:"CalculationBasis"`              // Indicates how to calculate the tax. 0 = based on the gross amount, 1 = based on the gross amount + another tax
	Charged                       edm.Boolean    `json:"Charged"`                       // Indicates if transactions using the VAT code are transactions of the domestic VAT charging regulation (such as those for subcontractors) or transactions that are registered within the EU. If Charged=1 and linked to a purchase invoice, both a line for the VAT to pay and a line for the VAT to claim are being created
	Code                          edm.String     `json:"Code"`                          // VAT code
	Country                       edm.String     `json:"Country"`                       // Obsolete
	Created                       edm.DateTime   `json:"Created"`                       // Creation date
	Creator                       edm.GUID       `json:"Creator"`                       // User ID of creator
	CreatorFullName               edm.String     `json:"CreatorFullName"`               // Name of creator
	Description                   edm.String     `json:"Description"`                   // Description of the VAT code
	Division                      edm.Int32      `json:"Division"`                      // Division code
	EUSalesListing                edm.String     `json:"EUSalesListing"`                // Used in all legislations except France. Indicates if and how transactions using the VAT code appear on the ICT return (EU sales list). L = Listing goods, N = No listing, S = Listing services, T = Triangulation
	GLDiscountPurchase            edm.GUID       `json:"GLDiscountPurchase"`            // Indicates the purchase discount GL account linked to the VAT codes for German legislation
	GLDiscountPurchaseCode        edm.String     `json:"GLDiscountPurchaseCode"`        // Code of GLDiscountPurchase
	GLDiscountPurchaseDescription edm.String     `json:"GLDiscountPurchaseDescription"` // Description of GLDiscountPurchase
	GLDiscountSales               edm.GUID       `json:"GLDiscountSales"`               // Indicates the sales discount GL account linked to the VAT codes for German legislation
	GLDiscountSalesCode           edm.String     `json:"GLDiscountSalesCode"`           // Code of GLDiscountSales
	GLDiscountSalesDescription    edm.String     `json:"GLDiscountSalesDescription"`    // Description of GLDiscountSales
	GLToClaim                     edm.GUID       `json:"GLToClaim"`                     // G/L account that is used to book the VAT to claim. If you enter purchases with a VAT code, the VAT amount to be claimed is entered to this VAT account. Must be of type VAT
	GLToClaimCode                 edm.String     `json:"GLToClaimCode"`                 // Code of GLToClaim
	GLToClaimDescription          edm.String     `json:"GLToClaimDescription"`          // Description of GLToClaim
	GLToPay                       edm.GUID       `json:"GLToPay"`                       // G/L account that is used to book the VAT to pay. If you enter sales with a VAT code, the VAT amount to be paid is entered to this VAT account. Must be of type VAT
	GLToPayCode                   edm.String     `json:"GLToPayCode"`                   // Code of GLToPay
	GLToPayDescription            edm.String     `json:"GLToPayDescription"`            // Description of GLToPay
	IntraStat                     edm.Boolean    `json:"IntraStat"`                     // Used in all legislations except France. Indicates if intrastat is used
	IntrastatType                 edm.String     `json:"IntrastatType"`                 // Used in France legislation only. Indicates if and how transactions using the VAT code appear on the DEB/DES return. L = Goods, N = Empty, S = Services
	IsBlocked                     edm.Boolean    `json:"IsBlocked"`                     // Indicates if the VAT code may still be used
	LegalText                     edm.String     `json:"LegalText"`                     // Legal description for VAT code to print in the total block of the invoice
	Modified                      edm.DateTime   `json:"Modified"`                      // Last modified date
	Modifier                      edm.GUID       `json:"Modifier"`                      // User ID of modifier
	ModifierFullName              edm.String     `json:"ModifierFullName"`              // User name of modifier
	Percentage                    edm.Double     `json:"Percentage"`                    // Percentage of the VAT code
	TaxReturnType                 edm.Int16      `json:"TaxReturnType"`                 // Indicates what type of Taxcode it is: can be VAT, IncomeTax
	Type                          Type           `json:"Type"`                          // Indicates how the VAT amount should be calculated in relation to the invoice amount. B = VAT 0% (Only base amount), E = Excluding, I = Including, N = No VAT
	VatDocType                    edm.String     `json:"VatDocType"`                    // Field in VAT code maintenance to calculate different VATs depending on the selected document type. P = purchase invoice, F = freelance invoice, E = expense voucher. The field is valid for witholding tax type
	VatMargin                     edm.Byte       `json:"VatMargin"`                     // The VAT margin scheme is used for the trade of secondhand goods which are purchased without VAT (for example when a company buys a secondhand good from a private person). In the VAT margin scheme, the VAT is not calculated based on the sales price. Instead of that, the VAT is calculated based on the margin (gross sales price minus the gross purchase price)
	VATPartialRatio               edm.Int16      `json:"VATPartialRatio"`               // Partial ratio explains which part of the VAT the company has to pay. Used in some branches where the sellers have a bad reputation, so the buyers have to take over the VAT-liability
	VATPercentages                VATPercentages `json:"VATPercentages"`                // VAT percentages. You can have several VAT percentages, with start and end dates
	VATTransactionType            edm.String     `json:"VATTransactionType"`            // Indicates the type of transactions for which the VAT code may be used. B = Both, P = Purchase, S = Sales
}

func (c *VatCode) UnmarshalJSON(data []byte) error {
	type Alias VatCode
	alias := Alias(*c)
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	*c = VatCode(alias)
	c.Code = edm.String(strings.TrimSpace(c.Code.String()))
	c.AccountCode = edm.String(strings.TrimSpace(c.AccountCode.String()))
	return nil
}

type VATPercentages []VATPercentage

// standalone: "SalesInvoiceLines": []
// deferred: "SalesInvoiceLines": {"__deferred": {}}
// embedded: "SalesInvoiceLines": {"results": []}
func (p *VATPercentages) UnmarshalJSON(data []byte) error {
	type Results VATPercentages

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

		*p = VATPercentages(*results)
		return nil
	}

	// lines are embedded or deferred: only try embedded
	envelope := &Envelope{Results: Results(*p)}
	err = json.Unmarshal(data, envelope)
	if err != nil {
		return err
	}

	*p = VATPercentages(envelope.Results)
	return nil
}

type VATPercentage struct {
	ID               edm.GUID     `json:"ID"`               // Primary key
	Created          edm.DateTime `json:"Created"`          // Creation date
	Creator          edm.GUID     `json:"Creator"`          // User ID of creator
	CreatorFullName  edm.String   `json:"CreatorFullName"`  // Name of creator
	Division         edm.Int32    `json:"Division"`         // Division code
	EndDate          edm.DateTime `json:"EndDate"`          // End date of the date range during which this percentage is valid
	LineNumber       edm.Int32    `json:"LineNumber"`       // Line number
	Modified         edm.DateTime `json:"Modified"`         // Last modified date
	Modifier         edm.GUID     `json:"Modifier"`         // User ID of modifier
	ModifierFullName edm.String   `json:"ModifierFullName"` // Name of modifier
	Percentage       edm.Double   `json:"Percentage"`       // Percentage
	StartDate        edm.DateTime `json:"StartDate"`        // Start date of the date range during which this percentage is valid
	Type             edm.Int16    `json:"Type"`             // 0 = Normal, 1 = Extra duty
	VATCodeID        edm.GUID     `json:"VATCodeID"`        // VAT code
}

type Type edm.String

const (
	TypeBase      Type = "B"
	TypeExcluding Type = "E"
	TypeIncluding Type = "I"
	TypeNo        Type = "N"
)
