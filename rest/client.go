package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/omniboost/go-exactonline/utils"
)

const (
	mediaType                 = "application/json"
	charset                   = "utf-8"
	customDescriptionLanguage = "EN-US"
)

var (
	rateLimits = make(map[int]*rateLimit)
)

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func New(http *http.Client) *Client {
	return &Client{
		http:                      http,
		customDescriptionLanguage: customDescriptionLanguage,
	}
}

type Client struct {
	// HTTP client used to communicate with the API.
	http *http.Client

	// Url pointing to base Exact Online API
	baseURL *url.URL

	divisionID int

	// Debugging flag
	debug bool

	disallowUnknownFields bool

	// Retrieve language sensitive properties such as descriptions in a specific language
	customDescriptionLanguage string

	// User agent for client
	userAgent string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback
}

func (c *Client) SetBaseURL(baseURL *url.URL) {
	// set base url for use in http client
	c.baseURL = baseURL
}

func (c *Client) SetDivisionID(divisionID int) {
	c.divisionID = divisionID
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	path = c.SubPath(path)
	u := c.GetEndpoint(path)

	var b io.Reader
	if body != nil {
		// determine if body is an io.Reader or should be serialized
		if r, ok := body.(io.Reader); ok {
			b = r
		} else {
			buf := new(bytes.Buffer)
			if body != nil {
				err := json.NewEncoder(buf).Encode(body)
				if err != nil {
					return nil, err
				}
			}
			b = buf
		}
	}

	req, err := http.NewRequest(method, u.String(), b)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", mediaType, charset))
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("CustomDescriptionLanguage", c.customDescriptionLanguage)
	return req, nil
}

func (c *Client) SubPath(path string) string {
	divisionID := strconv.Itoa(c.divisionID)
	path = strings.Replace(path, "{division}", divisionID, 1)
	path = strings.Replace(path, "{id}", "", 1)
	return path
}

func (c *Client) SubPathWithID(path string, id string) string {
	divisionID := strconv.Itoa(c.divisionID)
	path = strings.Replace(path, "{division}", divisionID, 1)

	if id == "" {
		path = strings.Replace(path, "{id}", id, 1)
	} else {
		id = fmt.Sprintf("(guid'%s')", id)
		path = strings.Replace(path, "{id}", id, 1)
	}
	return path
}

func (c *Client) GetEndpoint(path string) *url.URL {
	basePath := strings.TrimSuffix(c.baseURL.Path, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	u := *c.baseURL
	u.Path = basePath + path
	return &u
}

// Do sends an API request and returns the API response. The API response is XML decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	c.sleepUntilRequestLimit()

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	c.registerRequestLimit(httpResp)

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, err
	}

	if responseBody == nil {
		return httpResp, err
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// $top=1
	// {
	// 	"d" : {
	// 		"results" : [
	// 			{}
	// 		]
	// 	}
	// }

	// $top=2
	// {
	// 	"d" : [
	// 		{}
	// 	]
	// }

	// no $top
	// {
	// 	"d" : {
	// 		"results": [
	// 			{}
	// 		]
	// 	}
	// }

	type Envelope struct {
		D utils.JsonTester `json:"d"`
	}

	envelope := &Envelope{}
	dec := json.NewDecoder(httpResp.Body)
	if c.disallowUnknownFields {
		dec.DisallowUnknownFields()
	}

	err = dec.Decode(envelope)
	if err != nil {
		return httpResp, err
	}

	// get bytes
	b := []byte(envelope.D.RawMessage)

	// check if interface has ".Results" field
	r := reflect.ValueOf(responseBody)
	val := reflect.Indirect(r)
	hasResults := false
	if val.Kind() == reflect.Struct {
		field := val.FieldByName("Results")
		hasResults = field.IsValid()
	}

	// check if json is an object
	isArray := envelope.D.IsArray()

	// conversie doen
	if hasResults && isArray {
		b = append([]byte(`{"results":`), b...)
		b = append(b, []byte("}")...)

		err = json.Unmarshal(b, responseBody)
		return httpResp, err
	}

	err = json.Unmarshal(b, responseBody)
	return httpResp, err
}

type rateLimit struct {
	Limit     int
	Remaining int
	Reset     time.Time
}

func (c *Client) sleepUntilRequestLimit() {
	if rateLimits[c.divisionID] == nil {
		return
	}

	if rateLimits[c.divisionID].Remaining > 0 {
		return
	}

	sleep := rateLimits[c.divisionID].Reset.Sub(time.Now())
	time.Sleep(sleep)
}

func (c *Client) registerRequestLimit(req *http.Response) error {
	if rateLimits[c.divisionID] == nil {
		rateLimits[c.divisionID] = &rateLimit{}
	}

	var err error
	rateLimits[c.divisionID].Limit, err = strconv.Atoi(req.Header.Get("X-Ratelimit-Minutely-Limit"))
	if err != nil {
		return err
	}

	rateLimits[c.divisionID].Remaining, err = strconv.Atoi(req.Header.Get("X-Ratelimit-Minutely-Remaining"))
	if err != nil {
		return err
	}

	timestamp, err := strconv.ParseInt(req.Header.Get("X-Ratelimit-Minutely-Reset"), 10, 64)
	if err != nil {
		return err
	}

	rateLimits[c.divisionID].Reset = time.Unix(0, timestamp*int64(time.Millisecond))
	return nil
}
