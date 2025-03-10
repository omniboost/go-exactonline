package purchaseentry

import "github.com/omniboost/go-exactonline/rest"

const (
	PurchaseEntriesEndpoint = "/v1/{division}/purchaseentry/PurchaseEntries"
)

func NewService(rest *rest.Client) *Service {
	return &Service{rest: rest}
}

type Service struct {
	rest *rest.Client
}
