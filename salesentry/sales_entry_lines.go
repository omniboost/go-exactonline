package salesentry

import (
	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/omitempty"
)

type NewSalesEntryLines []NewSalesEntryLine

type NewSalesEntryLine struct {
	ID                       edm.GUID     `json:"ID,omitempty"`             // Primary key
	AmountFC                 edm.Double   `json:"AmountFC"`                 // For normal lines it's the amount excluding VAT
	Asset                    edm.GUID     `json:"Asset"`                    // Reference to Asset
	CostCenter               edm.String   `json:"CostCenter"`               // Reference to Cost center
	CostUnit                 edm.String   `json:"CostUnit"`                 // Reference to Cost unit
	Description              edm.String   `json:"Description"`              // Description. Can be different for header and lines
	EntryID                  edm.GUID     `json:"EntryID,omitempty"`        // The unique ID of the entry. Via this ID all transaction lines of a single entry can be retrieved
	ExtraDutyAmountFC        edm.Double   `json:"ExtraDutyAmountFC"`        // Extra duty amount in the currency of the transaction. Both extra duty amount and VAT amount need to be specified in order to differ this property from automatically calculated.
	ExtraDutyPercentage      edm.Double   `json:"ExtraDutyPercentage"`      // Extra duty percentage for the item
	From                     edm.DateTime `json:"From,omitempty"`           // From date for deferred revenue
	GLAccount                edm.GUID     `json:"GLAccount"`                // The GL Account of the invoice line. This field is generated based on the revenue account of the item (or the related item group). G/L Account is also used to determine whether the costcenter / costunit is mandatory
	IntraStatArea            edm.String   `json:"IntraStatArea"`            // IntraStat area
	IntraStatCountry         edm.String   `json:"IntraStatCountry"`         // IntraStatCountry
	IntraStatDeliveryTerm    edm.String   `json:"IntraStatDeliveryTerm"`    // IntraStat delivery term
	IntraStatTransactionA    edm.String   `json:"IntraStatTransactionA"`    // IntraStat transaction a
	IntraStatTransportMethod edm.String   `json:"IntraStatTransportMethod"` // IntraStat transport method
	Notes                    edm.String   `json:"Notes"`                    // Extra notes
	Project                  edm.GUID     `json:"Project"`                  // The project to which the sales transaction line is linked. The project can be different per line. Sometimes also the project in the header is filled although this is not really used
	Quantity                 edm.Double   `json:"Quantity"`                 // The number of items sold in default units. The quantity shown in the entry screen is Quantity * UnitFactor
	SerialNumber             edm.String   `json:"SerialNumber"`             // Serial number
	StatisticalNetWeight     edm.Double   `json:"StatisticalNetWeight"`     // Statistical NetWeight
	StatisticalNumber        edm.String   `json:"StatisticalNumber"`        // Statistical Number
	StatisticalQuantity      edm.Double   `json:"StatisticalQuantity"`      // Statistical Quantity
	StatisticalValue         edm.Double   `json:"StatisticalValue"`         // Statistical Value
	Subscription             edm.GUID     `json:"Subscription"`             // When generating invoices from subscriptions, this field records the link between invoice lines and subscription lines
	TaxSchedule              edm.GUID     `json:"TaxSchedule"`              // Obsolete
	To                       edm.DateTime `json:"To,omitempty"`             // To date for deferred revenue
	TrackingNumber           edm.GUID     `json:"TrackingNumber"`           // Reference to TrackingNumber
	VATAmountFC              edm.Double   `json:"VATAmountFC"`              // VAT amount in the currency of the transaction. Use this property to specify a VAT amount that differs from the VAT amount that is automatically calculated. However if the transaction uses extra duty, extra duty amount also needs to be specified.
	VATBaseAmountDC          edm.Double   `json:"VATBaseAmountDC"`          // The VAT base amount in the default currency of the company. This is calculated based on the VATBaseAmountFC
	VATBaseAmountFC          edm.Double   `json:"VATBaseAmountFC"`          // The VAT base amount in invoice currency. This is calculated with the use of VAT codes. It's an internal value
	VATCode                  edm.String   `json:"VATCode"`                  // The VAT code used when the invoice was registered
	VATPercentage            edm.Double   `json:"VATPercentage,omitempty"`  // The VAT percentage of the VAT code. This is the percentage at the moment the invoice is created. It's also used by the default calculation of VAT amounts and VAT base amounts
}

func (l NewSalesEntryLine) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(l)
}
