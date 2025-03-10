package document

import "github.com/omniboost/go-exactonline/edm"

type Documents []Document

type Document struct {
	ID                          edm.GUID     `json:"ID,omitempty"`                // Primary key
	Account                     edm.GUID     `json:"Account"`                     // ID of the related account of this document
	AccountCode                 edm.String   `json:"AccountCode"`                 // Code of Account
	AccountName                 edm.String   `json:"AccountName"`                 // Name of Account
	AmountFC                    edm.Double   `json:"AmountFC"`                    // Amount in the currency of the transaction
	Body                        edm.String   `json:"Body"`                        // Body of this document
	Category                    edm.GUID     `json:"Category"`                    // ID of the category of this document
	CategoryDescription         edm.String   `json:"CategoryDescription"`         // Description of Category
	Contact                     edm.GUID     `json:"Contact"`                     // ID of the related contact of this document
	ContactFullName             edm.String   `json:"ContactFullName"`             // Contact full name
	Created                     edm.DateTime `json:"Created"`                     // Creation date
	Creator                     edm.GUID     `json:"Creator"`                     // User ID of creator
	CreatorFullName             edm.String   `json:"CreatorFullName"`             // Name of creator
	Currency                    edm.String   `json:"Currency"`                    // Currency code
	Division                    edm.Int32    `json:"Division"`                    // Division code
	DocumentDate                edm.DateTime `json:"DocumentDate"`                // Entry date of the incoming document
	DocumentFolder              edm.GUID     `json:"DocumentFolder"`              // The Id of document folder
	DocumentFolderCode          edm.String   `json:"DocumentFolderCode"`          // The Code of document folder
	DocumentFolderDescription   edm.String   `json:"DocumentFolderDescription"`   // The Decsription of document folder
	DocumentViewUrl             edm.String   `json:"DocumentViewUrl"`             // Url to view the document
	ExpiryDate                  edm.DateTime `json:"ExpiryDate"`                  // Expiry date of this document
	FinancialTransactionEntryID edm.GUID     `json:"FinancialTransactionEntryID"` // Reference to the transaction lines of the financial entry. For a document of type sales invoice it will return the InvoiceID of the sales invoice (SalesInvoices API).
	HasEmptyBody                edm.Boolean  `json:"HasEmptyBody"`                // Indicates that the document body is empty
	HID                         edm.Int32    `json:"HID"`                         // Human-readable ID, formatted as xx.xxx.xxx. Unique. May not be equal to zero
	InheritShare                edm.Boolean  `json:"InheritShare"`                // InheritShare value
	Language                    edm.String   `json:"Language"`                    // The language code of the document
	Modified                    edm.DateTime `json:"Modified"`                    // Last modified date
	Modifier                    edm.GUID     `json:"Modifier"`                    // User ID of modifier
	ModifierFullName            edm.String   `json:"ModifierFullName"`            // Name of modifier
	Opportunity                 edm.GUID     `json:"Opportunity"`                 // The opportunity linked to the document
	Project                     edm.GUID     `json:"Project"`                     // The project linked to the document
	ProjectCode                 edm.String   `json:"ProjectCode"`                 // Code of project
	ProjectDescription          edm.String   `json:"ProjectDescription"`          // Description of project
	SalesInvoiceNumber          edm.Int32    `json:"SalesInvoiceNumber"`          // 'Our reference' of the transaction that belongs to this document
	SalesOrderNumber            edm.Int32    `json:"SalesOrderNumber"`            // Number of the sales order
	SendMethod                  edm.Int32    `json:"SendMethod"`                  // Send method
	ShopOrderNumber             edm.Int32    `json:"ShopOrderNumber"`             // Number of the shop order
	Subject                     edm.String   `json:"Subject"`                     // Subject of this document
	Type                        edm.Int32    `json:"Type"`                        // ID of the type of this document
	TypeDescription             edm.String   `json:"TypeDescription"`             // Description of Type debugger eval code:11:9
}

type DocumentAttachments []DocumentAttachment

type DocumentAttachment struct {
	ID         edm.GUID   `json:"ID,omitempty"` // Primary key
	Attachment edm.Binary `json:"Attachment"`   // Contains the attachment(Format: Base64 encoded)
	Document   edm.GUID   `json:"Document"`     // Reference to the Document
	FileName   edm.String `json:"FileName"`     // Filename of the attachment
	FileSize   edm.Double `json:"FileSize"`     // File size of the attachment
	Url        edm.String `json:"Url"`          // Url of the attachment. To get the file in its original format (xml, jpg, pdf, etc.) append &Download=1 to the url.
}

type DocumentTypeCategories []DocumentTypeCategory

type DocumentTypeCategory struct {
	ID          edm.Int32    `json:"ID,omitempty"`          // Primary key
	Created     edm.DateTime `json:"Created,omitempty"`     // Creation date
	Description edm.String   `json:"Description,omitempty"` // Document category type description
	Modified    edm.DateTime `json:"Modified,omitempty"`    // Last modified date debugger
}

type DocumentCategories []DocumentCategory

type DocumentCategory struct {
	ID          edm.GUID     `json:"ID,omitempty"`          // Primary key
	Created     edm.DateTime `json:"Created,omitempty"`     // Creation date
	Description edm.String   `json:"Description,omitempty"` // Document category description
	Modified    edm.DateTime `json:"Modified,omitempty"`    // Last modified date debugger eval code:12:9
}

type DocumentTypes []DocumentType

type DocumentType struct {
	ID                  edm.Int32    `json:"ID,omitempty"`                  // Primary key
	Created             edm.DateTime `json:"Created,omitempty"`             // Creation date
	Description         edm.String   `json:"Description,omitempty"`         // Document type description
	DocumentIsCreatable edm.Boolean  `json:"DocumentIsCreatable,omitempty"` // Indicates if documents of this type can be created
	DocumentIsDeletable edm.Boolean  `json:"DocumentIsDeletable,omitempty"` // Indicates if documents of this type can be deleted
	DocumentIsUpdatable edm.Boolean  `json:"DocumentIsUpdatable,omitempty"` // Indicates if documents of this type can be updated
	DocumentIsViewable  edm.Boolean  `json:"DocumentIsViewable,omitempty"`  // Indicates if documents of this type can be retrieved
	Modified            edm.DateTime `json:"Modified,omitempty"`            // Last modified date
	TypeCategory        edm.Int32    `json:"TypeCategory,omitempty"`        // ID of the document type category
}
