/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer( "%5B", "[", "%5D", "]" )
)

// APIClient manages communication with the IdentityNow Beta API API v3.1.0-beta
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccessProfilesApi *AccessProfilesApiService

	AccessRequestApprovalsApi *AccessRequestApprovalsApiService

	AccessRequestsApi *AccessRequestsApiService

	AccountActivitiesApi *AccountActivitiesApiService

	AccountAggregationsApi *AccountAggregationsApiService

	AccountsApi *AccountsApiService

	CertificationCampaignsApi *CertificationCampaignsApiService

	CertificationsApi *CertificationsApiService

	ConnectorRuleManagementApi *ConnectorRuleManagementApiService

	ConnectorsApi *ConnectorsApiService

	CustomPasswordInstructionsApi *CustomPasswordInstructionsApiService

	EntitlementsApi *EntitlementsApiService

	IAIAccessRequestRecommendationsApi *IAIAccessRequestRecommendationsApiService

	IAICommonAccessApi *IAICommonAccessApiService

	IAIOutliersApi *IAIOutliersApiService

	IAIPeerGroupStrategiesApi *IAIPeerGroupStrategiesApiService

	IAIRecommendationsApi *IAIRecommendationsApiService

	IAIRoleMiningApi *IAIRoleMiningApiService

	IdentitiesApi *IdentitiesApiService

	IdentityHistoryApi *IdentityHistoryApiService

	IdentityProfilesApi *IdentityProfilesApiService

	LifecycleStatesApi *LifecycleStatesApiService

	MFAConfigurationApi *MFAConfigurationApiService

	ManagedClientsApi *ManagedClientsApiService

	ManagedClustersApi *ManagedClustersApiService

	NonEmployeeLifecycleManagementApi *NonEmployeeLifecycleManagementApiService

	NotificationsApi *NotificationsApiService

	OAuthClientsApi *OAuthClientsApiService

	OrgConfigApi *OrgConfigApiService

	PasswordConfigurationApi *PasswordConfigurationApiService

	PasswordDictionaryApi *PasswordDictionaryApiService

	PasswordManagementApi *PasswordManagementApiService

	PasswordSyncGroupsApi *PasswordSyncGroupsApiService

	PersonalAccessTokensApi *PersonalAccessTokensApiService

	PublicIdentitiesConfigApi *PublicIdentitiesConfigApiService

	RequestableObjectsApi *RequestableObjectsApiService

	RoleInsightsApi *RoleInsightsApiService

	RolesApi *RolesApiService

	SODPolicyApi *SODPolicyApiService

	SODViolationsApi *SODViolationsApiService

	SPConfigApi *SPConfigApiService

	SearchAttributeConfigurationApi *SearchAttributeConfigurationApiService

	SegmentsApi *SegmentsApiService

	ServiceDeskIntegrationApi *ServiceDeskIntegrationApiService

	SourcesApi *SourcesApiService

	TaggedObjectsApi *TaggedObjectsApiService

	TransformsApi *TransformsApiService

	TriggersApi *TriggersApiService

	WorkItemsApi *WorkItemsApiService

	WorkflowsApi *WorkflowsApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = retryablehttp.NewClient()
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccessProfilesApi = (*AccessProfilesApiService)(&c.common)
	c.AccessRequestApprovalsApi = (*AccessRequestApprovalsApiService)(&c.common)
	c.AccessRequestsApi = (*AccessRequestsApiService)(&c.common)
	c.AccountActivitiesApi = (*AccountActivitiesApiService)(&c.common)
	c.AccountAggregationsApi = (*AccountAggregationsApiService)(&c.common)
	c.AccountsApi = (*AccountsApiService)(&c.common)
	c.CertificationCampaignsApi = (*CertificationCampaignsApiService)(&c.common)
	c.CertificationsApi = (*CertificationsApiService)(&c.common)
	c.ConnectorRuleManagementApi = (*ConnectorRuleManagementApiService)(&c.common)
	c.ConnectorsApi = (*ConnectorsApiService)(&c.common)
	c.CustomPasswordInstructionsApi = (*CustomPasswordInstructionsApiService)(&c.common)
	c.EntitlementsApi = (*EntitlementsApiService)(&c.common)
	c.IAIAccessRequestRecommendationsApi = (*IAIAccessRequestRecommendationsApiService)(&c.common)
	c.IAICommonAccessApi = (*IAICommonAccessApiService)(&c.common)
	c.IAIOutliersApi = (*IAIOutliersApiService)(&c.common)
	c.IAIPeerGroupStrategiesApi = (*IAIPeerGroupStrategiesApiService)(&c.common)
	c.IAIRecommendationsApi = (*IAIRecommendationsApiService)(&c.common)
	c.IAIRoleMiningApi = (*IAIRoleMiningApiService)(&c.common)
	c.IdentitiesApi = (*IdentitiesApiService)(&c.common)
	c.IdentityHistoryApi = (*IdentityHistoryApiService)(&c.common)
	c.IdentityProfilesApi = (*IdentityProfilesApiService)(&c.common)
	c.LifecycleStatesApi = (*LifecycleStatesApiService)(&c.common)
	c.MFAConfigurationApi = (*MFAConfigurationApiService)(&c.common)
	c.ManagedClientsApi = (*ManagedClientsApiService)(&c.common)
	c.ManagedClustersApi = (*ManagedClustersApiService)(&c.common)
	c.NonEmployeeLifecycleManagementApi = (*NonEmployeeLifecycleManagementApiService)(&c.common)
	c.NotificationsApi = (*NotificationsApiService)(&c.common)
	c.OAuthClientsApi = (*OAuthClientsApiService)(&c.common)
	c.OrgConfigApi = (*OrgConfigApiService)(&c.common)
	c.PasswordConfigurationApi = (*PasswordConfigurationApiService)(&c.common)
	c.PasswordDictionaryApi = (*PasswordDictionaryApiService)(&c.common)
	c.PasswordManagementApi = (*PasswordManagementApiService)(&c.common)
	c.PasswordSyncGroupsApi = (*PasswordSyncGroupsApiService)(&c.common)
	c.PersonalAccessTokensApi = (*PersonalAccessTokensApiService)(&c.common)
	c.PublicIdentitiesConfigApi = (*PublicIdentitiesConfigApiService)(&c.common)
	c.RequestableObjectsApi = (*RequestableObjectsApiService)(&c.common)
	c.RoleInsightsApi = (*RoleInsightsApiService)(&c.common)
	c.RolesApi = (*RolesApiService)(&c.common)
	c.SODPolicyApi = (*SODPolicyApiService)(&c.common)
	c.SODViolationsApi = (*SODViolationsApiService)(&c.common)
	c.SPConfigApi = (*SPConfigApiService)(&c.common)
	c.SearchAttributeConfigurationApi = (*SearchAttributeConfigurationApiService)(&c.common)
	c.SegmentsApi = (*SegmentsApiService)(&c.common)
	c.ServiceDeskIntegrationApi = (*ServiceDeskIntegrationApiService)(&c.common)
	c.SourcesApi = (*SourcesApiService)(&c.common)
	c.TaggedObjectsApi = (*TaggedObjectsApiService)(&c.common)
	c.TransformsApi = (*TransformsApiService)(&c.common)
	c.TriggersApi = (*TriggersApiService)(&c.common)
	c.WorkItemsApi = (*WorkItemsApiService)(&c.common)
	c.WorkflowsApi = (*WorkflowsApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

func parameterValueToString( obj interface{}, key string ) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return fmt.Sprintf("%v", obj)
	}
	var param,ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap,err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToQuery adds the provided object to the url query supporting deep object syntax
