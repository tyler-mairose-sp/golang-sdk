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

// checks if the Value type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Value{}

// Value struct for Value
type Value struct {
	// The type of attribute value
	Type *string `json:"type,omitempty"`
	// The attribute value
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Value Value

// NewValue instantiates a new Value object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValue() *Value {
	this := Value{}
	return &this
}

// NewValueWithDefaults instantiates a new Value object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueWithDefaults() *Value {
	this := Value{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Value) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Value) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Value) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Value) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Value) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Value) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Value) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Value) SetValue(v string) {
	o.Value = &v
}

func (o Value) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Value) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Value) UnmarshalJSON(bytes []byte) (err error) {
	varValue := _Value{}

	if err = json.Unmarshal(bytes, &varValue); err == nil {
		*o = Value(varValue)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValue struct {
	value *Value
	isSet bool
}

func (v NullableValue) Get() *Value {
	return v.value
}

func (v *NullableValue) Set(val *Value) {
	v.value = val
	v.isSet = true
}

func (v NullableValue) IsSet() bool {
	return v.isSet
}

func (v *NullableValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValue(val *Value) *NullableValue {
	return &NullableValue{value: val, isSet: true}
}

func (v NullableValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


