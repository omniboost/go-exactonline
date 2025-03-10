package payroll

import "github.com/omniboost/go-exactonline/edm"

type Employees []Employee

type Employee struct {
	ID                        edm.GUID     `json:"ID,omitempty"`                        // Primary key
	ActiveEmployment          edm.Byte     `json:"ActiveEmployment,omitempty"`          // Obsolete
	AddressLine2              edm.String   `json:"AddressLine2,omitempty"`              // Second address line
	AddressLine3              edm.String   `json:"AddressLine3,omitempty"`              // Third address line
	AddressStreet             edm.String   `json:"AddressStreet,omitempty"`             // Street of address
	AddressStreetNumber       edm.String   `json:"AddressStreetNumber,omitempty"`       // Street number of address
	AddressStreetNumberSuffix edm.String   `json:"AddressStreetNumberSuffix,omitempty"` // Street number suffix of address
	BirthDate                 edm.DateTime `json:"BirthDate,omitempty"`                 // Birth date
	BirthName                 edm.String   `json:"BirthName,omitempty"`                 // Birth name
	BirthNamePrefix           edm.String   `json:"BirthNamePrefix,omitempty"`           // Birth middle name
	BirthPlace                edm.String   `json:"BirthPlace,omitempty"`                // Birth place
	BusinessEmail             edm.String   `json:"BusinessEmail,omitempty"`             // Email of the employee at the office
	BusinessFax               edm.String   `json:"BusinessFax,omitempty"`               // Fax number of the employee at the office
	BusinessMobile            edm.String   `json:"BusinessMobile,omitempty"`            // Office mobile number of the employee
	BusinessPhone             edm.String   `json:"BusinessPhone,omitempty"`             // Phone number of the employee at the office
	BusinessPhoneExtension    edm.String   `json:"BusinessPhoneExtension,omitempty"`    // Phone extension of the employee at the office
	CASONumber                edm.String   `json:"CASONumber,omitempty"`                // Obsolete
	City                      edm.String   `json:"City,omitempty"`                      // City
	Code                      edm.String   `json:"Code,omitempty"`                      // Code of the employee
	Country                   edm.String   `json:"Country,omitempty"`                   // Country code
	Created                   edm.DateTime `json:"Created,omitempty"`                   // Creation date
	Creator                   edm.GUID     `json:"Creator,omitempty"`                   // User ID of creator
	CreatorFullName           edm.String   `json:"CreatorFullName,omitempty"`           // Name of creator
	Customer                  edm.GUID     `json:"Customer,omitempty"`                  // Customer ID
	Division                  edm.Int32    `json:"Division,omitempty"`                  // Division code
	Email                     edm.String   `json:"Email,omitempty"`                     // Email address
	EmployeeHID               edm.Int32    `json:"EmployeeHID,omitempty"`               // Employee number
	EndDate                   edm.DateTime `json:"EndDate,omitempty"`                   // End date of the employee
	FirstName                 edm.String   `json:"FirstName,omitempty"`                 // First name of the employee
	FullName                  edm.String   `json:"FullName,omitempty"`                  // Full name of the employee
	Gender                    edm.String   `json:"Gender,omitempty"`                    // Gender
	HID                       edm.Int32    `json:"HID,omitempty"`                       // Numeric ID of the employee
	Initials                  edm.String   `json:"Initials,omitempty"`                  // Initials
	IsActive                  edm.Boolean  `json:"IsActive,omitempty"`                  // IsActive
	IsAnonymised              edm.Byte     `json:"IsAnonymised,omitempty"`              // Indicates whether the employee is anonymised.
	Language                  edm.String   `json:"Language,omitempty"`                  // Language code
	LastName                  edm.String   `json:"LastName,omitempty"`                  // Last name
	LocationDescription       edm.String   `json:"LocationDescription,omitempty"`       // Description of the location of the employee (where am I?)
	Manager                   edm.GUID     `json:"Manager,omitempty"`                   // Direct manager of the employee
	MaritalDate               edm.DateTime `json:"MaritalDate,omitempty"`               // Date of marriage
	MaritalStatus             edm.Int16    `json:"MaritalStatus,omitempty"`             // Marital status
	MiddleName                edm.String   `json:"MiddleName,omitempty"`                // Middle name
	Mobile                    edm.String   `json:"Mobile,omitempty"`                    // Mobile phone
	Modified                  edm.DateTime `json:"Modified,omitempty"`                  // Last modified date
	Modifier                  edm.GUID     `json:"Modifier,omitempty"`                  // User ID of modifier
	ModifierFullName          edm.String   `json:"ModifierFullName,omitempty"`          // Name of modifier
	Municipality              edm.String   `json:"Municipality,omitempty"`              // Municipality
	NameComposition           edm.Int16    `json:"NameComposition,omitempty"`           //
	Nationality               edm.String   `json:"Nationality,omitempty"`               // Nationality
	NickName                  edm.String   `json:"NickName,omitempty"`                  // Nick name
	Notes                     edm.String   `json:"Notes,omitempty"`                     // Additional notes
	PartnerName               edm.String   `json:"PartnerName,omitempty"`               // Name of partner
	PartnerNamePrefix         edm.String   `json:"PartnerNamePrefix,omitempty"`         // Middle name of partner
	Person                    edm.GUID     `json:"Person,omitempty"`                    // Reference to the persons table in which the personal data of the employee is stored
	Phone                     edm.String   `json:"Phone,omitempty"`                     // Phone number
	PhoneExtension            edm.String   `json:"PhoneExtension,omitempty"`            // Phone number extension
	PictureFileName           edm.String   `json:"PictureFileName,omitempty"`           // Filename of picture
	PictureUrl                edm.String   `json:"PictureUrl,omitempty"`                // Url of picture
	Postcode                  edm.String   `json:"Postcode,omitempty"`                  // Postcode
	PrivateEmail              edm.String   `json:"PrivateEmail,omitempty"`              // Private email address
	SocialSecurityNumber      edm.String   `json:"SocialSecurityNumber,omitempty"`      // Social security number
	StartDate                 edm.DateTime `json:"StartDate,omitempty"`                 // Start date of the employee
	State                     edm.String   `json:"State,omitempty"`                     // State
	Title                     edm.String   `json:"Title,omitempty"`                     // Title
	User                      edm.GUID     `json:"User,omitempty"`                      // User ID of employee
	UserFullName              edm.String   `json:"UserFullName,omitempty"`              // Name of user
}
