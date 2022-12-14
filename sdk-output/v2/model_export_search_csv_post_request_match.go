/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"encoding/json"
)

// ExportSearchCsvPostRequestMatch struct for ExportSearchCsvPostRequestMatch
type ExportSearchCsvPostRequestMatch struct {
	DisplayName *string `json:"displayName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExportSearchCsvPostRequestMatch ExportSearchCsvPostRequestMatch

// NewExportSearchCsvPostRequestMatch instantiates a new ExportSearchCsvPostRequestMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportSearchCsvPostRequestMatch() *ExportSearchCsvPostRequestMatch {
	this := ExportSearchCsvPostRequestMatch{}
	return &this
}

// NewExportSearchCsvPostRequestMatchWithDefaults instantiates a new ExportSearchCsvPostRequestMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportSearchCsvPostRequestMatchWithDefaults() *ExportSearchCsvPostRequestMatch {
	this := ExportSearchCsvPostRequestMatch{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ExportSearchCsvPostRequestMatch) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportSearchCsvPostRequestMatch) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ExportSearchCsvPostRequestMatch) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ExportSearchCsvPostRequestMatch) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o ExportSearchCsvPostRequestMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ExportSearchCsvPostRequestMatch) UnmarshalJSON(bytes []byte) (err error) {
	varExportSearchCsvPostRequestMatch := _ExportSearchCsvPostRequestMatch{}

	if err = json.Unmarshal(bytes, &varExportSearchCsvPostRequestMatch); err == nil {
		*o = ExportSearchCsvPostRequestMatch(varExportSearchCsvPostRequestMatch)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportSearchCsvPostRequestMatch struct {
	value *ExportSearchCsvPostRequestMatch
	isSet bool
}

func (v NullableExportSearchCsvPostRequestMatch) Get() *ExportSearchCsvPostRequestMatch {
	return v.value
}

func (v *NullableExportSearchCsvPostRequestMatch) Set(val *ExportSearchCsvPostRequestMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableExportSearchCsvPostRequestMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableExportSearchCsvPostRequestMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportSearchCsvPostRequestMatch(val *ExportSearchCsvPostRequestMatch) *NullableExportSearchCsvPostRequestMatch {
	return &NullableExportSearchCsvPostRequestMatch{value: val, isSet: true}
}

func (v NullableExportSearchCsvPostRequestMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportSearchCsvPostRequestMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


