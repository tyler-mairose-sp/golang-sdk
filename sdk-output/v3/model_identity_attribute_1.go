/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// IdentityAttribute1 struct for IdentityAttribute1
type IdentityAttribute1 struct {
	// The attribute key
	Key *string `json:"key,omitempty"`
	// Human-readable display name of the attribute
	Name *string `json:"name,omitempty"`
	// The attribute value
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityAttribute1 IdentityAttribute1

// NewIdentityAttribute1 instantiates a new IdentityAttribute1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAttribute1() *IdentityAttribute1 {
	this := IdentityAttribute1{}
	return &this
}

// NewIdentityAttribute1WithDefaults instantiates a new IdentityAttribute1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAttribute1WithDefaults() *IdentityAttribute1 {
	this := IdentityAttribute1{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *IdentityAttribute1) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttribute1) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *IdentityAttribute1) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *IdentityAttribute1) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityAttribute1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttribute1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityAttribute1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityAttribute1) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IdentityAttribute1) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttribute1) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IdentityAttribute1) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IdentityAttribute1) SetValue(v string) {
	o.Value = &v
}

func (o IdentityAttribute1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityAttribute1) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityAttribute1 := _IdentityAttribute1{}

	if err = json.Unmarshal(bytes, &varIdentityAttribute1); err == nil {
		*o = IdentityAttribute1(varIdentityAttribute1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityAttribute1 struct {
	value *IdentityAttribute1
	isSet bool
}

func (v NullableIdentityAttribute1) Get() *IdentityAttribute1 {
	return v.value
}

func (v *NullableIdentityAttribute1) Set(val *IdentityAttribute1) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAttribute1) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAttribute1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAttribute1(val *IdentityAttribute1) *NullableIdentityAttribute1 {
	return &NullableIdentityAttribute1{value: val, isSet: true}
}

func (v NullableIdentityAttribute1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAttribute1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


