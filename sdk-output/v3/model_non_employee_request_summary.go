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

// NonEmployeeRequestSummary struct for NonEmployeeRequestSummary
type NonEmployeeRequestSummary struct {
	// The number of approved non-employee requests on all sources that *requested-for* user manages.
	Approved *int32 `json:"approved,omitempty"`
	// The number of rejected non-employee requests on all sources that *requested-for* user manages.
	Rejected *int32 `json:"rejected,omitempty"`
	// The number of pending non-employee requests on all sources that *requested-for* user manages.
	Pending *int32 `json:"pending,omitempty"`
	// The number of non-employee records on all sources that *requested-for* user manages.
	NonEmployeeCount *int32 `json:"nonEmployeeCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeRequestSummary NonEmployeeRequestSummary

// NewNonEmployeeRequestSummary instantiates a new NonEmployeeRequestSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeRequestSummary() *NonEmployeeRequestSummary {
	this := NonEmployeeRequestSummary{}
	return &this
}

// NewNonEmployeeRequestSummaryWithDefaults instantiates a new NonEmployeeRequestSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeRequestSummaryWithDefaults() *NonEmployeeRequestSummary {
	this := NonEmployeeRequestSummary{}
	return &this
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *NonEmployeeRequestSummary) GetApproved() int32 {
	if o == nil || isNil(o.Approved) {
		var ret int32
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestSummary) GetApprovedOk() (*int32, bool) {
	if o == nil || isNil(o.Approved) {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *NonEmployeeRequestSummary) HasApproved() bool {
	if o != nil && !isNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given int32 and assigns it to the Approved field.
func (o *NonEmployeeRequestSummary) SetApproved(v int32) {
	o.Approved = &v
}

// GetRejected returns the Rejected field value if set, zero value otherwise.
func (o *NonEmployeeRequestSummary) GetRejected() int32 {
	if o == nil || isNil(o.Rejected) {
		var ret int32
		return ret
	}
	return *o.Rejected
}

// GetRejectedOk returns a tuple with the Rejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestSummary) GetRejectedOk() (*int32, bool) {
	if o == nil || isNil(o.Rejected) {
		return nil, false
	}
	return o.Rejected, true
}

// HasRejected returns a boolean if a field has been set.
func (o *NonEmployeeRequestSummary) HasRejected() bool {
	if o != nil && !isNil(o.Rejected) {
		return true
	}

	return false
}

// SetRejected gets a reference to the given int32 and assigns it to the Rejected field.
func (o *NonEmployeeRequestSummary) SetRejected(v int32) {
	o.Rejected = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *NonEmployeeRequestSummary) GetPending() int32 {
	if o == nil || isNil(o.Pending) {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestSummary) GetPendingOk() (*int32, bool) {
	if o == nil || isNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *NonEmployeeRequestSummary) HasPending() bool {
	if o != nil && !isNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *NonEmployeeRequestSummary) SetPending(v int32) {
	o.Pending = &v
}

// GetNonEmployeeCount returns the NonEmployeeCount field value if set, zero value otherwise.
func (o *NonEmployeeRequestSummary) GetNonEmployeeCount() int32 {
	if o == nil || isNil(o.NonEmployeeCount) {
		var ret int32
		return ret
	}
	return *o.NonEmployeeCount
}

// GetNonEmployeeCountOk returns a tuple with the NonEmployeeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeRequestSummary) GetNonEmployeeCountOk() (*int32, bool) {
	if o == nil || isNil(o.NonEmployeeCount) {
		return nil, false
	}
	return o.NonEmployeeCount, true
}

// HasNonEmployeeCount returns a boolean if a field has been set.
func (o *NonEmployeeRequestSummary) HasNonEmployeeCount() bool {
	if o != nil && !isNil(o.NonEmployeeCount) {
		return true
	}

	return false
}

// SetNonEmployeeCount gets a reference to the given int32 and assigns it to the NonEmployeeCount field.
func (o *NonEmployeeRequestSummary) SetNonEmployeeCount(v int32) {
	o.NonEmployeeCount = &v
}

func (o NonEmployeeRequestSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	if !isNil(o.Rejected) {
		toSerialize["rejected"] = o.Rejected
	}
	if !isNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	if !isNil(o.NonEmployeeCount) {
		toSerialize["nonEmployeeCount"] = o.NonEmployeeCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NonEmployeeRequestSummary) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeRequestSummary := _NonEmployeeRequestSummary{}

	if err = json.Unmarshal(bytes, &varNonEmployeeRequestSummary); err == nil {
		*o = NonEmployeeRequestSummary(varNonEmployeeRequestSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "approved")
		delete(additionalProperties, "rejected")
		delete(additionalProperties, "pending")
		delete(additionalProperties, "nonEmployeeCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeRequestSummary struct {
	value *NonEmployeeRequestSummary
	isSet bool
}

func (v NullableNonEmployeeRequestSummary) Get() *NonEmployeeRequestSummary {
	return v.value
}

func (v *NullableNonEmployeeRequestSummary) Set(val *NonEmployeeRequestSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeRequestSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeRequestSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeRequestSummary(val *NonEmployeeRequestSummary) *NullableNonEmployeeRequestSummary {
	return &NullableNonEmployeeRequestSummary{value: val, isSet: true}
}

func (v NullableNonEmployeeRequestSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeRequestSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


