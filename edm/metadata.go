package edm

import "github.com/omniboost/go-exactonline/utils"

type MetaData struct {
	URL  utils.URL `json:"uri"`
	Type string    `json:"type"`
}
