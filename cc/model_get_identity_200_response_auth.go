/*
IdentityNow cc (private) APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cc

import (
	"encoding/json"
)

// checks if the GetIdentity200ResponseAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIdentity200ResponseAuth{}

// GetIdentity200ResponseAuth struct for GetIdentity200ResponseAuth
type GetIdentity200ResponseAuth struct {
	Service *string `json:"service,omitempty"`
	Encryption *string `json:"encryption,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIdentity200ResponseAuth GetIdentity200ResponseAuth

// NewGetIdentity200ResponseAuth instantiates a new GetIdentity200ResponseAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdentity200ResponseAuth() *GetIdentity200ResponseAuth {
	this := GetIdentity200ResponseAuth{}
	return &this
}

// NewGetIdentity200ResponseAuthWithDefaults instantiates a new GetIdentity200ResponseAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdentity200ResponseAuthWithDefaults() *GetIdentity200ResponseAuth {
	this := GetIdentity200ResponseAuth{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *GetIdentity200ResponseAuth) GetService() string {
	if o == nil || isNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdentity200ResponseAuth) GetServiceOk() (*string, bool) {
	if o == nil || isNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *GetIdentity200ResponseAuth) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *GetIdentity200ResponseAuth) SetService(v string) {
	o.Service = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *GetIdentity200ResponseAuth) GetEncryption() string {
	if o == nil || isNil(o.Encryption) {
		var ret string
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdentity200ResponseAuth) GetEncryptionOk() (*string, bool) {
	if o == nil || isNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *GetIdentity200ResponseAuth) HasEncryption() bool {
	if o != nil && !isNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given string and assigns it to the Encryption field.
func (o *GetIdentity200ResponseAuth) SetEncryption(v string) {
	o.Encryption = &v
}

func (o GetIdentity200ResponseAuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIdentity200ResponseAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.Encryption) {
		toSerialize["encryption"] = o.Encryption
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetIdentity200ResponseAuth) UnmarshalJSON(bytes []byte) (err error) {
	varGetIdentity200ResponseAuth := _GetIdentity200ResponseAuth{}

	if err = json.Unmarshal(bytes, &varGetIdentity200ResponseAuth); err == nil {
		*o = GetIdentity200ResponseAuth(varGetIdentity200ResponseAuth)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service")
		delete(additionalProperties, "encryption")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIdentity200ResponseAuth struct {
	value *GetIdentity200ResponseAuth
	isSet bool
}

func (v NullableGetIdentity200ResponseAuth) Get() *GetIdentity200ResponseAuth {
	return v.value
}

func (v *NullableGetIdentity200ResponseAuth) Set(val *GetIdentity200ResponseAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdentity200ResponseAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdentity200ResponseAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdentity200ResponseAuth(val *GetIdentity200ResponseAuth) *NullableGetIdentity200ResponseAuth {
	return &NullableGetIdentity200ResponseAuth{value: val, isSet: true}
}

func (v NullableGetIdentity200ResponseAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdentity200ResponseAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


