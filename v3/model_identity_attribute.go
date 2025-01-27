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

// checks if the IdentityAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAttribute{}

// IdentityAttribute struct for IdentityAttribute
type IdentityAttribute struct {
	// The attribute key
	Key *string `json:"key,omitempty"`
	// Human-readable display name of the attribute
	Name *string `json:"name,omitempty"`
	// The attribute value
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityAttribute IdentityAttribute

// NewIdentityAttribute instantiates a new IdentityAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAttribute() *IdentityAttribute {
	this := IdentityAttribute{}
	return &this
}

// NewIdentityAttributeWithDefaults instantiates a new IdentityAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAttributeWithDefaults() *IdentityAttribute {
	this := IdentityAttribute{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *IdentityAttribute) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttribute) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *IdentityAttribute) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *IdentityAttribute) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityAttribute) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttribute) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityAttribute) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityAttribute) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IdentityAttribute) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttribute) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IdentityAttribute) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IdentityAttribute) SetValue(v string) {
	o.Value = &v
}

func (o IdentityAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAttribute) ToMap() (map[string]interface{}, error) {
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

	return toSerialize, nil
}

func (o *IdentityAttribute) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityAttribute := _IdentityAttribute{}

	if err = json.Unmarshal(bytes, &varIdentityAttribute); err == nil {
		*o = IdentityAttribute(varIdentityAttribute)
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

type NullableIdentityAttribute struct {
	value *IdentityAttribute
	isSet bool
}

func (v NullableIdentityAttribute) Get() *IdentityAttribute {
	return v.value
}

func (v *NullableIdentityAttribute) Set(val *IdentityAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAttribute(val *IdentityAttribute) *NullableIdentityAttribute {
	return &NullableIdentityAttribute{value: val, isSet: true}
}

func (v NullableIdentityAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


