package system

import "github.com/omniboost/go-exactonline/rest"

func NewService(rest *rest.Client) *Service {
	return &Service{rest: rest}
}

type Service struct {
	rest *rest.Client
}

// 0 for Inactive, 1 for Active and 2 for Archived Divisions
type DivisionStatus int32

// M=Male, V=Female, O=Unknown
type Gender string
