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

// checks if the UpdateApplicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateApplicationRequest{}

// UpdateApplicationRequest struct for UpdateApplicationRequest
type UpdateApplicationRequest struct {
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateApplicationRequest UpdateApplicationRequest

// NewUpdateApplicationRequest instantiates a new UpdateApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApplicationRequest() *UpdateApplicationRequest {
	this := UpdateApplicationRequest{}
	return &this
}

// NewUpdateApplicationRequestWithDefaults instantiates a new UpdateApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApplicationRequestWithDefaults() *UpdateApplicationRequest {
	this := UpdateApplicationRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateApplicationRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateApplicationRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateApplicationRequest) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateApplicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateApplicationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateApplicationRequest := _UpdateApplicationRequest{}

	if err = json.Unmarshal(bytes, &varUpdateApplicationRequest); err == nil {
		*o = UpdateApplicationRequest(varUpdateApplicationRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateApplicationRequest struct {
	value *UpdateApplicationRequest
	isSet bool
}

func (v NullableUpdateApplicationRequest) Get() *UpdateApplicationRequest {
	return v.value
}

func (v *NullableUpdateApplicationRequest) Set(val *UpdateApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApplicationRequest(val *UpdateApplicationRequest) *NullableUpdateApplicationRequest {
	return &NullableUpdateApplicationRequest{value: val, isSet: true}
}

func (v NullableUpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


