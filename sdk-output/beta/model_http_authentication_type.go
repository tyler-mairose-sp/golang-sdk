/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
	"fmt"
)

// HttpAuthenticationType Defines the HTTP Authentication type. Additional values may be added in the future.  If *NO_AUTH* is selected, no extra information will be in HttpConfig.  If *BASIC_AUTH* is selected, HttpConfig will include BasicAuthConfig with Username and Password as strings.  If *BEARER_TOKEN* is selected, HttpConfig will include BearerTokenAuthConfig with Token as string.
type HttpAuthenticationType string

// List of HttpAuthenticationType
const (
	HTTPAUTHENTICATIONTYPE_NO_AUTH HttpAuthenticationType = "NO_AUTH"
	HTTPAUTHENTICATIONTYPE_BASIC_AUTH HttpAuthenticationType = "BASIC_AUTH"
	HTTPAUTHENTICATIONTYPE_BEARER_TOKEN HttpAuthenticationType = "BEARER_TOKEN"
)

// All allowed values of HttpAuthenticationType enum
var AllowedHttpAuthenticationTypeEnumValues = []HttpAuthenticationType{
	"NO_AUTH",
	"BASIC_AUTH",
	"BEARER_TOKEN",
}

func (v *HttpAuthenticationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HttpAuthenticationType(value)
	for _, existing := range AllowedHttpAuthenticationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HttpAuthenticationType", value)
}

// NewHttpAuthenticationTypeFromValue returns a pointer to a valid HttpAuthenticationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHttpAuthenticationTypeFromValue(v string) (*HttpAuthenticationType, error) {
	ev := HttpAuthenticationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HttpAuthenticationType: valid values are %v", v, AllowedHttpAuthenticationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HttpAuthenticationType) IsValid() bool {
	for _, existing := range AllowedHttpAuthenticationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HttpAuthenticationType value
func (v HttpAuthenticationType) Ptr() *HttpAuthenticationType {
	return &v
}

type NullableHttpAuthenticationType struct {
	value *HttpAuthenticationType
	isSet bool
}

func (v NullableHttpAuthenticationType) Get() *HttpAuthenticationType {
	return v.value
}

func (v *NullableHttpAuthenticationType) Set(val *HttpAuthenticationType) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpAuthenticationType) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpAuthenticationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpAuthenticationType(val *HttpAuthenticationType) *NullableHttpAuthenticationType {
	return &NullableHttpAuthenticationType{value: val, isSet: true}
}

func (v NullableHttpAuthenticationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpAuthenticationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

