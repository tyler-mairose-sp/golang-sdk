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

// FieldDetailsDto struct for FieldDetailsDto
type FieldDetailsDto struct {
	// The name of the attribute.
	Name *string `json:"name,omitempty"`
	// The transform to apply to the field
	Transform map[string]interface{} `json:"transform,omitempty"`
	// Attributes required for the transform
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Flag indicating whether or not the attribute is required.
	IsRequired *bool `json:"isRequired,omitempty"`
	// The type of the attribute.
	Type *string `json:"type,omitempty"`
	// Flag indicating whether or not the attribute is multi-valued.
	IsMultiValued *bool `json:"isMultiValued,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FieldDetailsDto FieldDetailsDto

// NewFieldDetailsDto instantiates a new FieldDetailsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldDetailsDto() *FieldDetailsDto {
	this := FieldDetailsDto{}
	var isMultiValued bool = false
	this.IsMultiValued = &isMultiValued
	return &this
}

// NewFieldDetailsDtoWithDefaults instantiates a new FieldDetailsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldDetailsDtoWithDefaults() *FieldDetailsDto {
	this := FieldDetailsDto{}
	var isMultiValued bool = false
	this.IsMultiValued = &isMultiValued
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FieldDetailsDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetailsDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FieldDetailsDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FieldDetailsDto) SetName(v string) {
	o.Name = &v
}

// GetTransform returns the Transform field value if set, zero value otherwise.
func (o *FieldDetailsDto) GetTransform() map[string]interface{} {
	if o == nil || isNil(o.Transform) {
		var ret map[string]interface{}
		return ret
	}
	return o.Transform
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetailsDto) GetTransformOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Transform) {
		return map[string]interface{}{}, false
	}
	return o.Transform, true
}

// HasTransform returns a boolean if a field has been set.
func (o *FieldDetailsDto) HasTransform() bool {
	if o != nil && !isNil(o.Transform) {
		return true
	}

	return false
}

// SetTransform gets a reference to the given map[string]interface{} and assigns it to the Transform field.
func (o *FieldDetailsDto) SetTransform(v map[string]interface{}) {
	o.Transform = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FieldDetailsDto) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetailsDto) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FieldDetailsDto) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *FieldDetailsDto) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *FieldDetailsDto) GetIsRequired() bool {
	if o == nil || isNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetailsDto) GetIsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *FieldDetailsDto) HasIsRequired() bool {
	if o != nil && !isNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *FieldDetailsDto) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FieldDetailsDto) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetailsDto) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FieldDetailsDto) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FieldDetailsDto) SetType(v string) {
	o.Type = &v
}

// GetIsMultiValued returns the IsMultiValued field value if set, zero value otherwise.
func (o *FieldDetailsDto) GetIsMultiValued() bool {
	if o == nil || isNil(o.IsMultiValued) {
		var ret bool
		return ret
	}
	return *o.IsMultiValued
}

// GetIsMultiValuedOk returns a tuple with the IsMultiValued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDetailsDto) GetIsMultiValuedOk() (*bool, bool) {
	if o == nil || isNil(o.IsMultiValued) {
		return nil, false
	}
	return o.IsMultiValued, true
}

// HasIsMultiValued returns a boolean if a field has been set.
func (o *FieldDetailsDto) HasIsMultiValued() bool {
	if o != nil && !isNil(o.IsMultiValued) {
		return true
	}

	return false
}

// SetIsMultiValued gets a reference to the given bool and assigns it to the IsMultiValued field.
func (o *FieldDetailsDto) SetIsMultiValued(v bool) {
	o.IsMultiValued = &v
}

func (o FieldDetailsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Transform) {
		toSerialize["transform"] = o.Transform
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.IsMultiValued) {
		toSerialize["isMultiValued"] = o.IsMultiValued
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FieldDetailsDto) UnmarshalJSON(bytes []byte) (err error) {
	varFieldDetailsDto := _FieldDetailsDto{}

	if err = json.Unmarshal(bytes, &varFieldDetailsDto); err == nil {
		*o = FieldDetailsDto(varFieldDetailsDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "transform")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "isRequired")
		delete(additionalProperties, "type")
		delete(additionalProperties, "isMultiValued")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFieldDetailsDto struct {
	value *FieldDetailsDto
	isSet bool
}

func (v NullableFieldDetailsDto) Get() *FieldDetailsDto {
	return v.value
}

func (v *NullableFieldDetailsDto) Set(val *FieldDetailsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldDetailsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldDetailsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldDetailsDto(val *FieldDetailsDto) *NullableFieldDetailsDto {
	return &NullableFieldDetailsDto{value: val, isSet: true}
}

func (v NullableFieldDetailsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldDetailsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


