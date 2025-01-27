/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// checks if the AggregationAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregationAllOf{}

// AggregationAllOf struct for AggregationAllOf
type AggregationAllOf struct {
	Status *string `json:"status,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	AvgDuration *int32 `json:"avgDuration,omitempty"`
	ChangedAccounts *int32 `json:"changedAccounts,omitempty"`
	// A date-time in ISO-8601 format
	NextScheduled NullableTime `json:"nextScheduled,omitempty"`
	// A date-time in ISO-8601 format
	StartTime NullableTime `json:"startTime,omitempty"`
	// John Doe
	SourceOwner *string `json:"sourceOwner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AggregationAllOf AggregationAllOf

// NewAggregationAllOf instantiates a new AggregationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregationAllOf() *AggregationAllOf {
	this := AggregationAllOf{}
	return &this
}

// NewAggregationAllOfWithDefaults instantiates a new AggregationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationAllOfWithDefaults() *AggregationAllOf {
	this := AggregationAllOf{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AggregationAllOf) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationAllOf) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AggregationAllOf) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AggregationAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AggregationAllOf) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationAllOf) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AggregationAllOf) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *AggregationAllOf) SetDuration(v int32) {
	o.Duration = &v
}

// GetAvgDuration returns the AvgDuration field value if set, zero value otherwise.
func (o *AggregationAllOf) GetAvgDuration() int32 {
	if o == nil || isNil(o.AvgDuration) {
		var ret int32
		return ret
	}
	return *o.AvgDuration
}

// GetAvgDurationOk returns a tuple with the AvgDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationAllOf) GetAvgDurationOk() (*int32, bool) {
	if o == nil || isNil(o.AvgDuration) {
		return nil, false
	}
	return o.AvgDuration, true
}

// HasAvgDuration returns a boolean if a field has been set.
func (o *AggregationAllOf) HasAvgDuration() bool {
	if o != nil && !isNil(o.AvgDuration) {
		return true
	}

	return false
}

// SetAvgDuration gets a reference to the given int32 and assigns it to the AvgDuration field.
func (o *AggregationAllOf) SetAvgDuration(v int32) {
	o.AvgDuration = &v
}

// GetChangedAccounts returns the ChangedAccounts field value if set, zero value otherwise.
func (o *AggregationAllOf) GetChangedAccounts() int32 {
	if o == nil || isNil(o.ChangedAccounts) {
		var ret int32
		return ret
	}
	return *o.ChangedAccounts
}

// GetChangedAccountsOk returns a tuple with the ChangedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationAllOf) GetChangedAccountsOk() (*int32, bool) {
	if o == nil || isNil(o.ChangedAccounts) {
		return nil, false
	}
	return o.ChangedAccounts, true
}

// HasChangedAccounts returns a boolean if a field has been set.
func (o *AggregationAllOf) HasChangedAccounts() bool {
	if o != nil && !isNil(o.ChangedAccounts) {
		return true
	}

	return false
}

// SetChangedAccounts gets a reference to the given int32 and assigns it to the ChangedAccounts field.
func (o *AggregationAllOf) SetChangedAccounts(v int32) {
	o.ChangedAccounts = &v
}

// GetNextScheduled returns the NextScheduled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AggregationAllOf) GetNextScheduled() time.Time {
	if o == nil || isNil(o.NextScheduled.Get()) {
		var ret time.Time
		return ret
	}
	return *o.NextScheduled.Get()
}

// GetNextScheduledOk returns a tuple with the NextScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AggregationAllOf) GetNextScheduledOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextScheduled.Get(), o.NextScheduled.IsSet()
}

// HasNextScheduled returns a boolean if a field has been set.
func (o *AggregationAllOf) HasNextScheduled() bool {
	if o != nil && o.NextScheduled.IsSet() {
		return true
	}

	return false
}

// SetNextScheduled gets a reference to the given NullableTime and assigns it to the NextScheduled field.
func (o *AggregationAllOf) SetNextScheduled(v time.Time) {
	o.NextScheduled.Set(&v)
}
// SetNextScheduledNil sets the value for NextScheduled to be an explicit nil
func (o *AggregationAllOf) SetNextScheduledNil() {
	o.NextScheduled.Set(nil)
}

// UnsetNextScheduled ensures that no value is present for NextScheduled, not even an explicit nil
func (o *AggregationAllOf) UnsetNextScheduled() {
	o.NextScheduled.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AggregationAllOf) GetStartTime() time.Time {
	if o == nil || isNil(o.StartTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AggregationAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *AggregationAllOf) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *AggregationAllOf) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *AggregationAllOf) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *AggregationAllOf) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetSourceOwner returns the SourceOwner field value if set, zero value otherwise.
func (o *AggregationAllOf) GetSourceOwner() string {
	if o == nil || isNil(o.SourceOwner) {
		var ret string
		return ret
	}
	return *o.SourceOwner
}

// GetSourceOwnerOk returns a tuple with the SourceOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationAllOf) GetSourceOwnerOk() (*string, bool) {
	if o == nil || isNil(o.SourceOwner) {
		return nil, false
	}
	return o.SourceOwner, true
}

// HasSourceOwner returns a boolean if a field has been set.
func (o *AggregationAllOf) HasSourceOwner() bool {
	if o != nil && !isNil(o.SourceOwner) {
		return true
	}

	return false
}

// SetSourceOwner gets a reference to the given string and assigns it to the SourceOwner field.
func (o *AggregationAllOf) SetSourceOwner(v string) {
	o.SourceOwner = &v
}

func (o AggregationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregationAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.AvgDuration) {
		toSerialize["avgDuration"] = o.AvgDuration
	}
	if !isNil(o.ChangedAccounts) {
		toSerialize["changedAccounts"] = o.ChangedAccounts
	}
	if o.NextScheduled.IsSet() {
		toSerialize["nextScheduled"] = o.NextScheduled.Get()
	}
	if o.StartTime.IsSet() {
		toSerialize["startTime"] = o.StartTime.Get()
	}
	if !isNil(o.SourceOwner) {
		toSerialize["sourceOwner"] = o.SourceOwner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AggregationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAggregationAllOf := _AggregationAllOf{}

	if err = json.Unmarshal(bytes, &varAggregationAllOf); err == nil {
		*o = AggregationAllOf(varAggregationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "avgDuration")
		delete(additionalProperties, "changedAccounts")
		delete(additionalProperties, "nextScheduled")
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "sourceOwner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAggregationAllOf struct {
	value *AggregationAllOf
	isSet bool
}

func (v NullableAggregationAllOf) Get() *AggregationAllOf {
	return v.value
}

func (v *NullableAggregationAllOf) Set(val *AggregationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationAllOf(val *AggregationAllOf) *NullableAggregationAllOf {
	return &NullableAggregationAllOf{value: val, isSet: true}
}

func (v NullableAggregationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


