package purchaseentry

import (
	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/omitempty"
)

type NewPurchaseEntryLines []NewPurchaseEntryLine

type NewPurchaseEntryLine struct {
	ID                         edm.GUID     `json:"ID,omitempty"`                         // Primary key
	AmountFC                   edm.Double   `json:"AmountFC"`                             // Amount in the currency of the transaction
	Asset                      edm.GUID     `json:"Asset,omitempty"`                      // Reference to asset
	CostCenter                 edm.String   `json:"CostCenter,omitempty"`                 // Reference to cost center
	CostUnit                   edm.String   `json:"CostUnit,omitempty"`                   // Reference to cost unit
	Description                edm.String   `json:"Description,omitempty"`                // Description
	EntryID                    edm.GUID     `json:"EntryID,omitempty"`                              // Reference to header of the purchase entry
	From                       edm.DateTime `json:"From,omitempty"`                       // From date to identify the range for deferred costs. This is used in combination with the property 'To' that defines the end date
	GLAccount                  edm.GUID     `json:"GLAccount"`                            // General ledger account
	IntraStatArea              edm.String   `json:"IntraStatArea,omitempty"`              // IntraStat area (only relevant when IntraStat for purchase is enabled in the administration)
	IntraStatCountry           edm.String   `json:"IntraStatCountry,omitempty"`           // IntraStatCountry (only relevant when IntraStat for purchase is enabled in the administration)
	IntraStatDeliveryTerm      edm.String   `json:"IntraStatDeliveryTerm,omitempty"`      // IntraStat delivery term (only relevant when IntraStat for purchase is enabled in the administration)
	IntraStatTransactionA      edm.String   `json:"IntraStatTransactionA,omitempty"`      // IntraStat transaction A (only relevant when IntraStat for purchase is enabled in the administration)
	IntraStatTransactionB      edm.String   `json:"IntraStatTransactionB,omitempty"`      // IntraStat transaction B (only relevant when IntraStat for purchase is enabled in the Belgium, Luxembourg, France & United Kingdom administration)
	IntraStatTransportMethod   edm.String   `json:"IntraStatTransportMethod,omitempty"`   // IntraStat transport method (only relevant when IntraStat for purchase is enabled in the administration)
	Notes                      edm.String   `json:"Notes,omitempty"`                      // Extra remarks
	PrivateUsePercentage       edm.Double   `json:"PrivateUsePercentage,omitempty"`       // Percentage of re-invoice part of a cost to the owner of the company.
	Project                    edm.GUID     `json:"Project,omitempty"`                    // Reference to project
	Quantity                   edm.Double   `json:"Quantity,omitempty"`                   // Quantity
	SerialNumber               edm.String   `json:"SerialNumber,omitempty"`               // Serial number
	StatisticalNetWeight       edm.Double   `json:"StatisticalNetWeight,omitempty"`       // Statistical NetWeight (only relevant when IntraStat for purchase is enabled in the administration)
	StatisticalNumber          edm.String   `json:"StatisticalNumber,omitempty"`          // Statistical Number (only relevant when IntraStat for purchase is enabled in the administration)
	StatisticalQuantity        edm.Double   `json:"StatisticalQuantity,omitempty"`        // Statistical Quantity (only relevant when IntraStat for purchase is enabled in the administration)
	StatisticalValue           edm.Double   `json:"StatisticalValue,omitempty"`           // Statistical Value (only relevant when IntraStat for purchase is enabled in the administration)
	Subscription               edm.GUID     `json:"Subscription,omitempty"`               // Reference to subscription
	To                         edm.DateTime `json:"To,omitempty"`                         // To date to identify the range for deferred costs. This is used in combination with the property 'From' that defines the start date
	TrackingNumber             edm.GUID     `json:"TrackingNumber,omitempty"`             // Reference to tracking number
	VATAmountFC                edm.Double   `json:"VATAmountFC,omitempty"`                // VAT amount in the currency of the transaction. Use this property to specify a VAT amount that differs from the VAT amount that is automatically calculated.
	VATBaseAmountFC            edm.Double   `json:"VATBaseAmountFC,omitempty"`            // VAT base amount in the currency of the transaction
	VATCode                    edm.String   `json:"VATCode,omitempty"`                    // VAT code. If this property is not filled, it will use the default VAT code of the G/L account property
	VATNonDeductiblePercentage edm.Double   `json:"VATNonDeductiblePercentage,omitempty"` // If not the full amount of the VAT is deductible, you can indicate a percentage for the non decuctible part. This is used during the entry of purchase invoices.
	VATPercentage              edm.Double   `json:"VATPercentage,omitempty"`              // VAT percentage
	WithholdingAmountDC        edm.Double   `json:"WithholdingAmountDC,omitempty"`        // Withholding tax amount for spanish legislation
	WithholdingTax             edm.String   `json:"WithholdingTax,omitempty"`             // Withholding tax key for spanish legislation
}

func (l NewPurchaseEntryLine) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(l)
}
