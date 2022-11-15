/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpoint

import (
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

type ClientCredentials struct {
	ClientId string
	ClientSecret string
}

var (
	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = "accesstoken"

	ContextClientCredentials = "clientCredentials"
)

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL string
	Description string
	Variables map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
	Tenant			 string
}

// NewConfiguration returns a new Configuration object
func NewConfiguration(tenant string) *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/0.1.0/go",
		Debug:            false,
		Tenant: 		  tenant,		
		Servers:          ServerConfigurations{
			{
				URL: "https://" + tenant + ".api.identitynow.com/v3",
				Description: "This is the production API server.",
			},
		},
		OperationServers: map[string]ServerConfigurations{
		},
	}
	return cfg
}

