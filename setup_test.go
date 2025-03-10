package exact_test

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"testing"

	exact "github.com/omniboost/go-exactonline"
	"golang.org/x/oauth2"
)

var (
	client *exact.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	tokenURL := os.Getenv("TOKEN_URL")
	divisionIDString := os.Getenv("DIVISION_ID")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL
	var divisionID int

	if divisionIDString != "" {
		divisionID, err = strconv.Atoi(divisionIDString)
		if err != nil {
			log.Fatal(err)
		}
	}

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	oauthConfig := exact.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(oauth2.NoContext, token)

	client = exact.NewClient(httpClient, divisionID)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(baseURL)
	}

	client.SetDisallowUnknownFields(true)

	m.Run()
}
