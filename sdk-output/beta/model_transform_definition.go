/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
)

// TransformDefinition struct for TransformDefinition
type TransformDefinition struct {
	// The type of the transform definition.
	Type *string `json:"type,omitempty"`
	// Arbitrary key-value pairs to store any metadata for the object
	Attributes map[string]map[string]interface{} `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransformDefinition TransformDefinition

// NewTransformDefinition instantiates a new TransformDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformDefinition() *TransformDefinition {
	this := TransformDefinition{}
	return &this
}

// NewTransformDefinitionWithDefaults instantiates a new TransformDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformDefinitionWithDefaults() *TransformDefinition {
	this := TransformDefinition{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransformDefinition) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformDefinition) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransformDefinition) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransformDefinition) SetType(v string) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TransformDefinition) GetAttributes() map[string]map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformDefinition) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TransformDefinition) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *TransformDefinition) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = v
}

func (o TransformDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransformDefinition) UnmarshalJSON(bytes []byte) (err error) {
	varTransformDefinition := _TransformDefinition{}

	if err = json.Unmarshal(bytes, &varTransformDefinition); err == nil {
		*o = TransformDefinition(varTransformDefinition)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransformDefinition struct {
	value *TransformDefinition
	isSet bool
}

func (v NullableTransformDefinition) Get() *TransformDefinition {
	return v.value
}

func (v *NullableTransformDefinition) Set(val *TransformDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformDefinition(val *TransformDefinition) *NullableTransformDefinition {
	return &NullableTransformDefinition{value: val, isSet: true}
}

func (v NullableTransformDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


