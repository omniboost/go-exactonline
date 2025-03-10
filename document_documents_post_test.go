package exact_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/omniboost/go-exactonline/document"
	"github.com/omniboost/go-exactonline/edm"
)

func TestDocumentsPost(t *testing.T) {
	cat, _ := uuid.FromString("cf507b3a-cfea-45d8-9695-4d49964ddada")
	account, _ := uuid.FromString("44b9a2d3-7fe5-45e6-9f3c-01efbf40ec2d")
	body := document.DocumentsPostBody{
		Category: edm.GUID{cat},
		Type:     10,
		Account:  edm.GUID{account},
		// AmountFC                    edm.Double   `json:"AmountFC,omitempty"`                    // Amount in the currency of the transaction
		// Body                        edm.String   `json:"Body,omitempty"`                        // Body of this document
		// Category                    edm.GUID     `json:"Category,omitempty"`                    // ID of the category of this document
		// CategoryDescription         edm.String   `json:"CategoryDescription,omitempty"`         // Description of Category
		// Contact                     edm.GUID     `json:"Contact,omitempty"`                     // ID of the related contact of this document
		// CreatorFullName             edm.String   `json:"CreatorFullName,omitempty"`             // Name of creator
		// Currency                    edm.String   `json:"Currency,omitempty"`                    // Currency code
		DocumentDate: edm.DateTime{time.Now()},
		// DocumentFolder              edm.GUID     `json:"DocumentFolder,omitempty"`              // The Id of document folder
		// ExpiryDate                  edm.DateTime `json:"ExpiryDate,omitempty"`                  // Expiry date of this document
		FinancialTransactionEntryID: edm.GUID{cat},
		// HID                         edm.Int32    `json:"HID,omitempty"`                         // Human-readable ID, formatted as xx.xxx.xxx. Unique. May not be equal to zero
		// InheritShare                edm.Boolean  `json:"InheritShare,omitempty"`                // InheritShare value
		// Language                    edm.String   `json:"Language,omitempty"`                    // The language code of the document
		// Opportunity                 edm.GUID     `json:"Opportunity,omitempty"`                 // The opportunity linked to the document
		// Project                     edm.GUID     `json:"Project,omitempty"`                     // The project linked to the document
		// SalesInvoiceNumber          edm.Int32    `json:"SalesInvoiceNumber,omitempty"`          // 'Our reference' of the transaction that belongs to this document
		// SalesOrderNumber            edm.Int32    `json:"SalesOrderNumber,omitempty"`            // Number of the sales order
		// ShopOrderNumber             edm.Int32    `json:"ShopOrderNumber,omitempty"`             // Number of the shop order
		Subject: "TEST",
	}
	resp, err := client.Document.DocumentsPost(&body, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
