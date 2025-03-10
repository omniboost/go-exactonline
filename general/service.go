package general

import "github.com/omniboost/go-exactonline/rest"

func NewService(rest *rest.Client) *Service {
	service := &Service{rest: rest}
	service.Currencies = NewCurrenciesResource(rest)
	return service
}

type Service struct {
	Currencies *CurrenciesResource

	rest *rest.Client
}
