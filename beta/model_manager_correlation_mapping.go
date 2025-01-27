/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the ManagerCorrelationMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagerCorrelationMapping{}

// ManagerCorrelationMapping struct for ManagerCorrelationMapping
type ManagerCorrelationMapping struct {
	// Name of the attribute to use for manager correlation. The value found on the account attribute will be used to lookup the manager's identity.
	AccountAttribute *string `json:"accountAttribute,omitempty"`
	// Name of the identity attribute to search when trying to find a manager using the value from the accountAttribute.
	IdentityAttribute *string `json:"identityAttribute,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagerCorrelationMapping ManagerCorrelationMapping

// NewManagerCorrelationMapping instantiates a new ManagerCorrelationMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagerCorrelationMapping() *ManagerCorrelationMapping {
	this := ManagerCorrelationMapping{}
	return &this
}

// NewManagerCorrelationMappingWithDefaults instantiates a new ManagerCorrelationMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagerCorrelationMappingWithDefaults() *ManagerCorrelationMapping {
	this := ManagerCorrelationMapping{}
	return &this
}

// GetAccountAttribute returns the AccountAttribute field value if set, zero value otherwise.
func (o *ManagerCorrelationMapping) GetAccountAttribute() string {
	if o == nil || isNil(o.AccountAttribute) {
		var ret string
		return ret
	}
	return *o.AccountAttribute
}

// GetAccountAttributeOk returns a tuple with the AccountAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagerCorrelationMapping) GetAccountAttributeOk() (*string, bool) {
	if o == nil || isNil(o.AccountAttribute) {
		return nil, false
	}
	return o.AccountAttribute, true
}

// HasAccountAttribute returns a boolean if a field has been set.
func (o *ManagerCorrelationMapping) HasAccountAttribute() bool {
	if o != nil && !isNil(o.AccountAttribute) {
		return true
	}

	return false
}

// SetAccountAttribute gets a reference to the given string and assigns it to the AccountAttribute field.
func (o *ManagerCorrelationMapping) SetAccountAttribute(v string) {
	o.AccountAttribute = &v
}

// GetIdentityAttribute returns the IdentityAttribute field value if set, zero value otherwise.
func (o *ManagerCorrelationMapping) GetIdentityAttribute() string {
	if o == nil || isNil(o.IdentityAttribute) {
		var ret string
		return ret
	}
	return *o.IdentityAttribute
}

// GetIdentityAttributeOk returns a tuple with the IdentityAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagerCorrelationMapping) GetIdentityAttributeOk() (*string, bool) {
	if o == nil || isNil(o.IdentityAttribute) {
		return nil, false
	}
	return o.IdentityAttribute, true
}

// HasIdentityAttribute returns a boolean if a field has been set.
func (o *ManagerCorrelationMapping) HasIdentityAttribute() bool {
	if o != nil && !isNil(o.IdentityAttribute) {
		return true
	}

	return false
}

// SetIdentityAttribute gets a reference to the given string and assigns it to the IdentityAttribute field.
func (o *ManagerCorrelationMapping) SetIdentityAttribute(v string) {
	o.IdentityAttribute = &v
}

func (o ManagerCorrelationMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagerCorrelationMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountAttribute) {
		toSerialize["accountAttribute"] = o.AccountAttribute
	}
	if !isNil(o.IdentityAttribute) {
		toSerialize["identityAttribute"] = o.IdentityAttribute
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagerCorrelationMapping) UnmarshalJSON(bytes []byte) (err error) {
	varManagerCorrelationMapping := _ManagerCorrelationMapping{}

	if err = json.Unmarshal(bytes, &varManagerCorrelationMapping); err == nil {
		*o = ManagerCorrelationMapping(varManagerCorrelationMapping)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountAttribute")
		delete(additionalProperties, "identityAttribute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagerCorrelationMapping struct {
	value *ManagerCorrelationMapping
	isSet bool
}

func (v NullableManagerCorrelationMapping) Get() *ManagerCorrelationMapping {
	return v.value
}

func (v *NullableManagerCorrelationMapping) Set(val *ManagerCorrelationMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableManagerCorrelationMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableManagerCorrelationMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagerCorrelationMapping(val *ManagerCorrelationMapping) *NullableManagerCorrelationMapping {
	return &NullableManagerCorrelationMapping{value: val, isSet: true}
}

func (v NullableManagerCorrelationMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagerCorrelationMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


