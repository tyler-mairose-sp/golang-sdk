/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the DisplayReferenceAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisplayReferenceAllOf{}

// DisplayReferenceAllOf struct for DisplayReferenceAllOf
type DisplayReferenceAllOf struct {
	DisplayName *string `json:"displayName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DisplayReferenceAllOf DisplayReferenceAllOf

// NewDisplayReferenceAllOf instantiates a new DisplayReferenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisplayReferenceAllOf() *DisplayReferenceAllOf {
	this := DisplayReferenceAllOf{}
	return &this
}

// NewDisplayReferenceAllOfWithDefaults instantiates a new DisplayReferenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisplayReferenceAllOfWithDefaults() *DisplayReferenceAllOf {
	this := DisplayReferenceAllOf{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DisplayReferenceAllOf) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayReferenceAllOf) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DisplayReferenceAllOf) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DisplayReferenceAllOf) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o DisplayReferenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisplayReferenceAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DisplayReferenceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDisplayReferenceAllOf := _DisplayReferenceAllOf{}

	if err = json.Unmarshal(bytes, &varDisplayReferenceAllOf); err == nil {
		*o = DisplayReferenceAllOf(varDisplayReferenceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDisplayReferenceAllOf struct {
	value *DisplayReferenceAllOf
	isSet bool
}

func (v NullableDisplayReferenceAllOf) Get() *DisplayReferenceAllOf {
	return v.value
}

func (v *NullableDisplayReferenceAllOf) Set(val *DisplayReferenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDisplayReferenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDisplayReferenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisplayReferenceAllOf(val *DisplayReferenceAllOf) *NullableDisplayReferenceAllOf {
	return &NullableDisplayReferenceAllOf{value: val, isSet: true}
}

func (v NullableDisplayReferenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisplayReferenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


