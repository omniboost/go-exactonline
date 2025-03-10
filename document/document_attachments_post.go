package document

import (
	"context"
	"net/http"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/omitempty"
)

func (s *Service) DocumentAttachmentsPost(body *DocumentAttachmentsPostBody, ctx context.Context) (*DocumentAttachmentsPostResponse, error) {
	method := http.MethodPost
	responseBody := s.NewDocumentAttachmentsPostResponse()
	path := s.rest.SubPath(DocumentAttachmentsEndpoint)

	// create a new HTTP request
	httpReq, err := s.rest.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.rest.Do(httpReq, responseBody)
	return responseBody, err
}

type DocumentAttachmentsPostBody NewDocumentAttachment

func (d DocumentAttachmentsPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(d)
}

func (s *Service) NewDocumentAttachmentsPostBody() *DocumentAttachmentsPostBody {
	return &DocumentAttachmentsPostBody{}
}

func (s *Service) NewDocumentAttachmentsPostResponse() *DocumentAttachmentsPostResponse {
	return &DocumentAttachmentsPostResponse{}
}

type DocumentAttachmentsPostResponse struct {
	MetaData edm.MetaData `json:"__metadata"`
	DocumentAttachment
}

type NewDocumentAttachment struct {
	ID         edm.GUID   `json:"ID,omitempty"` // Primary key
	Attachment edm.Binary `json:"Attachment"`   // Contains the attachment(Format: Base64 encoded)
	Document   edm.GUID   `json:"Document"`     // Reference to the Document
	FileName   edm.String `json:"FileName"`     // Filename of the attachment
}

func (d NewDocumentAttachment) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(d)
}
