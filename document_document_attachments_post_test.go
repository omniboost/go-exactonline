package exact_test

import (
	"encoding/json"
	"log"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/omniboost/go-exactonline/document"
	"github.com/omniboost/go-exactonline/edm"
)

func TestDocumentAttachmentsPost(t *testing.T) {
	doc, _ := uuid.FromString("c9d49de9-3142-4d29-b163-6a8b8b1979b6")
	body := document.DocumentAttachmentsPostBody{
		Attachment: edm.Binary([]byte("TEST")),
		Document:   edm.GUID{doc},
		FileName:   "test.txt",
	}
	resp, err := client.Document.DocumentAttachmentsPost(&body, nil)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
