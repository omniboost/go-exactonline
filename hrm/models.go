package hrm

import "github.com/omniboost/go-exactonline/edm"

type Costcenters []Costcenter

type Costcenter struct {
	ID               edm.GUID     `json:"ID"`               // Primary key
	Active           edm.Boolean  `json:"Active"`           // Indicates if the cost center is active: 0 = inactive 1 = active
	Code             edm.String   `json:"Code"`             // Code (user-defined ID)
	Created          edm.DateTime `json:"Created"`          // Creation date
	Creator          edm.GUID     `json:"Creator"`          // User ID of creator
	CreatorFullName  edm.String   `json:"CreatorFullName"`  // Name of creator
	Description      edm.String   `json:"Description"`      // Description (text)
	Division         edm.Int32    `json:"Division"`         // Division code
	EndDate          edm.DateTime `json:"EndDate"`          // The end date by which the cost center has to be inactive
	Modified         edm.DateTime `json:"Modified"`         // Last modified date
	Modifier         edm.GUID     `json:"Modifier"`         // User ID of modifier
	ModifierFullName edm.String   `json:"ModifierFullName"` // Name of modifier
}

type Costunits []Costunit

type Costunit struct {
	ID               edm.GUID     `json:"ID"`               // Primary key
	Code             edm.String   `json:"Code"`             // Code (user-defined ID)
	Created          edm.DateTime `json:"Created"`          // Creation date
	Creator          edm.GUID     `json:"Creator"`          // User ID of creator
	CreatorFullName  edm.String   `json:"CreatorFullName"`  // Name of creator
	Description      edm.String   `json:"Description"`      // Description (text)
	Division         edm.Int32    `json:"Division"`         // Division code
	EndDate          edm.DateTime `json:"EndDate"`          // The end date by which the cost center has to be inactive
	Modified         edm.DateTime `json:"Modified"`         // Last modified date
	Modifier         edm.GUID     `json:"Modifier"`         // User ID of modifier
	ModifierFullName edm.String   `json:"ModifierFullName"` // Name of modifier
}
