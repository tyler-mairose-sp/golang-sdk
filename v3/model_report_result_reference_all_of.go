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

// checks if the ReportResultReferenceAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportResultReferenceAllOf{}

// ReportResultReferenceAllOf struct for ReportResultReferenceAllOf
type ReportResultReferenceAllOf struct {
	// Status of a violation report
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReportResultReferenceAllOf ReportResultReferenceAllOf

// NewReportResultReferenceAllOf instantiates a new ReportResultReferenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportResultReferenceAllOf() *ReportResultReferenceAllOf {
	this := ReportResultReferenceAllOf{}
	return &this
}

// NewReportResultReferenceAllOfWithDefaults instantiates a new ReportResultReferenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportResultReferenceAllOfWithDefaults() *ReportResultReferenceAllOf {
	this := ReportResultReferenceAllOf{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReportResultReferenceAllOf) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResultReferenceAllOf) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReportResultReferenceAllOf) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReportResultReferenceAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o ReportResultReferenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportResultReferenceAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReportResultReferenceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varReportResultReferenceAllOf := _ReportResultReferenceAllOf{}

	if err = json.Unmarshal(bytes, &varReportResultReferenceAllOf); err == nil {
		*o = ReportResultReferenceAllOf(varReportResultReferenceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReportResultReferenceAllOf struct {
	value *ReportResultReferenceAllOf
	isSet bool
}

func (v NullableReportResultReferenceAllOf) Get() *ReportResultReferenceAllOf {
	return v.value
}

func (v *NullableReportResultReferenceAllOf) Set(val *ReportResultReferenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportResultReferenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportResultReferenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportResultReferenceAllOf(val *ReportResultReferenceAllOf) *NullableReportResultReferenceAllOf {
	return &NullableReportResultReferenceAllOf{value: val, isSet: true}
}

func (v NullableReportResultReferenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportResultReferenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


