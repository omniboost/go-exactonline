package project

import "github.com/omniboost/go-exactonline/rest"

func NewService(rest *rest.Client) *Service {
	return &Service{rest: rest}
}

type Service struct {
	rest *rest.Client
}
