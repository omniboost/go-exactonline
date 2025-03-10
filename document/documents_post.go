package document

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/omitempty"
)

func (s *Service) DocumentsPost(body *DocumentsPostBody, ctx context.Context) (*DocumentsPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewDocumentsPostResponse()
	path := s.rest.SubPath(DocumentsEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type DocumentsPostBody NewDocument

func (d DocumentsPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(d)
}

func (s *Service) NewDocumentsPostBody() *DocumentsPostBody {
	return &DocumentsPostBody{}
}

func (s *Service) NewDocumentsPostResponse() *DocumentsPostResponse {
	return &DocumentsPostResponse{}
}

type DocumentsPostResponse struct {
	MetaData edm.MetaData `json:"__metadata"`
	Document
}

type NewDocument struct {
	ID                          edm.GUID     `json:"ID,omitempty"`                          // Primary key
	Account                     edm.GUID     `json:"Account,omitempty"`                     // ID of the related account of this document
	AmountFC                    edm.Double   `json:"AmountFC,omitempty"`                    // Amount in the currency of the transaction
	Body                        edm.String   `json:"Body,omitempty"`                        // Body of this document
	Category                    edm.GUID     `json:"Category,omitempty"`                    // ID of the category of this document
	CategoryDescription         edm.String   `json:"CategoryDescription,omitempty"`         // Description of Category
	Contact                     edm.GUID     `json:"Contact,omitempty"`                     // ID of the related contact of this document
	CreatorFullName             edm.String   `json:"CreatorFullName,omitempty"`             // Name of creator
	Currency                    edm.String   `json:"Currency,omitempty"`                    // Currency code
	DocumentDate                edm.DateTime `json:"DocumentDate,omitempty"`                // Entry date of the incoming document
	DocumentFolder              edm.GUID     `json:"DocumentFolder,omitempty"`              // The Id of document folder
	ExpiryDate                  edm.DateTime `json:"ExpiryDate,omitempty"`                  // Expiry date of this document
	FinancialTransactionEntryID edm.GUID     `json:"FinancialTransactionEntryID,omitempty"` // Reference to the transaction lines of the financial entry. For a document of type sales invoice it will return the InvoiceID of the sales invoice (SalesInvoices API).
	HID                         edm.Int32    `json:"HID,omitempty"`                         // Human-readable ID, formatted as xx.xxx.xxx. Unique. May not be equal to zero
	InheritShare                edm.Boolean  `json:"InheritShare,omitempty"`                // InheritShare value
	Language                    edm.String   `json:"Language,omitempty"`                    // The language code of the document
	Opportunity                 edm.GUID     `json:"Opportunity,omitempty"`                 // The opportunity linked to the document
	Project                     edm.GUID     `json:"Project,omitempty"`                     // The project linked to the document
	SalesInvoiceNumber          edm.Int32    `json:"SalesInvoiceNumber,omitempty"`          // 'Our reference' of the transaction that belongs to this document
	SalesOrderNumber            edm.Int32    `json:"SalesOrderNumber,omitempty"`            // Number of the sales order
	ShopOrderNumber             edm.Int32    `json:"ShopOrderNumber,omitempty"`             // Number of the shop order
	Subject                     edm.String   `json:"Subject"`                               // Subject of this document
	Type                        edm.Int32    `json:"Type"`                                  // ID of the type of this document
}

func (d NewDocument) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(d)
}