func parameterAddToQuery(queryParams interface{}, keyPrefix string, obj interface{}, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
			case reflect.Invalid:
				value = "invalid"

			case reflect.Struct:
				if t,ok := obj.(MappedNullable); ok {
					dataMap,err := t.ToMap()
					if err != nil {
						return
					}
					parameterAddToQuery(queryParams, keyPrefix, dataMap, collectionType)
					return
				}
				if t, ok := obj.(time.Time); ok {
					parameterAddToQuery(queryParams, keyPrefix, t.Format(time.RFC3339), collectionType)
					return
				}
				value = v.Type().String() + " value"
			case reflect.Slice:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				var lenIndValue = indValue.Len()
				for i:=0;i<lenIndValue;i++ {
					var arrayValue = indValue.Index(i)
					parameterAddToQuery(queryParams, keyPrefix, arrayValue.Interface(), collectionType)
				}
				return

			case reflect.Map:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				iter := indValue.MapRange()
				for iter.Next() {
					k,v := iter.Key(), iter.Value()
					parameterAddToQuery(queryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), collectionType)
				}
				return

			case reflect.Interface:
				fallthrough
			case reflect.Ptr:
				parameterAddToQuery(queryParams, keyPrefix, v.Elem().Interface(), collectionType)
				return

			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				value = strconv.FormatInt(v.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				value = strconv.FormatUint(v.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
			case reflect.Bool:
				value = strconv.FormatBool(v.Bool())
			case reflect.String:
				value = v.String()
			default:
				value = v.Type().String() + " value"
		}
	}

	switch valuesMap := queryParams.(type) {
		case url.Values:
			if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
				valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix) + "," + value)
			} else {
				valuesMap.Add(keyPrefix, value)
			}
			break
		case map[string]string:
			valuesMap[keyPrefix] = value
			break
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.StandardClient().Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

func (c *APIClient) GetAPIClient() *APIClient {
	return c.common.client
}

func (c *APIClient) GetCommon() *service {
	return &c.common
}

type formFile struct {
		fileBytes []byte
		fileName string
		formFileName string
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
						return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
						return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)
	localVarRequest.Header.Add("X-SailPoint-SDK", "1.0.0")

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

type AccessToken struct {
	AccessToken         string `json:"access_token"`
	TokenType           string `json:"token_type"`
	ExpiresIn           int    `json:"expires_in"`
	Scope               string `json:"scope"`
	TenantId            string `json:"tenant_id"`
	Pod                 string `json:"pod"`
	StrongAuthSupported bool   `json:"strong_auth_supported"`
	Org                 string `json:"org"`
	IdentityId          string `json:"identity_id"`
	UserName            string `json:"user_name"`
	StrongAuth          bool   `json:"strong_auth"`
	Jti                 string `json:"jti"`
}

func getAccessToken(clientId string, clientSecret string, tokenURL string) (string, error) {
	url := tokenURL + "?grant_type=client_credentials&client_id=" + clientId + "&client_secret=" + clientSecret
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var jsonMap AccessToken
	json.Unmarshal([]byte(body), &jsonMap)

	return jsonMap.AccessToken, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = ioutil.TempFile("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = ioutil.TempFile("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if jsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// A wrapper for strict JSON decoding
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	field := metaValue.FieldByName("Title")
	if field != (reflect.Value{}) {
		str = fmt.Sprintf("%s", field.Interface())
	}

	field = metaValue.FieldByName("Detail")
	if field != (reflect.Value{}) {
		str = fmt.Sprintf("%s (%s)", str, field.Interface())
	}

	// status title (detail)
	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}
