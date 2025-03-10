package general

import "github.com/omniboost/go-exactonline/edm"

type Currency struct {
	Code            edm.String   `json:"Code"`            // Primary key
	AmountPrecision edm.Double   `json:"AmountPrecision"` // Amount precision
	Created         edm.DateTime `json:"Created"`         // Creation date
	Description     edm.String   `json:"Description"`     // Description of the currency
	Modified        edm.DateTime `json:"Modified"`        // Last modified date
	PricePrecision  edm.Double   `json:"PricePrecision"`  // Price precision
}

type Currencies []Currency
