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

// checks if the IdentityAttributeConfig1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAttributeConfig1{}

// IdentityAttributeConfig1 Defines all the identity attribute mapping configurations. This defines how to generate or collect data for each identity attributes in identity refresh process.
type IdentityAttributeConfig1 struct {
	// The backend will only promote values if the profile/mapping is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	AttributeTransforms []IdentityAttributeTransform1 `json:"attributeTransforms,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityAttributeConfig1 IdentityAttributeConfig1

// NewIdentityAttributeConfig1 instantiates a new IdentityAttributeConfig1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAttributeConfig1() *IdentityAttributeConfig1 {
	this := IdentityAttributeConfig1{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewIdentityAttributeConfig1WithDefaults instantiates a new IdentityAttributeConfig1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAttributeConfig1WithDefaults() *IdentityAttributeConfig1 {
	this := IdentityAttributeConfig1{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IdentityAttributeConfig1) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttributeConfig1) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IdentityAttributeConfig1) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IdentityAttributeConfig1) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAttributeTransforms returns the AttributeTransforms field value if set, zero value otherwise.
func (o *IdentityAttributeConfig1) GetAttributeTransforms() []IdentityAttributeTransform1 {
	if o == nil || isNil(o.AttributeTransforms) {
		var ret []IdentityAttributeTransform1
		return ret
	}
	return o.AttributeTransforms
}

// GetAttributeTransformsOk returns a tuple with the AttributeTransforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAttributeConfig1) GetAttributeTransformsOk() ([]IdentityAttributeTransform1, bool) {
	if o == nil || isNil(o.AttributeTransforms) {
		return nil, false
	}
	return o.AttributeTransforms, true
}

// HasAttributeTransforms returns a boolean if a field has been set.
func (o *IdentityAttributeConfig1) HasAttributeTransforms() bool {
	if o != nil && !isNil(o.AttributeTransforms) {
		return true
	}

	return false
}

// SetAttributeTransforms gets a reference to the given []IdentityAttributeTransform1 and assigns it to the AttributeTransforms field.
func (o *IdentityAttributeConfig1) SetAttributeTransforms(v []IdentityAttributeTransform1) {
	o.AttributeTransforms = v
}

func (o IdentityAttributeConfig1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAttributeConfig1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AttributeTransforms) {
		toSerialize["attributeTransforms"] = o.AttributeTransforms
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityAttributeConfig1) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityAttributeConfig1 := _IdentityAttributeConfig1{}

	if err = json.Unmarshal(bytes, &varIdentityAttributeConfig1); err == nil {
		*o = IdentityAttributeConfig1(varIdentityAttributeConfig1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "attributeTransforms")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityAttributeConfig1 struct {
	value *IdentityAttributeConfig1
	isSet bool
}

func (v NullableIdentityAttributeConfig1) Get() *IdentityAttributeConfig1 {
	return v.value
}

func (v *NullableIdentityAttributeConfig1) Set(val *IdentityAttributeConfig1) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAttributeConfig1) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAttributeConfig1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAttributeConfig1(val *IdentityAttributeConfig1) *NullableIdentityAttributeConfig1 {
	return &NullableIdentityAttributeConfig1{value: val, isSet: true}
}

func (v NullableIdentityAttributeConfig1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAttributeConfig1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


