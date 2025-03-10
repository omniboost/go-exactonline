package utils

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/aodin/date"
	"github.com/gorilla/schema"
	"github.com/omniboost/go-exactonline/odata"
)

type SchemaMarshaler interface {
	MarshalSchema() string
}

func AddQueryParamsToRequest(requestParams interface{}, req *http.Request, skipEmpty bool) error {
	params := url.Values{}
	encoder := schema.NewEncoder()

	// register custom encoders
	encodeSchemaMarshaler := func(v reflect.Value) string {
		marshaler, ok := v.Interface().(SchemaMarshaler)
		if ok == false {
			return ""
		}

		return marshaler.MarshalSchema()
	}

	// encodeInt := func(v reflect.Value) string {
	// 	i := int64(v.Int())
	// 	if i == 0 {
	// 		return ""
	// 	}
	// 	return strconv.FormatInt(i, 10)
	// }

	encoder.RegisterEncoder(&odata.Expand{}, encodeSchemaMarshaler)
	encoder.RegisterEncoder(&odata.Filter{}, encodeSchemaMarshaler)
	encoder.RegisterEncoder(&odata.Select{}, encodeSchemaMarshaler)
	encoder.RegisterEncoder(&odata.Top{}, encodeSchemaMarshaler)
	encoder.RegisterEncoder(&odata.OrderBy{}, encodeSchemaMarshaler)
	encoder.RegisterEncoder(&odata.Skip{}, encodeSchemaMarshaler)

	err := encoder.Encode(requestParams, params)
	if err != nil {
		return err
	}

	query := req.URL.Query()

	for k, vals := range params {
		for _, v := range vals {
			if skipEmpty && v == "" {
				continue
			}

			query.Add(k, v)
		}
	}

	req.URL.RawQuery = query.Encode()
	// force $ in query parameters
	req.URL.RawQuery = strings.Replace(req.URL.RawQuery, "%24", "$", -1)
	return nil
}

type DateNLNL struct {
	date.Date
}

func NewDateNLNL(year int, month time.Month, day int) DateNLNL {
	d := date.New(year, month, day)
	return DateNLNL{Date: d}
}

// func (d *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	t := time.Time(*d)
// 	return e.EncodeElement(t.Format("20060102"), start)
// }

func (d *DateNLNL) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// 28-1-2008
	layout := "2-1-2006"
	time, err := time.Parse(layout, value)
	d.Date = date.FromTime(time)
	return err
}
